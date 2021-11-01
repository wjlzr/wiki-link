package bitGo

// {
//   "balance": 138048685021,
//   "confirmedBalance": 138048685021,
//   "unconfirmedSends": 0,
//   "unconfirmedReceives": 0,
//   "spendableBalance": 138048685021,
//   "sent": 342383427327833,
//   "received": 342521476012854,
//   "address": "17A16QmavnUfCW11DAApiJxp7ARnxN5pGX"
// }
type result interface {
	GetCode() int
	GetMsg() string
	GetData() interface{}
	GetErrorCode() int
}

type BaseResp struct {
	Balance int `json:"balance"`
}

func (base BaseResp) GetCode() int {
	return 1
}

func (base BaseResp) GetMsg() string {
	return ""
}

func (base BaseResp) GetData() interface{} {
	return base
}

func (base BaseResp) GetErrorCode() int {
	return 0
}

func (base BaseResp) GetBalance() int {
	return base.Balance
}
