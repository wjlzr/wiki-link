package model

type BaseReq struct {
	Chain string `json:"chain" form:"chain" binding:"required"`
}

//链
type ChainInfoReq struct {
	BaseReq
}

//交易
type TransactionSearchReq struct {
	BaseReq
	Info string `json:"info" form:"info" binding:"required"`
	Pagination
}

//
type ChainCommonReq struct {
	BaseReq
	Pagination
}

// 富豪查询
type RicherReq struct {
	BaseReq
	Pagination
}

type RichersStat struct {
	BaseReq
	TokenId string `json:"tokenId" form:"tokenId"`
}

type BlockInfoReq struct {
	BaseReq
	Info string `json:"info" form:"info" binding:"required"`
}

type BlockPoolReq struct {
	ChainCommonReq
	Type string `json:"type" form:"type" `
	Pool string `json:"pool" form:"pool" `
}

type Transactions struct {
	ChainCommonReq
	Address string `json:"address" form:"address" binding:"required"`
}
