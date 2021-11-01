package finance

import (
	"wiki-link/core/http"
	"wiki-link/core/util"
)

func FindBalanceByAddress(req *AddressReq, result http.Result) error {
	var query map[string]interface{}
	if req != nil {
		query = util.Struct2MapJson(*req)
	}
	err := Invoke("findBalanceByAddress", result, query)
	return err
}
