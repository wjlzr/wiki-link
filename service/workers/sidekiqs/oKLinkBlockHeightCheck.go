package sidekiqs

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/oldfritter/sidekiq-go"

	"wiki-link/common"
	"wiki-link/core/oklink"
	"wiki-link/db"
	hmodel "wiki-link/handler/model"
	"wiki-link/handler/util"
	"wiki-link/model"
)

var (
	DefaultOKLinkPageSize = 1000
	addressCheckWorker    sidekiq.WorkerI
)

func initAddressCheck() {
	if addressCheckWorker == nil {
		for _, w := range common.AllWorkers {
			if w.Name == "AddressCheck" {
				addressCheckWorker = common.AllWorkerIs[w.Name](&w)
				addressCheckWorker.SetClient(db.RedisClient())
			}
		}
	}
}

func CreateOKLinkBlockHeightCheck(w *sidekiq.Worker) sidekiq.WorkerI {
	initAddressCheck()
	return &OKLinkBlockHeightCheck{*w}
}

type OKLinkBlockHeightCheck struct {
	sidekiq.Worker
}

func (worker *OKLinkBlockHeightCheck) Work() (err error) {
	payload := make(map[string]string)
	json.Unmarshal([]byte(worker.Payload), &payload)
	start := time.Now().UnixNano()
	var (
		pageIndex = 1
		findNext  = true
	)
	for findNext {
		var req = hmodel.TransactionSearchReq{
			Info:       payload["blockHeight"],
			BaseReq:    hmodel.BaseReq{Chain: payload["chain"]},
			Pagination: hmodel.Pagination{PageSize: DefaultOKLinkPageSize, PageIndex: pageIndex},
		}
		resp, e := worker.getBlockHeight(&req)
		if e != nil {
			err = e
			continue
		}
		if resp.Size() < DefaultOKLinkPageSize {
			findNext = false
		} else {
			pageIndex += 1
		}
		switch req.Chain {
		case common.ETH:
			r := resp.(oklink.EthTransactionResp)
			if r.Code != 0 {
				err = fmt.Errorf(r.Msg)
				continue
			}
			worker.parseEthTransactionResp(&r, payload)
		case common.BTC:
			r := resp.(oklink.BtcTransactionResp)
			if r.Code != 0 {
				fmt.Errorf(r.Msg)
				continue
			}
			worker.parseBtcTransactionResp(&r, payload)
		}
	}
	worker.LogInfo("payload: ", worker.Payload, ", time:", (time.Now().UnixNano()-start)/1000000, " ms")
	return
}

func (worker *OKLinkBlockHeightCheck) getBlockHeight(req *hmodel.TransactionSearchReq) (oklink.TransactionResp, error) {
	var transactionsReq oklink.TransactionsReq
	if util.IsNumeric(req.Info) {
		util.Pagination(&req.PageSize, &req.PageIndex)
		transactionsReq.PageSize = req.PageSize
		transactionsReq.PageIndex = req.PageIndex
		transactionsReq.BlockHeight = req.Info
		return oklink.Transactions(req.Chain, &transactionsReq)
	}
	return nil, fmt.Errorf("无效的块高")
}

func (worker *OKLinkBlockHeightCheck) parseEthTransactionResp(resp *oklink.EthTransactionResp, payload map[string]string) {
	for _, hit := range resp.Data.Hits {
		worker.saveEth(&hit, payload)
	}
}

func (worker *OKLinkBlockHeightCheck) parseBtcTransactionResp(resp *oklink.BtcTransactionResp, payload map[string]string) {
	for _, hit := range resp.Data.Hits {
		worker.saveBtc(&hit, payload)
	}
}

// 解析 & 存储 ETH
func (worker *OKLinkBlockHeightCheck) saveEth(hit *oklink.EthTransaction, payload map[string]string) {
	tx := db.BeginTx()
	defer tx.DbRollback()
	fromTag := hit.FromTag.([]interface{})
	if hit.From != "" && len(fromTag) > 0 {
		f := fromTag[0].(map[string]interface{})
		var from model.Address
		var newFound bool
		tx.Where(model.Address{Tag: f["tag"].(string), Address: hit.From, Coin: "eth"}).FirstOrInit(&from)
		from.Tag = f["tag"].(string)
		from.Coin = payload["chain"]
		from.Type = f["type"].(string)
		from.Project = f["project"].(string)
		from.Address = hit.From
		from.DataSource = "OKLink"
		from.UpdatedAt, _ = time.Parse(time.RFC3339, "2020-01-02T15:04:05Z")
		if from.ID == 0 {
			newFound = true
			if err := tx.Save(&from).Error; err != nil {
				worker.LogError(err)
			}
		}
		if payload["toward"] == "right" && time.Unix(hit.Blocktime, 0).Add(time.Hour).After(time.Now()) {
			worker.priority(from.Coin, from.Address, "", strconv.Itoa(int(from.ID)))
		} else if newFound {
			worker.perform(from.Coin, from.Address, "", strconv.Itoa(int(from.ID)))
		}
	}
	toTag := hit.ToTag.([]interface{})
	if hit.To != "" && len(toTag) > 0 {
		t := toTag[0].(map[string]interface{})
		var to model.Address
		var newFound bool
		tx.Where(model.Address{Tag: t["tag"].(string), Address: hit.To, Coin: "eth"}).FirstOrInit(&to)
		to.Tag = t["tag"].(string)
		to.Coin = payload["chain"]
		to.Type = t["type"].(string)
		to.Project = t["project"].(string)
		to.Address = hit.To
		to.DataSource = "OKLink"
		to.UpdatedAt, _ = time.Parse(time.RFC3339, "2020-01-02T15:04:05Z")
		if to.ID == 0 {
			newFound = true
			if err := tx.Save(&to).Error; err != nil {
				worker.LogError(err)
			}
		}
		if payload["toward"] == "right" && time.Unix(hit.Blocktime, 0).Add(time.Hour).After(time.Now()) {
			worker.priority(to.Coin, to.Address, "NewBlock", strconv.Itoa(int(to.ID)))
		} else if newFound {
			worker.perform(to.Coin, to.Address, "", strconv.Itoa(int(to.ID)))
		}
	}
	number, _ := strconv.Atoi(payload["blockHeight"])
	var bh model.BlockHeight
	tx.Where(model.BlockHeight{Coin: "eth", Number: number}).FirstOrInit(&bh)
	tx.Save(&bh)
	tx.DbCommit()
}

