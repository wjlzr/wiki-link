package infura

type BaseReq struct {
	Id      int    `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
}

type AddressBalanceReq struct {
	Id      int      `json:"id"`
	Jsonrpc string   `json:"jsonrpc"`
	Method  string   `json:"method"`
	Params  []string `json:"params"`
}
