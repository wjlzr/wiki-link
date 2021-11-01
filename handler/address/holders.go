package address

import (
	"wiki-link/core/log"
	"wiki-link/core/oklink"
	"wiki-link/handler/util"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetHolders(c *gin.Context) {
	var req oklink.PageReq
	if err := c.ShouldBindQuery(&req); err != nil {
		log.Logger().Error("chain info param error:", zap.Error(err))
		util.ResponseErrorJson(c, 1010001, util.GetLang(c), err)
		return
	}
	var resp oklink.Base
	if err := oklink.AddressHolders(req.Chain, &req, &resp); err != nil {
		log.Logger().Error("chain info error:", zap.Error(err))
		util.ResponseErrorJson(c, 2000001, util.GetLang(c), err)
		return
	}
	util.ResponseJson(c, "", resp.GetData())
}

func GetTransactionsByClassfy(c *gin.Context) {
	var req oklink.PageReq
	if err := c.ShouldBindQuery(&req); err != nil {
		log.Logger().Error("chain info param error:", zap.Error(err))
		util.ResponseErrorJson(c, 1010001, util.GetLang(c), err)
		return
	}
	resp, err := oklink.TransactionsByClassfy(req.Chain, &req)
	if err != nil {
		log.Logger().Error("chain info error:", zap.Error(err))
		util.ResponseErrorJson(c, 2000001, util.GetLang(c), err)
		return
	}
	util.ResponseJson(c, "", resp.GetData())
}
