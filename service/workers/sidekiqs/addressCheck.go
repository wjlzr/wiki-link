package sidekiqs

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/oldfritter/sidekiq-go"
	"github.com/shopspring/decimal"

	"wiki-link/common"
	"wiki-link/core/address"
	"wiki-link/core/bitGo"
	"wiki-link/core/blockChain"
	"wiki-link/core/btcCom"
	"wiki-link/core/chainInfo"
	"wiki-link/core/finance"
	"wiki-link/core/infura"
	"wiki-link/core/oklink"
	"wiki-link/db"
	hmodel "wiki-link/handler/model"
	"wiki-link/model"
)

var (
	index   = 0
	sources = []string{"BlockChain", "ChainInfo", "BtcCom", "BitGo"}
)

func CreateAddressCheck(w *sidekiq.Worker) sidekiq.WorkerI {
	return &AddressCheck{*w}
}

type AddressCheck struct {
	sidekiq.Worker
}

func (worker *AddressCheck) Work() (err error) {
	start := time.Now().UnixNano()
	payload := make(map[string]string)
	json.Unmarshal([]byte(worker.Payload), &payload)

	switch payload["chain"] {
	case common.ETH:
		// Infura 仅支持 ETH
		if err = worker.checkAddressFromInfura(payload); err != nil {
			return
		}
		// if err = worker.checkAddressFromFinance(payload); err != nil {
		//   return
		// }
		// if err = worker.checkAddressFromOKLink(payload); err != nil {
		//   return
		// }
	case common.BTC:
		switch sources[index] {
		case "BlockChain":
			err = worker.checkAddressFromBlockChain(payload)
		case "ChainInfo":
			err = worker.checkAddressFromChainInfo(payload)
		case "BtcCom":
			err = worker.checkAddressFromBtcCom(payload)
		case "BitGo":
			err = worker.checkAddressFromBitGo(payload)
		}
		if err != nil && errors.Is(err, ErrMaxRequest) && !BtcChannelChange.isLocked() {
			BtcChannelChange.doLock()
			index += 1
			if index >= len(sources) {
				index = 0
			}
			return
		}
	}

	// if err = worker.checkAddressFromOKLink(payload); err != nil {
	//   return
	// }

	worker.LogInfo("payload: ", worker.Payload, ", time:", (time.Now().UnixNano()-start)/1000000, " ms")
	return
}

func (worker *AddressCheck) checkAddressFromOKLink(payload map[string]string) (err error) {
	var req = hmodel.TransactionSearchReq{
		Info:    payload["address"],
		BaseReq: hmodel.BaseReq{Chain: payload["chain"]},
	}
	isAddress := address.Validation(req.Chain, req.Info)
	if isAddress {
		switch req.Chain {
		case common.ETH:
			var resp oklink.EthFindByAddressResp
			if e := oklink.FindByAddress(req.Chain, req.Info, &resp); e != nil {
				return e
			}
			if resp.Code != 0 {
				return fmt.Errorf(resp.Msg)
			}
			var address model.Address
			if err = db.MainDb.
				Where("address = ?", payload["address"]).
				Where("coin = ?", payload["chain"]).
				Where("id = ?", payload["id"]).
				First(&address).Error; err != nil {
				worker.LogError("没有数据:", err)
				return
			}
			address.Balance = decimal.NewFromFloat(resp.Data.Balance)
			var r struct {
				oklink.Base
				Data struct {
					Hits []struct {
						TokenContractAddress string  `json:"tokenContractAddress"`
						Value                float64 `json:"value"`
					} `json:"hits"`
				} `json:"data"`
			}
			if err := oklink.AddressHolders(req.Chain, &oklink.PageReq{
				ChainPagination: oklink.ChainPagination{Chain: payload["chain"]},
				Address:         payload["address"],
				TokenAddress:    "0xdac17f958d2ee523a2206206994597c13d831ec7",
			},
				&r,
			); err != nil {
				t := db.BeginTx()
				defer t.DbRollback()
				t.Save(&address)
				t.DbCommit()
			}

			tx := db.BeginTx()
			defer tx.DbRollback()
			var token model.Token
			tx.Where(model.Token{Coin: "usdt-erc20", AddressId: address.ID}).FirstOrInit(&token)
			for _, hit := range r.Data.Hits {
				if hit.TokenContractAddress == "0xdac17f958d2ee523a2206206994597c13d831ec7" {
					token.Balance = decimal.NewFromFloat(hit.Value)
				}
			}
			tx.Save(&address)
			tx.Save(&token)
			tx.DbCommit()

		case common.BTC:
			var resp oklink.BtcFindByAddressResp
			if e := oklink.FindByAddress(req.Chain, req.Info, &resp); e != nil {
				return e
			}
			if resp.Code != 0 {
				return fmt.Errorf(resp.Msg)
			}
			tx := db.BeginTx()
			defer tx.DbRollback()
			var address model.Address
			if err = tx.Where("address = ?", payload["address"]).
				Where("coin = ?", payload["chain"]).
				Where("id = ?", payload["id"]).
				First(&address).Error; err != nil {
				worker.LogError("没有数据:", err)
				return
			}
			address.Balance = decimal.NewFromFloat(resp.Data.Balance)
			tx.Save(&address)
			var token model.Token
			tx.Where(model.Token{Coin: "usdt-omni", AddressId: address.ID}).FirstOrInit(&token)
			token.Balance = decimal.NewFromFloat(resp.Data.UsdtBalance)
			tx.Save(&token)
			tx.DbCommit()
		}
	}
	return
}

