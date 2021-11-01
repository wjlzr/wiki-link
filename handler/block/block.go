package block

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"wiki-link/common"
	"wiki-link/core/log"
	"wiki-link/core/oklink"
	"wiki-link/handler/model"
	"wiki-link/handler/util"
)

func GetRichers(c *gin.Context) {
	var req oklink.ChainPagination
	if err := c.ShouldBindQuery(&req); err != nil {
		log.Logger().Error("chain info param error:", zap.Error(err))
		util.ResponseErrorJson(c, 1010001, util.GetLang(c), err)
		return
	}
	resp, err := oklink.FindRichers(req.Chain, &req)
	if err != nil {
		log.Logger().Error("chain info error:", zap.Error(err))
		util.ResponseErrorJson(c, 2000001, util.GetLang(c), err)
		return
	}
	util.ResponseJson(c, "", resp.GetData())
}

func GetStatisticRichersStat(c *gin.Context) {
	var req model.RichersStat
	if err := c.ShouldBindQuery(&req); err != nil {
		log.Logger().Error("chain info param error:", zap.Error(err))
		util.ResponseErrorJson(c, 1010001, util.GetLang(c), err)
		return
	}
	resp, err := oklink.StatisticRichersStat(req.Chain, req.TokenId)
	if err != nil {
		log.Logger().Error("chain info error:", zap.Error(err))
		util.ResponseErrorJson(c, 2000001, util.GetLang(c), err)
		return
	}
	util.ResponseJson(c, "", resp.Data)
}

func GetBlockInfo(c *gin.Context) {
	var req oklink.BlockInfoReq
	if err := c.ShouldBindQuery(&req); err != nil {
		log.Logger().Error("chain info param error:", zap.Error(err))
		util.ResponseErrorJson(c, 1010001, util.GetLang(c), err)
		return
	}
	var resp oklink.Base
	if err := oklink.GetBlockInfo(req.Chain, &req, &resp); err != nil {
		log.Logger().Error("chain info error:", zap.Error(err))
		util.ResponseErrorJson(c, 2000001, util.GetLang(c), err)
		return
	}
	util.ResponseJson(c, "", resp.GetData())
}

func GetBlocks(c *gin.Context) {
	var req oklink.BlocksReq
	if err := c.ShouldBindQuery(&req); err != nil {
		log.Logger().Error("chain info param error:", zap.Error(err))
		util.ResponseErrorJson(c, 1010001, util.GetLang(c), err)
		return
	}
	resp, err := oklink.Blocks(req.Chain, &req)
	if err != nil {
		log.Logger().Error("chain info error:", zap.Error(err))
		util.ResponseErrorJson(c, 2000001, util.GetLang(c), err)
		return
	}
	util.ResponseJson(c, "", resp.GetData())
}

func GetUncleBlocks(c *gin.Context) {
	var req oklink.UncleBlocksReq
	if err := c.ShouldBindQuery(&req); err != nil {
		log.Logger().Error("chain info param error:", zap.Error(err))
		util.ResponseErrorJson(c, 1010001, util.GetLang(c), err)
		return
	}
	resp, err := oklink.UncleBlocks(req.Chain, &req)
	if err != nil {
		log.Logger().Error("chain info error:", zap.Error(err))
		util.ResponseErrorJson(c, 2000001, util.GetLang(c), err)
		return
	}
	util.ResponseJson(c, "", resp.GetData())
}

func GetUncleBlock(c *gin.Context) {
	var req oklink.UncleBlockReq
	if err := c.ShouldBindQuery(&req); err != nil {
		log.Logger().Error("chain info param error:", zap.Error(err))
		util.ResponseErrorJson(c, 1010001, util.GetLang(c), err)
		return
	}
	resp, err := oklink.UncleBlock(req.Chain, &req)
	if err != nil {
		log.Logger().Error("chain info error:", zap.Error(err))
		util.ResponseErrorJson(c, 2000001, util.GetLang(c), err)
		return
	}
	util.ResponseJson(c, "", resp.GetData())
}

