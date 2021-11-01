package coinMarketCap

import "strconv"

type result interface {
	GetCode() int
	GetMsg() string
	GetData() interface{}
}

type BaseResp struct {
	Data struct {
		CryptoCurrencyList []interface{} `json:"cryptoCurrencyList"`
	} `json:"data"`
	Status struct {
		ErrorCode    string `json:"error_code"`
		ErrorMessage string `json:"error_message"`
	} `json:"status"`
}

func (base BaseResp) GetCode() int {
	code, _ := strconv.Atoi(base.Status.ErrorCode)
	return code
}

func (base BaseResp) GetMsg() string {
	return base.Status.ErrorMessage
}

func (base BaseResp) GetData() interface{} {
	return base.Data.CryptoCurrencyList
}
