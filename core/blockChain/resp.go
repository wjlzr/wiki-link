package blockChain

type result interface {
	GetCode() int
	GetMsg() string
	GetData() interface{}
}

type BaseResp map[string]interface{}

func (base BaseResp) GetCode() int {
	return 1
}

func (base BaseResp) GetMsg() string {
	return ""
}

func (base BaseResp) GetData() interface{} {
	return base
}
