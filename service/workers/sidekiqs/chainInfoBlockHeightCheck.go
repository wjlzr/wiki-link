package sidekiqs

import (
	"encoding/json"
	// "fmt"
	"strconv"
	"time"

	"github.com/oldfritter/sidekiq-go"

	"wiki-link/common"
	"wiki-link/core/chainInfo"
	"wiki-link/db"
	"wiki-link/model"
)

var (
	DefaultChianInfoPageSize = 100
)

func CreateChainInfoBlockHeightCheck(w *sidekiq.Worker) sidekiq.WorkerI {
	initAddressCheck()
	return &ChainInfoBlockHeightCheck{*w}
}

type ChainInfoBlockHeightCheck struct {
	sidekiq.Worker
}

func (worker *ChainInfoBlockHeightCheck) Work() (err error) {
	payload := make(map[string]string)
	json.Unmarshal([]byte(worker.Payload), &payload)
	start := time.Now().UnixNano()
	var (
		pageIndex = 1
		findNext  = true
	)
	for findNext {
		var req = chainInfo.TransactionsByBlockReq{
			Pagination:  chainInfo.Pagination{PageSize: DefaultChianInfoPageSize, PageIndex: pageIndex},
			BlockHeight: payload["blockHeight"],
		}
		resp, e := chainInfo.FindTransactionsByBlock(&req)
		if e != nil {
			return e
		}
		if (pageIndex+1)*DefaultChianInfoPageSize >= resp.TotalCount() {
			findNext = false
		} else {
			pageIndex += 1
		}
		switch payload["chain"] {
		case common.BTC:
			err = worker.parseBtcTransactionResp(resp, payload)
		}
	}
	worker.LogInfo("payload: ", worker.Payload, ", time:", (time.Now().UnixNano()-start)/1000000, " ms")
	return
}

func (worker *ChainInfoBlockHeightCheck) parseBtcTransactionResp(resp *chainInfo.BaseResp, payload map[string]string) (err error) {
	for _, item := range resp.Transactions() {
		re, e := chainInfo.FindTransaction(item.(map[string]interface{})["id"].(string))
		if e != nil {
			return e
		}
		worker.saveBtc(re, payload)
	}
	return
}

// 解析 & 存储 BTC
func (worker *ChainInfoBlockHeightCheck) saveBtc(re *chainInfo.BaseResp, payload map[string]string) {

	tx := db.BeginTx()
	defer tx.DbRollback()
	for _, input := range re.Inputs() {

		if input.Tag != "" {
			var i model.Address
			tx.Where(model.Address{Tag: input.Tag, Address: input.Address, Coin: payload["chain"]}).FirstOrInit(&i)
			i.Tag = input.Tag
			i.Coin = payload["chain"]
			i.Address = input.Address
			i.DataSource = "ChainInfo"
			if i.ID == 0 {
				i.UpdatedAt, _ = time.Parse(time.RFC3339, "2020-01-02T15:04:05Z")
			}
			if err := tx.Save(&i).Error; err != nil {
				worker.LogError(err)
			} else if payload["toward"] == "right" && time.Unix(re.Timestamp(), 0).Add(time.Hour).After(time.Now()) {
				addressCheckWorker.Priority(map[string]string{
					"address": i.Address,
					"chain":   i.Coin,
					"source":  "NewBlock",
					"id":      strconv.Itoa(int(i.ID)),
				})

			}
		}
	}

	for _, output := range re.Outputs() {
		if output.Tag != "" {
			var o model.Address
			tx.Where(model.Address{Tag: output.Tag, Address: output.Address, Coin: payload["chain"]}).FirstOrInit(&o)
			o.Tag = output.Tag
			o.Coin = payload["chain"]
			o.Address = output.Address
			o.DataSource = "ChainInfo"
			if o.ID == 0 {
				o.UpdatedAt, _ = time.Parse(time.RFC3339, "2020-01-02T15:04:05Z")
			}
			if err := tx.Save(&o).Error; err != nil {
				worker.LogError(err)
			} else if payload["toward"] == "right" && time.Unix(re.Timestamp(), 0).Add(time.Hour).After(time.Now()) {
				addressCheckWorker.Priority(map[string]string{
					"address": o.Address,
					"chain":   o.Coin,
					"source":  "NewBlock",
					"id":      strconv.Itoa(int(o.ID)),
				})
			}
		}
	}
	number, _ := strconv.Atoi(payload["blockHeight"])
	var bh model.BlockHeight
	tx.Where(model.BlockHeight{Coin: payload["chain"], Number: number}).FirstOrInit(&bh)
	tx.Save(&bh)
	tx.DbCommit()
}
