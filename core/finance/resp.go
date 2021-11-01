package finance

type result interface {
	GetCode() int
	GetMsg() string
	GetData() interface{}
}

type BaseResp struct {
	Success bool        `json:"success"`
	Error   string      `json:"error"`
	Data    interface{} `json:"data"`
}

func (base BaseResp) GetCode() int {
	return 1
}

func (base BaseResp) GetMsg() string {
	return base.Error
}

func (base BaseResp) GetData() interface{} {
	return base.Data
}

func (base BaseResp) GetBalance() float64 {
	data := base.GetData()
	if data == nil {
		return 0
	}
	return data.(map[string]interface{})["balance"].(float64)
}
