package chainInfo

// {
//   success: true,
//   error: "",
//   data: {
//     balance: "0",
//     totalReceived: "1380821864",
//     totalSent: "1380821864",
//     unconfirmed: 0,
//     transactionCount: 945,
//     firstAppearance: 1590667876,
//     lastAppearance: 1620595140,
//     tag: ""
//   }
// }
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

func (base BaseResp) TotalCount() int {
	data := base.GetData()
	if data == nil {
		return 0
	}
	return data.(map[string]interface{})["totalCount"].(int)
}

func (base BaseResp) Transactions() []interface{} {
	return base.GetData().(map[string]interface{})["transactions"].([]interface{})
}

func (base BaseResp) Inputs() []AddressInfo {
	return base.GetData().(map[string]interface{})["inputs"].([]AddressInfo)
}

func (base BaseResp) Outputs() []AddressInfo {
	return base.GetData().(map[string]interface{})["outputs"].([]AddressInfo)
}

func (base BaseResp) Timestamp() int64 {
	return base.GetData().(map[string]interface{})["timestamp"].(int64)
}

type AddressInfo struct {
	Address string `json:"address"`
	Tag     string `json:"tag"`
}
