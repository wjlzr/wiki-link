package blockHeight

import (
	"strconv"

	"github.com/oldfritter/sidekiq-go"

	"wiki-link/common"
	"wiki-link/core/oklink"
	"wiki-link/db"
	hModel "wiki-link/handler/model"
	"wiki-link/model"
)

var (
	oKLinkBlockHeightCheck, chainInfoBlockHeightCheck sidekiq.WorkerI

	lastestEthBlockHeight int
	ethA                  = "eth:blockHeight:a"
	ethB                  = "eth:blockHeight:b"

	lastestBtcBlockHeight int
	btcA                  = "btc:blockHeight:a"
	btcB                  = "btc:blockHeight:b"
)

func InitWorker() {
	for _, w := range common.AllWorkers {
		if w.Name == "OKLinkBlockHeightCheck" {
			oKLinkBlockHeightCheck = common.AllWorkerIs[w.Name](&w)
			oKLinkBlockHeightCheck.SetClient(db.RedisClient())
		}
	}
}

func ReRunErrors() {
	oKLinkBlockHeightCheck.ReRunErrors()
}

func CheckEthBlockHeight() {
	if oKLinkBlockHeightCheck.GetMaxQuery() < oKLinkBlockHeightCheck.GetQuerySize() {
		return
	}

	client := db.RedisClient()

	// 0----------<-a-------b->-------c
	var a, b int

	// 设置初始点
	if client.Get(ethB).Val() == "" {
		client.Set(ethB, 12281877, 0)
		b = 12281877
	} else {
		b, _ = strconv.Atoi(client.Get(ethB).Val())
	bFor:
		for i := 1; i < 10; i++ {
			if b+i > lastestEthBlockHeight {
				break bFor
			}
			var count int64
			if db.MainDb.Model(&model.BlockHeight{}).Where("coin = ?", "eth").Where("number = ?", b+i).Count(&count); count == int64(0) {
				oKLinkPriority("eth", strconv.Itoa(b+i), "right")
			}
			client.Set(ethB, b+i, 0)
		}
	}

	if client.Get(ethA).Val() == "" {
		client.Set(ethA, 12281877, 0)
		a = 12281877
	} else {
		a, _ = strconv.Atoi(client.Get(ethA).Val())
	aFor:
		for i := 1; i < 50; i++ {
			if a-i > 0 {
				var count int64
				if db.MainDb.Model(&model.BlockHeight{}).Where("coin = ?", "eth").Where("number = ?", a-i).Count(&count); count == int64(0) {
					oKLinkPerform("eth", strconv.Itoa(a-i), "left")
				}
				client.Set(ethA, a-i, 0)
			} else {
				break aFor
			}
		}
	}

}

func GetEthLastestBlockHeight() {
	var req = hModel.ChainInfoReq{hModel.BaseReq{Chain: "eth"}}
	resp, _ := oklink.ChainInfo(req.Chain)
	lastestEthBlockHeight = int(resp.Data.Block.Height)
}

func CheckBtcBlockHeight() {
	if oKLinkBlockHeightCheck.GetMaxQuery() < oKLinkBlockHeightCheck.GetQuerySize() {
		return
	}

	client := db.RedisClient()

	// 0----------<-a-------b->-------c
	var a, b int

	// 设置初始点
	if client.Get(btcB).Val() == "" {
		client.Set(btcB, 684267, 0)
		b = 684267
	} else {
		b, _ = strconv.Atoi(client.Get(btcB).Val())
	bFor:
		for i := 1; i < 10; i++ {
			if b+i > lastestBtcBlockHeight {
				break bFor
			}
			var count int64
			if db.MainDb.Model(&model.BlockHeight{}).Where("coin = ?", "btc").Where("number = ?", b+i).Count(&count); count == int64(0) {
				oKLinkPriority("btc", strconv.Itoa(b+i), "right")
				// chainInfoPriority("btc", strconv.Itoa(b+i), "right")
			}
			client.Set(btcB, b+i, 0)
		}
	}

	if client.Get(btcA).Val() == "" {
		client.Set(btcA, 684267, 0)
		a = 684267
	} else {
		a, _ = strconv.Atoi(client.Get(btcA).Val())
	aFor:
		for i := 1; i < 50; i++ {
			if a-i > 0 {
				var count int64
				if db.MainDb.Model(&model.BlockHeight{}).Where("coin = ?", "btc").Where("number = ?", a-i).Count(&count); count == int64(0) {
					oKLinkPerform("btc", strconv.Itoa(a-i), "left")
					// chainInfoPerform("btc", strconv.Itoa(a-i), "left")
				}
				client.Set(btcA, a-i, 0)
			} else {
				break aFor
			}
		}
	}
}

func GetBtcLastestBlockHeight() {
	var req = hModel.ChainInfoReq{hModel.BaseReq{Chain: "btc"}}
	resp, _ := oklink.ChainInfo(req.Chain)
	lastestBtcBlockHeight = int(resp.Data.Block.Height)
}

func oKLinkPerform(coin, blockHeight, toward string) {
	oKLinkBlockHeightCheck.Perform(map[string]string{"blockHeight": blockHeight, "chain": coin, "toward": toward})
}

func oKLinkPriority(coin, blockHeight, toward string) {
	oKLinkBlockHeightCheck.Priority(map[string]string{"blockHeight": blockHeight, "chain": coin, "toward": toward})
}

func chainInfoPerform(coin, blockHeight, toward string) {
	chainInfoBlockHeightCheck.Perform(map[string]string{"blockHeight": blockHeight, "chain": coin, "toward": toward})
}

func chainInfoPriority(coin, blockHeight, toward string) {
	chainInfoBlockHeightCheck.Priority(map[string]string{"blockHeight": blockHeight, "chain": coin, "toward": toward})
}
