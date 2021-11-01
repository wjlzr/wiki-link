package router

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/newrelic/go-agent/v3/integrations/nrgin"
	"github.com/newrelic/go-agent/v3/newrelic"
	"go.uber.org/zap"

	"wiki-link/handler/address"
	"wiki-link/handler/block"
	"wiki-link/handler/chain"
	"wiki-link/handler/token"
	"wiki-link/handler/transaction"
	"wiki-link/middleware"
)

//路由配置
func RouterEngine(zapLogger *zap.Logger) *gin.Engine {

	engine := gin.New()
	if app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("wikilink-api"),
		newrelic.ConfigLicense("b96c746e9a168a6969315a21804178af1e00NRAL"),
	); err != nil {
		os.Exit(1)
	} else {
		engine.Use(nrgin.Middleware(app))
	}

	//gin中间件使用
	engine.Use(middleware.Cors())
	engine.Use(middleware.Secure())
	engine.Use(middleware.Language())
	engine.Use(middleware.CustomIntercept())
	engine.Use(middleware.Ginzap(zapLogger, time.RFC3339, true))
	engine.Use(middleware.RecoveryWithZap(zapLogger, true))

	apiGroup := engine.Group("api")
	apiV1Group := apiGroup.Group("v1")
	apiV1sGroup(apiV1Group)

	//区块
	blockGroup := apiV1Group.Group("block")
	blocksGroup(blockGroup)

	//链
	chainGroup := apiV1Group.Group("chain")
	chainsGroup(chainGroup)

	tokenGroup := apiV1Group.Group("token")
	tokensGroup(tokenGroup)

	addressGroup := apiV1Group.Group("address")
	addresssGroup(addressGroup)

	//交易
	transactionGroup := apiV1Group.Group("transaction")
	transactionsGroup(transactionGroup)

	// apiV2Group := engine.Group("v2")
	return engine
}

func apiV1sGroup(rg *gin.RouterGroup) {
	rg.GET("/richers", block.GetRichers)
	rg.GET("/richers/stat", block.GetStatisticRichersStat)

	rg.GET("/blocks", block.GetBlocks)
	rg.GET("/block", block.GetBlockInfo)
	rg.GET("/block/pool", block.GetBlockPool)
	rg.GET("/match", block.GetMatch)
}

//链
func chainsGroup(rg *gin.RouterGroup) {

	rg.GET("/rank", chain.GetChainRank)
	//链基本信息
	rg.GET("/info", chain.GetChainInfo)
	//查询支持区块链的汇总信息
	rg.GET("/summary", chain.GetChainSummary)
	//查询链的基本统计信息
	rg.GET("/common/statistic", chain.GetChainCommonStatistic)
}

func tokensGroup(rg *gin.RouterGroup) {
	rg.GET("/holders", token.GetHolders)
	rg.GET("/tokens", token.GetTokens)
	rg.GET("/match", token.GetTokensMatch)
}

func addresssGroup(rg *gin.RouterGroup) {
	rg.GET("/analysis", address.GetAnalysisByAddress)
	rg.GET("/analysis/addresses/edge/txs", address.GetAnalysisAddressesEdgeTxs)
	rg.GET("/internalTx", address.GetInternalTxByAddress)
	rg.GET("/contract", address.GetContractByAddress)
	rg.GET("/transfers", address.GetTransfersByAddress)
	rg.GET("/holders", address.GetHolders)
	rg.GET("/transactionsByClassfy", address.GetTransactionsByClassfy)
	rg.GET("/state", address.GetAddressState)
	rg.GET("/analysis/statistic", address.Addressstatistic)
}

//区块
func blocksGroup(rg *gin.RouterGroup) {
	rg.GET("/transfers", block.GetTransfers)
	rg.GET("/transactionsNoRestrict", block.GetTransactionsNoRestrict)
	rg.GET("/internalTransactions", block.GetInternalTransactions)
	rg.GET("/uncle", block.GetUncleBlocks)
	rg.GET("/uncle/info", block.GetUncleBlock)
	rg.GET("/analysis/transactions", block.GetAnalysisTransactions)
}

//交易
func transactionsGroup(rg *gin.RouterGroup) {
	//交易查找
	rg.GET("/search", transaction.TransactionSearch)
	rg.GET("/search/without/chain", transaction.SearchWithoutChain)

	rg.GET("/pending", transaction.TransactionPending)
	rg.GET("/lastest", transaction.GetTransaction)
	rg.GET("/address", transaction.TransactionByAddress)
	rg.GET("/logs", transaction.GetTransactionLogs)
}
