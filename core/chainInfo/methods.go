package chainInfo

var (
	ChainInfoUrl = map[string]string{
		"findTransaction":         "https://api.chain.info/blockbook/v1/api/tx/%v",
		"findTransactionsByBlock": "https://api.chain.info/blockbook/v1/api/block-txs/%v",
		"findBalanceByAddress":    "https://api.chain.info/blockbook/v1/api/address/%v",
	}
)