func GetBlockPool(c *gin.Context) {
	var req oklink.BlocksReq
	if err := c.ShouldBindQuery(&req); err != nil {
		log.Logger().Error("chain info param error:", zap.Error(err))
		util.ResponseErrorJson(c, 1010001, util.GetLang(c), err)
		return
	}
	req.Type = "pool"
	resp, err := oklink.Blocks(req.Chain, &req)
	if err != nil {
		log.Logger().Error("chain info error:", zap.Error(err))
		util.ResponseErrorJson(c, 2000001, util.GetLang(c), err)
		return
	}
	util.ResponseJson(c, "", resp.Data)
}

func GetTransfers(c *gin.Context) {
	var req oklink.TransfersReq
	if err := c.ShouldBindQuery(&req); err != nil {
		log.Logger().Error("chain info param error:", zap.Error(err))
		util.ResponseErrorJson(c, 1010001, util.GetLang(c), err)
		return
	}
	resp, err := oklink.Transfers(req.Chain, &req)
	if err != nil {
		log.Logger().Error("chain info error:", zap.Error(err))
		util.ResponseErrorJson(c, 2000001, util.GetLang(c), err)
		return
	}
	util.ResponseJson(c, "", resp.GetData())
}

func GetTransactionsNoRestrict(c *gin.Context) {

	var req oklink.PageReq
	if err := c.ShouldBindQuery(&req); err != nil {
		log.Logger().Error("chain info param error:", zap.Error(err))
		util.ResponseErrorJson(c, 1010001, util.GetLang(c), err)
		return
	}
	resp, err := oklink.TransactionsNoRestrict(req.Chain, &req)
	if err != nil {
		log.Logger().Error("chain info error:", zap.Error(err))
		util.ResponseErrorJson(c, 2000001, util.GetLang(c), err)
		return
	}

	if util.GeClient(c) == "app" {
		switch req.Chain {
		case common.BTC, common.LTC:
			// 价格换算
			for k, v := range resp.Data.Hits {
				if v.Fee > 0 {
					resp.Data.Hits[k].Fee = v.Fee / 100000000
				}
			}
		}
	}

	util.ResponseJson(c, "", resp.Data)
}

func GetInternalTransactions(c *gin.Context) {
	var req oklink.InternalTransactions
	if err := c.ShouldBindQuery(&req); err != nil {
		log.Logger().Error("chain info param error:", zap.Error(err))
		util.ResponseErrorJson(c, 1010001, util.GetLang(c), err)
		return
	}
	resp, err := oklink.TransactionsInternal(req.Chain, &req)
	if err != nil {
		log.Logger().Error("chain info error:", zap.Error(err))
		util.ResponseErrorJson(c, 2000001, util.GetLang(c), err)
		return
	}
	util.ResponseJson(c, "", resp.GetData())
}

func GetMatch(c *gin.Context) {
	var req oklink.MatchReq
	if err := c.ShouldBindQuery(&req); err != nil {
		log.Logger().Error("chain info param error:", zap.Error(err))
		util.ResponseErrorJson(c, 1010001, util.GetLang(c), err)
		return
	}
	resp, err := oklink.MatchTokens(req.Chain, req.Q)
	if err != nil {
		log.Logger().Error("chain info error:", zap.Error(err))
		util.ResponseErrorJson(c, 2000001, util.GetLang(c), err)
		return
	}
	util.ResponseJson(c, "", resp.GetData())
}

func GetAnalysisTransactions(c *gin.Context) {
	var req oklink.AnalysisTransactionsReq
	if err := c.ShouldBindQuery(&req); err != nil {
		log.Logger().Error("chain info param error:", zap.Error(err))
		util.ResponseErrorJson(c, 1010001, util.GetLang(c), err)
		return
	}
	resp, err := oklink.AnalysisTransactions(req.Chain, &req)
	if err != nil {
		log.Logger().Error("chain info error:", zap.Error(err))
		util.ResponseErrorJson(c, 2000001, util.GetLang(c), err)
		return
	}
	if c.GetHeader("x-client") == "app" {
		data := resp.GetData().([]interface{})
		first := data[0].(map[string]interface{})
		inputs := first["inputs"].([]interface{})
		outputs := first["outputs"].([]interface{})
		if len(inputs)+len(outputs) > 100 {
			util.ResponseErrorJson(c, 2000002, util.GetLang(c), errors.New("too much data"))
			return
		}
	}
	util.ResponseJson(c, "", resp.GetData())
}