// 解析 & 存储 BTC
func (worker *OKLinkBlockHeightCheck) saveBtc(hit *oklink.BtcTransaction, payload map[string]string) {
	tx := db.BeginTx()
	defer tx.DbRollback()
	for _, input := range hit.Inputs {
		for _, tag := range input.PreAddressesTags {
			for _, tagLog := range tag.TagLogos.([]interface{}) {
				attrs := tagLog.(map[string]interface{})

				if attrs["item"] != nil {
					var i model.Address
					var newFound bool
					tx.Where(model.Address{Tag: attrs["item"].(string), Address: tag.Address, Coin: payload["chain"]}).FirstOrInit(&i)
					i.Tag = attrs["item"].(string)
					i.Coin = payload["chain"]
					if attrs["type"] != nil {
						i.Type = attrs["type"].(string)
					}
					if attrs["project"] != nil {
						i.Project = attrs["project"].(string)
					}
					i.Address = tag.Address
					i.DataSource = "OKLink"
					i.UpdatedAt, _ = time.Parse(time.RFC3339, "2020-01-02T15:04:05Z")
					if i.ID == 0 {
						newFound = true
						if err := tx.Save(&i).Error; err != nil {
							worker.LogError(err)
						}
					}
					if payload["toward"] == "right" && time.Unix(hit.Blocktime, 0).Add(time.Hour).After(time.Now()) {
						worker.priority(i.Coin, i.Address, "NewBlock", strconv.Itoa(int(i.ID)))
					} else if newFound {
						worker.perform(i.Coin, i.Address, "", strconv.Itoa(int(i.ID)))
					}
				}
			}
		}
	}

	for _, output := range hit.Outputs {
		for _, tag := range output.AddressesTags {
			for _, tagLog := range tag.TagLogos.([]interface{}) {
				attrs := tagLog.(map[string]interface{})
				if attrs["item"] != nil {
					var o model.Address
					var newFound bool
					tx.Where(model.Address{Tag: attrs["item"].(string), Address: tag.Address, Coin: payload["chain"]}).FirstOrInit(&o)
					o.Tag = attrs["item"].(string)
					o.Coin = payload["chain"]
					if attrs["type"] != nil {
						o.Type = attrs["type"].(string)
					}
					if attrs["project"] != nil {
						o.Project = attrs["project"].(string)
					}
					o.Address = tag.Address
					o.DataSource = "OKLink"
					o.UpdatedAt, _ = time.Parse(time.RFC3339, "2020-01-02T15:04:05Z")
					if o.ID == 0 {
						newFound = true
						if err := tx.Save(&o).Error; err != nil {
							worker.LogError(err)
						}
					}
					if payload["toward"] == "right" && time.Unix(hit.Blocktime, 0).Add(time.Hour).After(time.Now()) {
						worker.priority(o.Coin, o.Address, "NewBlock", strconv.Itoa(int(o.ID)))
					} else if newFound {
						worker.perform(o.Coin, o.Address, "NewBlock", strconv.Itoa(int(o.ID)))
					}
				}
			}
		}
	}
	number, _ := strconv.Atoi(payload["blockHeight"])
	var bh model.BlockHeight
	tx.Where(model.BlockHeight{Coin: "btc", Number: number}).FirstOrInit(&bh)
	tx.Save(&bh)
	tx.DbCommit()
}

func (worker *OKLinkBlockHeightCheck) perform(chain, address, source, id string) {
	addressCheckWorker.Perform(map[string]string{"address": address, "chain": chain, "source": source, "id": id})
}

func (worker *OKLinkBlockHeightCheck) priority(chain, address, source, id string) {
	addressCheckWorker.Priority(map[string]string{"address": address, "chain": chain, "source": source, "id": id})
}
