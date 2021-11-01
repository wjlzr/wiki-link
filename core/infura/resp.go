package infura

type result interface {
	GetCode() int
	GetMsg() string
	GetData() interface{}
	GetErrorCode() int
}

type BaseResp struct {
	Id      int    `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  string `json:"result"`
	Error   struct {
		Code int `json:"code"`
	} `json:"error"`
}

func (base BaseResp) GetCode() int {
	return 1
}

func (base BaseResp) GetMsg() string {
	return ""
}

func (base BaseResp) GetData() interface{} {
	return base.Result
}
func (base BaseResp) GetErrorCode() int {
	return base.Error.Code
}
