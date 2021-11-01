package transaction

import (
	"strings"

	"wiki-link/common"
	"wiki-link/core/address"
	"wiki-link/core/log"
	"wiki-link/core/oklink"
	"wiki-link/handler/model"
	"wiki-link/handler/util"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//交易查询
func TransactionSearch(c *gin.Context) {
	var req model.TransactionSearchReq

	if err := c.ShouldBindQuery(&req); err != nil {
		log.Logger().Error("transaction search param error:", zap.Error(err))
		util.ResponseErrorJson(c, 1010001, util.GetLang(c), err)
		return
	}

	//地址
	if isAddress := address.Validation(req.Chain, req.Info); isAddress {
		var resp oklink.Base
		if e := oklink.FindByAddress(req.Chain, req.Info, &resp); e == nil {
			util.ResponseJson(c, "", resp.Data)
			return
		} else {
			util.ResponseErrorJson(c, 2000001, util.GetLang(c), e)
			return
		}
	}

	//区块高度
	var transactionsReq oklink.TransactionsReq
	isBlockHeight := util.IsNumeric(req.Info)
	if isBlockHeight {
		transactionsReq.PageSize = req.PageSize
		transactionsReq.PageIndex = req.PageIndex
		transactionsReq.BlockHeight = req.Info
		var resp oklink.Base
		if err := oklink.TransactionsNewEdition(req.Chain, &transactionsReq, &resp); err != nil {
			log.Logger().Error("transaction search error:", zap.Error(err))
			util.ResponseErrorJson(c, 2000001, util.GetLang(c), err)
			return
		}
		util.ResponseJson(c, "", resp.GetData())
		return
	}

	//区块哈希
	var resp oklink.Base
	if err := oklink.GetTransactionsByHash(req.Chain, req.Info, &resp); err != nil {
		log.Logger().Error("transaction search error:", zap.Error(err))
		util.ResponseErrorJson(c, 2000001, util.GetLang(c), err)
		return
	}
	util.ResponseJson(c, "", resp.GetData())
}

// 通用查询
func SearchWithoutChain(c *gin.Context) {
	var req oklink.SearchReq

	if err := c.ShouldBindQuery(&req); err != nil {
		util.ResponseErrorJson(c, 1010001, util.GetLang(c), err)
		return
	}

	if req.Chain != "" {
		TransactionSearch(c)
		return
	}
	var re []interface{}

	//地址
	if chain := address.ValidationWithoutChain(req.Info); chain != "" {
		for _, coin := range common.AllVisitableCoins() {
			if strings.ToLower(coin.Symbol) != strings.ToLower(chain) {
				continue
			}
			var resp oklink.SimpleAddressResp
			if err := oklink.FindByAddress(chain, req.Info, &resp); err == nil {
				data := resp.Data
				data.SetCoin(coin)
				re = append(re, data)
			}
		}
		util.ResponseJson(c, "", re)
		return
	}

	//区块高度
	if isBlockHeight := util.IsNumeric(req.Info); isBlockHeight {
		var blockInfoReq = oklink.BlockInfoReq{Info: req.Info}
		for _, chain := range common.AllVisitableCoins() {
			var resp oklink.SimpleBlockInfoResp
			if err := oklink.GetBlockInfo(strings.ToLower(chain.Symbol), &blockInfoReq, &resp); err == nil {
				data := resp.Data
				data.SetCoin(chain)
				re = append(re, data)
			}
		}
		util.ResponseJson(c, "", re)
		return
	}

	//区块哈希
	for _, chain := range common.AllVisitableCoins() {
		var resp oklink.SimpleTransactionsByHashResp
		if err := oklink.GetTransactionsByHash(strings.ToLower(chain.Symbol), req.Info, &resp); err == nil {
			data := resp.Data
			data.SetCoin(chain)
			re = append(re, data)
		}
	}
	util.ResponseJson(c, "", re)
}

func GetTransaction(c *gin.Context) {
	var req model.ChainCommonReq
	var oReq oklink.TransactionsReq
	if err := c.ShouldBindQuery(&req); err != nil {
		log.Logger().Error("chain info param error:", zap.Error(err))
		util.ResponseErrorJson(c, 1010001, util.GetLang(c), err)
		return
	}
	oReq.PageSize = req.PageSize
	oReq.PageIndex = req.PageIndex
	var resp oklink.Base
	if err := oklink.TransactionsNewEdition(req.Chain, &oReq, &resp); err != nil {
		log.Logger().Error("chain info error:", zap.Error(err))
		util.ResponseErrorJson(c, 2000001, util.GetLang(c), err)
		return
	}
	util.ResponseJson(c, "", resp.GetData())
}

func TransactionPending(c *gin.Context) {
	req := oklink.TransactionsReq{Type: "pending"}
	if err := c.ShouldBindQuery(&req); err != nil {
		log.Logger().Error("chain info param error:", zap.Error(err))
		util.ResponseErrorJson(c, 1010001, util.GetLang(c), err)
		return
	}
	var resp oklink.Base
	if err := oklink.TransactionsNewEdition(req.Chain, &req, &resp); err != nil {
		log.Logger().Error("chain info error:", zap.Error(err))
		util.ResponseErrorJson(c, 2000001, util.GetLang(c), err)
		return
	}
	util.ResponseJson(c, "", resp.GetData())
}

func TransactionByAddress(c *gin.Context) {
	var oReq oklink.TransactionsReq
	if err := c.ShouldBindQuery(&oReq); err != nil {
		log.Logger().Error("chain info param error:", zap.Error(err))
		util.ResponseErrorJson(c, 1010001, util.GetLang(c), err)
		return
	}
	resp, err := oklink.FindTransactionsByAddress(oReq.Chain, &oReq)
	if err != nil {
		log.Logger().Error("chain info error:", zap.Error(err))
		util.ResponseErrorJson(c, 2000001, util.GetLang(c), err)
		return
	}
	util.ResponseJson(c, "", resp.GetData())
}

func GetTransactionLogs(c *gin.Context) {
	var req oklink.TransactionsLogsReq
	if err := c.ShouldBindQuery(&req); err != nil {
		log.Logger().Error("chain info param error:", zap.Error(err))
		util.ResponseErrorJson(c, 1010001, util.GetLang(c), err)
		return
	}
	resp, err := oklink.TransactionsLogs(req.Chain, &req)
	if err != nil {
		log.Logger().Error("chain info error:", zap.Error(err))
		util.ResponseErrorJson(c, 2000001, util.GetLang(c), err)
		return
	}
	util.ResponseJson(c, "", resp.GetData())
}
