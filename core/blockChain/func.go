package blockChain

import (
	"wiki-link/core/http"
	"wiki-link/core/util"
)

func FindBalanceByAddress(address string) (*BaseResp, error) {
	var result BaseResp
	query := map[string]interface{}{"active": address}
	err := http.BlockChainInvoke("findBalanceByAddress", &result, query)
	return &result, err
}

func FindTransactionsByBlock(req *TransactionsByBlockReq) (*BaseResp, error) {
	var result BaseResp
	var query map[string]interface{}
	if req != nil {
		req.InitPagination()
		query = util.Paged2MapJson(*req)
	}
	err := http.BlockChainInvoke("findTransactionsByBlock", &result, query, req.BlockHash)
	return &result, err
}
