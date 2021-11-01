package infura

//根据地址查询地址的余额
//chain    待查询区块链的符号小写
//address
func FindBalanceByAddress(address string) (*BaseResp, error) {
	req := AddressBalanceReq{
		Method: "eth_getBalance",
		Params: []string{address, "latest"},
	}
	req.Id = 1
	req.Jsonrpc = "2.0"
	var result BaseResp
	err := InfuraPost("addressInfo", &result, &req)
	return &result, err
}
