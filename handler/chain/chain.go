package chain

import (
	"sort"
	"strings"

	"wiki-link/common"
	"wiki-link/core/coinMarketCap"
	"wiki-link/core/log"
	"wiki-link/core/oklink"
	"wiki-link/handler/model"
	"wiki-link/handler/util"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//查询一个区块链的详情信息
func GetChainInfo(c *gin.Context) {
	var req model.ChainInfoReq

	if err := c.ShouldBindQuery(&req); err != nil {
		log.Logger().Error("chain info param error:", zap.Error(err))
		util.ResponseErrorJson(c, 1010001, util.GetLang(c), err)
		return
	}

	resp, err := oklink.ChainInfo(req.Chain)
	if err != nil {
		log.Logger().Error("chain info error:", zap.Error(err))
		util.ResponseErrorJson(c, 2000001, util.GetLang(c), err)
		return
	}
	resp.SetIcon()
	util.ResponseJson(c, "", resp.Data)
}

//查询支持的区块链的汇总信息。
func GetChainSummary(c *gin.Context) {
	var req oklink.SummaryInfoReq
	c.ShouldBindQuery(&req)
	resp, err := oklink.SummaryInfo()
	if err != nil {
		log.Logger().Error("chain summary error:", zap.Error(err))
		util.ResponseErrorJson(c, 2000001, util.GetLang(c), err)
		return
	}

	//
	var mapCoins map[string]string = make(map[string]string, 0)
	if req.Coins != "" {
		req.Coins = strings.ToUpper(req.Coins)
		cs := strings.Split(req.Coins, ",")
		for _, c := range cs {
			mapCoins[c] = c
		}
	}

	var coins []oklink.Summary
	for _, d := range resp.Data {
		coin := common.AllCoins[d.Symbol]
		if coin != nil && (mapCoins[d.Symbol] != "" || coin.Visitable) {
			d.SetCoin(coin)
			coins = append(coins, d)
		}
	}
	//从小到大排序
	sort.SliceStable(coins, func(i, j int) bool {
		if coins[i].Index < coins[j].Index {
			return true
		}
		return false
	})

	util.ResponseJson(c, "", coins)
}

func GetChainRank(c *gin.Context) {
	var req coinMarketCap.CurrencyListReq
	if err := c.ShouldBindQuery(&req); err != nil {
		util.ResponseErrorJson(c, 1010001, util.GetLang(c), err)
		return
	}
	resp, err := coinMarketCap.CurrencyList(&req)
	if err != nil {
		util.ResponseErrorJson(c, 2000001, util.GetLang(c), err)
		return
	}
	util.ResponseJson(c, "", resp.Data)
}

//查询链的基本统计信息
func GetChainCommonStatistic(c *gin.Context) {
	var req oklink.ChainPagination

	if err := c.ShouldBindQuery(&req); err != nil {
		log.Logger().Error("chain common statistic param error:", zap.Error(err))
		util.ResponseErrorJson(c, 1010001, util.GetLang(c), err)
		return
	}

	resp, err := oklink.StatisticCommon(req.Chain, &req)
	if err != nil {
		log.Logger().Error("chain common statistic error:", zap.Error(err))
		util.ResponseErrorJson(c, 2000001, util.GetLang(c), err)
		return
	}

	util.ResponseJson(c, "", resp.GetData())
}
