package http

var (
	BlockChainUrl = map[string]string{
		"findBalanceByAddress":    "https://blockchain.info/balance",
		"findTransactionsByBlock": "https://api.blockchain.info/haskoin-store/btc/block/%v",
	}
)
