package coinMarketCap

type Pagination struct {
	PageSize  int `form:"pageSize"`
	PageIndex int `form:"pageIndex"`
	Start     int `json:"start"`
	Limit     int `json:"limit"`
}

func (p *Pagination) InitPagination() {
	if p.Limit == 0 {
		if p.PageSize <= 0 {
			p.Limit = 20
		} else {
			p.Limit = p.PageSize
		}
	}
	if p.Start == 0 {
		if p.PageIndex < 1 {
			p.Start = 1
		} else if p.PageIndex > 1 {
			p.Start = (p.PageIndex - 1) * p.PageSize
		}
	}
}

func (pagination Pagination) GetPagination() interface{} {
	return pagination
}

type CurrencyListReq struct {
	Pagination
	SortBy   string `json:"sortBy" form:"sortBy"`
	SortType string `json:"sortType" form:"sortType"`
	Convert  string `json:"convert" form:"convert"`
}

func (req *CurrencyListReq) SetDefault() {
	if req.SortBy == "" {
		req.SortBy = "volume_24h"
	}
	if req.SortType == "" {
		req.SortType = "desc"
	}
	if req.Convert == "" {
		req.Convert = "USD"
	}
}
