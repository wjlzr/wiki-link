package http

var (
	OKLinkUrl = map[string]string{
		//token 模糊匹配接口
		"matchTokens": "https://www.oklink.com/api/explorer/v1/%v/matchTokens",
		//主链通用查询接口
		"chainSearch": "https://www.oklink.com/api/explorer/v1/%v/search",
		//查询链的详情
		"chainInfo": "https://www.oklink.com/api/explorer/v1/%v/info",
		//查询支持的区块链的汇总信息
		"summary": "https://www.oklink.com/api/explorer/v1/info/summary",
		//查询链的基本统计信息
		"statisticCommon": "https://www.oklink.com/api/explorer/v1/%v/statistic/common",
		//查询当前大额交易统计
		"statisticLargeTransferLastest": "https://www.oklink.com/api/explorer/v1/%v/statistic/large/transfer/lastest",
		//查询一个区块链的主链币或代币的持仓分布统计。
		"statisticRichersStat": "https://www.oklink.com/api/explorer/v1/%v/statistic/richers/stat",
		//查询一个区块链的区块列表,区块的排序默认为区块高度倒序
		"blocks": "https://www.oklink.com/api/explorer/v1/%v/blocks",
		//查询一个区块链的区块详情。
		"blocksInfo": "https://www.oklink.com/api/explorer/v1/%v/blocks/%v",
		//查询一个区块链的矿池出块列表，默认为区块高度倒序。
		"transactions": "https://www.oklink.com/api/explorer/v1/%v/transactions",
		//查询一个区块链的交易详情
		"transactionsByHash": "https://www.oklink.com/api/explorer/v1/%v/transactions/%v",
		//查询一个高度的区块的交易列表
		"transactionNoRestrict": "https://www.oklink.com/api/explorer/v1/%v/transactionNoRestrict",
		//查询一个区块链的未确认交易列表，接口存在问题
		"transactionsTrade": "https://www.oklink.com/api/explorer/v1/%v/transactions/trade",
		//查询一个地址的基本信息
		"findByAddress": "https://www.oklink.com/api/explorer/v1/%v/addresses/%v",
		//根据地址查询地址的余额
		"findBalanceByAddress": "https://www.oklink.com/api/explorer/v1/%v/addresses/balance/%v",
		// 查询一个地址的交易列表
		"findTransactionsByAddress": "https://www.oklink.com/api/explorer/v1/%v/addresses/%v/transactions",
		// 查询一个区块链的富豪地址排行信息
		"findRichers": "https://www.oklink.com/api/explorer/v1/%v/richers",

		// 以下为非文档提供接口
		// 叔块列表
		"uncleBlocks": "https://www.oklink.com/api/explorer/v1/%v/uncleBlocks",
		// 叔块详情
		"uncleBlock": "https://www.oklink.com/api/explorer/v1/%v/uncleBlocks/%v",
		// 查询地址分析
		"analysis": "https://www.oklink.com/api/explorer/v1/%v/analysis/addresses",
		// chain
		"transfers": "https://www.oklink.com/api/explorer/v1/%v/transfers",
		// chain, address
		"tokensHolders": "https://www.oklink.com/api/explorer/v1/%v/tokens/holders/%v",
		// chain, address
		"tokens": "https://www.oklink.com/api/explorer/v1/%v/tokens/%v",
		// chain
		"transactionsNoRestrict": "https://www.oklink.com/api/explorer/v1/%v/transactionsNoRestrict",
		// internalTransactions
		"internalTransactions": "https://www.oklink.com/api/explorer/v1/%v/internalTransactions",
		// contract
		"contract": "https://www.oklink.com/api/explorer/v1/%v/addresses/%v/contract",
		// address transfers
		"addressTransfers": "https://www.oklink.com/api/explorer/v1/%v/addresses/%v/transfers",
		// chain, address
		"addressHolders": "https://www.oklink.com/api/explorer/v1/%v/addresses/%v/holders",
		// chain, address
		"transactionsByClassfy": "https://www.oklink.com/api/explorer/v1/%v/addresses/%v/transactionsByClassfy",
		"search":                "https://www.oklink.com/api/explorer/v1/%v/search",
		"transactionsLogs":      "https://www.oklink.com/api/explorer/v1/%v/transactions/%v/logs",
		"tokensMatch":           "https://www.oklink.com/api/explorer/v1/%v/tokens/match/%v",
		// https://www.oklink.com/api/explorer/v1/btc/analysis/transactions?t=1620877954269
		"analysisTransactions": "https://www.oklink.com/api/explorer/v1/%v/analysis/transactions",
		// https://www.oklink.com/api/explorer/v1/eth/addresses/0xa1d8d972560c2f8144af871db508f0b0b10a3fbf/internalTx?t=1621824580734&offset=0&limit=20&filterValue=true
		"internalTx": "https://www.oklink.com/api/explorer/v1/%v/addresses/%v/internalTx",
		// https://www.oklink.com/api/explorer/v1/btc/analysis/addresses/edge/txs
		"analysisAddressesEdgeTxs": "https://www.oklink.com/api/explorer/v1/%v/analysis/addresses/edge/txs",
		// https://www.oklink.com/api/explorer/v1/btc/statistic/addresses/1Edu4yBtfAKwGGsQSa45euTSAG6A2Zbone/state
		"addressState": "https://www.oklink.com/api/explorer/v1/%v/statistic/addresses/%v/state",
		//
		"addressStatistic": "https://www.oklink.com/api/explorer/v1/eth/analysis/addresses/statistic",
	}
)
