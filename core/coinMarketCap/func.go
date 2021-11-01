package coinMarketCap

import (
	"wiki-link/core/util"
)

func CurrencyList(req *CurrencyListReq) (*BaseResp, error) {
	var query map[string]interface{}
	if req != nil {
		req.SetDefault()
		req.InitPagination()
		query = util.Paged2MapJson(*req)
	}
	var result BaseResp
	err := Invoke("currencyList", &result, query)
	return &result, err
}
