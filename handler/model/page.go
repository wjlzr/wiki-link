package model

//
type Pagination struct {
	PageSize  int `form:"pageSize" json:"pageSize"`
	PageIndex int `form:"pageIndex" json:"pageIndex"`
}
