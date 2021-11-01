package address

import (
	"wiki-link/core/log"
	"wiki-link/core/oklink"
	"wiki-link/handler/util"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetAnalysisByAddress(c *gin.Context) {
	var req oklink.AnalysisReq
	if err := c.ShouldBindQuery(&req); err != nil {
		log.Logger().Error("chain info param error:", zap.Error(err))
		util.ResponseErrorJson(c, 1010001, util.GetLang(c), err)
		return
	}
	resp, err := oklink.AnalysisByAddress(req.Chain, &req)
	if err != nil {
		log.Logger().Error("chain info error:", zap.Error(err))
		util.ResponseErrorJson(c, 2000001, util.GetLang(c), err)
		return
	}
	util.ResponseJson(c, "", resp.GetData())
}

func GetContractByAddress(c *gin.Context) {
	var req oklink.ChainAndAddressReq
	if err := c.ShouldBindQuery(&req); err != nil {
		log.Logger().Error("chain info param error:", zap.Error(err))
		util.ResponseErrorJson(c, 1010001, util.GetLang(c), err)
		return
	}
	resp, err := oklink.ContractByAddress(req.Chain, req.Address)
	if err != nil {
		log.Logger().Error("chain info error:", zap.Error(err))
		util.ResponseErrorJson(c, 2000001, util.GetLang(c), err)
		return
	}
	util.ResponseJson(c, "", resp.GetData())
}

func GetTransfersByAddress(c *gin.Context) {
	var req oklink.AddressTransfers
	if err := c.ShouldBindQuery(&req); err != nil {
		log.Logger().Error("chain info param error:", zap.Error(err))
		util.ResponseErrorJson(c, 1010001, util.GetLang(c), err)
		return
	}
	resp, err := oklink.TransfersByAddress(req.Chain, &req)
	if err != nil {
		log.Logger().Error("chain info error:", zap.Error(err))
		util.ResponseErrorJson(c, 2000001, util.GetLang(c), err)
		return
	}
	util.ResponseJson(c, "", resp.GetData())
}

func GetInternalTxByAddress(c *gin.Context) {
	var req oklink.InternalTxReq
	if err := c.ShouldBindQuery(&req); err != nil {
		util.ResponseErrorJson(c, 1010001, util.GetLang(c), err)
		return
	}
	resp, err := oklink.InternalTx(req.Chain, &req)
	if err != nil {
		util.ResponseErrorJson(c, 2000001, util.GetLang(c), err)
		return
	}
	util.ResponseJson(c, "", resp.GetData())
}

func GetAnalysisAddressesEdgeTxs(c *gin.Context) {
	var req oklink.AnalysisAddressesEdgeTxsReq
	if err := c.ShouldBindQuery(&req); err != nil {
		util.ResponseErrorJson(c, 1010001, util.GetLang(c), err)
		return
	}
	resp, err := oklink.AnalysisAddressesEdgeTxs(req.Chain, &req)
	if err != nil {
		util.ResponseErrorJson(c, 2000001, util.GetLang(c), err)
		return
	}
	util.ResponseJson(c, "", resp.GetData())
}

//
func GetAddressState(c *gin.Context) {
	var req oklink.AddressStateReq
	if err := c.ShouldBindQuery(&req); err != nil {
		util.ResponseErrorJson(c, 1010001, util.GetLang(c), err)
		return
	}
	resp, err := oklink.AddressState(req.Chain, &req)
	if err != nil {
		util.ResponseErrorJson(c, 2000001, util.GetLang(c), err)
		return
	}
	util.ResponseJson(c, "", resp.GetData())
}

//
func Addressstatistic(c *gin.Context) {
	var req oklink.AddressstatisticReq

	if err := c.ShouldBindQuery(&req); err != nil {
		util.ResponseErrorJson(c, 1010001, util.GetLang(c), err)
		return
	}
	resp, err := oklink.Addressstatistic(&req)
	if err != nil {
		util.ResponseErrorJson(c, 2000001, util.GetLang(c), err)
		return
	}
	util.ResponseJson(c, "", resp.GetData())

}
