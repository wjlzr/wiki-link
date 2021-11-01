package chainInfo

type Pagination struct {
	PageSize  int `form:"pageSize"`
	PageIndex int `form:"pageIndex"`
	Offset    int `json:"offset"`
	Limit     int `json:"limit"`
}

func (p *Pagination) InitPagination() {
	if p.PageSize <= 0 {
		p.Limit = 20
	} else {
		p.Limit = p.PageSize
	}
	if p.PageIndex <= 1 {
		p.Offset = 0
	} else {
		p.Offset = (p.PageIndex - 1) * p.PageSize
	}
}

func (pagination Pagination) GetPagination() interface{} {
	return pagination
}

type TransactionsByBlockReq struct {
	Pagination
	BlockHeight string
}
