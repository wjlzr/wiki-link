package bitGo

import (
	"wiki-link/core/http"
)

//根据地址查询地址的余额
//address
func FindBalanceByAddress(address string, result http.Result) (err error) {
	err = Invoke("findBalanceByAddress", result, nil, address)
	return
}
