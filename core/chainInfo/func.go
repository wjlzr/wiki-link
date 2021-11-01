package chainInfo

import (
	"wiki-link/core/util"
)

func FindTransaction(tx string) (*BaseResp, error) {
	var result BaseResp
	err := ChainInfoInvoke("findTransaction", &result, nil, tx)
	return &result, err
}

// 暂不使用，减少接口访问数
func FindTransactionsByBlock(req *TransactionsByBlockReq) (*BaseResp, error) {
	var result BaseResp
	var query map[string]interface{}
	if req != nil {
		req.InitPagination()
		query = util.Paged2MapJson(*req)
	}
	err := ChainInfoInvoke("findTransactionsByBlock", &result, query, req.BlockHeight)
	return &result, err
}

func FindBalanceByAddress(address string) (*BaseResp, error) {
	var result BaseResp
	err := ChainInfoInvoke("findBalanceByAddress", &result, nil, address)
	return &result, err
}
