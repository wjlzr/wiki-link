package btcCom

// {
//   "data": {
//     "address": "17A16QmavnUfCW11DAApiJxp7ARnxN5pGX",
//     "received": 13456600053996106,
//     "sent": 13456462082976426,
//     "balance": 137971019680,
//     "tx_count": 433448,
//     "unconfirmed_tx_count": 0,
//     "unconfirmed_received": 0,
//     "unconfirmed_sent": 0,
//     "unspent_tx_count": 174
//   },
//   "err_code": 0,
//   "err_no": 0,
//   "message": "success",
//   "status": "success"
// }
type result interface {
	GetCode() int
	GetMsg() string
	GetData() interface{}
	GetErrorCode() int
}

type BaseResp struct {
	Data    map[string]interface{} `json:"data"`
	ErrCode int                    `json:"err_code"`
	ErrNo   int                    `json:"err_no"`
	Message string                 `json:"message"`
	Status  string                 `json:"status"`
}

func (base BaseResp) GetCode() int {
	return 1
}

func (base BaseResp) GetMsg() string {
	return base.Message
}

func (base BaseResp) GetData() interface{} {
	return base.Data
}
func (base BaseResp) GetErrorCode() int {
	return base.ErrCode
}