func (worker *AddressCheck) checkAddressFromInfura(payload map[string]string) (err error) {
	resp, err := infura.FindBalanceByAddress(payload["address"])
	if err != nil {
		return
	}
	tx := db.BeginTx()
	defer tx.DbRollback()
	var address model.Address
	if err = tx.
		Where("address = ?", payload["address"]).
		Where("coin = ?", payload["chain"]).
		Where("id = ?", payload["id"]).
		First(&address).Error; err != nil {
		return
	}
	val := resp.GetData().(string)[2:]
	i := new(big.Int)
	i.SetString(val, 16)
	address.Balance = decimal.NewFromBigInt(i, -18)
	tx.Save(&address)
	tx.DbCommit()
	return
}

func (worker *AddressCheck) checkAddressFromFinance(payload map[string]string) (err error) {
	var resp finance.BaseResp
	req := finance.AddressReq{
		Address: payload["address"],
		Chain:   payload["chain"],
	}
	if err = finance.FindBalanceByAddress(&req, &resp); err != nil {
		return
	}
	tx := db.BeginTx()
	defer tx.DbRollback()
	var address model.Address
	if err = tx.
		Where("address = ?", payload["address"]).
		Where("coin = ?", payload["chain"]).
		Where("id = ?", payload["id"]).
		First(&address).Error; err != nil {
		return
	}
	address.Balance = decimal.NewFromFloat(resp.GetBalance())
	tx.Save(&address)
	tx.DbCommit()
	return
}

func (worker *AddressCheck) checkAddressFromBlockChain(payload map[string]string) (err error) {
	resp, err := blockChain.FindBalanceByAddress(payload["address"])
	if err != nil {
		return ErrMaxRequest
	}
	tx := db.BeginTx()
	defer tx.DbRollback()
	var address model.Address
	if err = tx.
		Where("address = ?", payload["address"]).
		Where("coin = ?", payload["chain"]).
		Where("id = ?", payload["id"]).
		First(&address).Error; err != nil {
		return
	}
	v := (*resp)[payload["address"]].(map[string]interface{})
	address.Balance = decimal.NewFromFloat(v["final_balance"].(float64)).Div(decimal.New(1, 8))
	address.DataSource = "BlockChain"
	tx.Save(&address)
	tx.DbCommit()
	return
}

func (worker *AddressCheck) checkAddressFromChainInfo(payload map[string]string) (err error) {
	resp, err := chainInfo.FindBalanceByAddress(payload["address"])
	if err != nil {
		return ErrMaxRequest
	}
	tx := db.BeginTx()
	defer tx.DbRollback()
	var address model.Address
	if err = tx.
		Where("address = ?", payload["address"]).
		Where("coin = ?", payload["chain"]).
		Where("id = ?", payload["id"]).
		First(&address).Error; err != nil {
		return
	}
	address.Balance, _ = decimal.NewFromString(resp.GetData().(map[string]interface{})["balance"].(string))
	address.Balance = address.Balance.Div(decimal.New(1, 8))
	address.DataSource = "ChainInfo"
	tx.Save(&address)
	tx.DbCommit()
	return
}

func (worker *AddressCheck) checkAddressFromBtcCom(payload map[string]string) (err error) {
	var resp btcCom.BaseResp
	if err = btcCom.FindBalanceByAddress(payload["address"], &resp); err != nil {
		return ErrMaxRequest
	}
	tx := db.BeginTx()
	defer tx.DbRollback()
	var address model.Address
	if err = tx.
		Where("address = ?", payload["address"]).
		Where("coin = ?", payload["chain"]).
		Where("id = ?", payload["id"]).
		First(&address).Error; err != nil {
		return
	}
	address.Balance = decimal.NewFromFloat(resp.Data["balance"].(float64))
	address.Balance = address.Balance.Div(decimal.New(1, 8))
	address.DataSource = "BtcCom"
	tx.Save(&address)
	tx.DbCommit()
	return
}

func (worker *AddressCheck) checkAddressFromBitGo(payload map[string]string) (err error) {
	var resp bitGo.BaseResp
	if err = bitGo.FindBalanceByAddress(payload["address"], &resp); err != nil {
		return ErrMaxRequest
	}
	tx := db.BeginTx()
	defer tx.DbRollback()
	var address model.Address
	if err = tx.
		Where("address = ?", payload["address"]).
		Where("coin = ?", payload["chain"]).
		Where("id = ?", payload["id"]).
		First(&address).Error; err != nil {
		return
	}
	address.Balance = decimal.New(int64(resp.Balance), -8)
	address.DataSource = "BitGo"
	tx.Save(&address)
	tx.DbCommit()
	return
}
