package oklink

type Pagination struct {
	PageSize  int `json:"-" form:"pageSize"`
	PageIndex int `json:"-" form:"pageIndex"`
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

type ChainPagination struct {
	Pagination
	Chain string `json:"chain" form:"chain"`
}

type MatchReq struct {
	Chain string `json:"chain" form:"chain"`
	Q     string `json:"q" form:"q"`
}

//块查询
type BlocksReq struct {
	ChainPagination
	Type string `json:"type" form:"type"`
	Pool string `json:"pool" form:"pool"`
	Sort string `json:"sort" form:"sort"`
}

type BlockInfoReq struct {
	ChainPagination
	Info string `json:"info" form:"info"`
}

type SummaryInfoReq struct {
	Coins string `json:"coins" form:"coins"`
}

//交易查询
type TransactionsReq struct {
	Pagination
	Chain       string `json:"chain" form:"chain"`
	BlockHash   string `json:"blockHash" form:"blockHash"`
	BlockHeight string `json:"blockHeight" form:"blockHeight"`
	Sort        string `json:"sort" form:"sort"`
	Type        string `json:"type" form:"type"`
	Tag         string `json:"tag" form:"tag"`
	Address     string `json:"address" form:"address"`
}

//
type BalanceByAddressResp struct {
	Base
	Data []BalanceData `json:"data"`
}

//
type BalanceData struct {
	Address    string  `json:"address"`
	Balance    float64 `json:"balance"`
	BalanceSat float64 `json:"balanceSat"`
}

type TransactionPendingReq struct {
	Pagination
	Sort string `json:"sort"`
	Type string `json:"type"`
}

type AnalysisReq struct {
	Chain                string `json:"chain" form:"chain"`
	SourceAddress        string `json:"sourceAddress" form:"sourceAddress"`
	TokenContractAddress string `json:"tokenContractAddress" form:"tokenContractAddress"`
	StartTime            int64  `json:"startTime" form:"startTime"`
	EndTime              int64  `json:"endTime" form:"endTime"`
	MinValue             string `json:"minValue" form:"minValue"`
	MaxValue             string `json:"maxValue" form:"maxValue"`
	ShowTagData          string `json:"showTagData" form:"showTagData"`
	Direction            string `json:"direction" form:"direction"`
	Depth                string `json:"depth" form:"depth"`
	Mixed                string `json:"mixed" form:"mixed"`
}

type InternalTransactions struct {
	ChainPagination
	BlockHeight string `json:"blockHeight" form:"blockHeight"`
	TranHash    string `json:"tranHash" form:"tranHash"`
}

type TransfersReq struct {
	ChainPagination
	BlockHeight          string `json:"blockHeight" form:"blockHeight"`
	TokenContractAddress string `json:"tokenContractAddress" form:"tokenContractAddress"`
	TranHash             string `json:"tranHash" form:"tranHash"`
}

type PageReq struct {
	ChainPagination
	Address      string `json:"address" form:"address"`
	Type         string `json:"type" form:"type"`
	BlockHeight  string `json:"blockHeight" form:"blockHeight"`
	TokenAddress string `json:"tokenAddress" form:"tokenAddress"`
}

type ChainAndAddressReq struct {
	Chain   string `json:"chain" form:"chain" binding:"required"`
	Address string `json:"address" form:"address" binding:"required"`
}

type AddressTransfers struct {
	Pagination
	Chain        string `json:"chain" form:"chain" binding:"required"`
	Address      string `json:"address" form:"address" binding:"required"`
	TokenType    string `json:"tokenType" form:"tokenType" binding:"required"`
	TokenAddress string `json:"tokenAddress" form:"tokenAddress"`
}

type UncleBlocksReq struct {
	ChainPagination
	Height int `json:"height" form:"blockHeight" binding:"required"`
}

type UncleBlockReq struct {
	Chain string `json:"chain" form:"chain"`
	Info  string `json:"info" form:"info"`
}

type TransactionsLogsReq struct {
	Chain string `json:"chain" form:"chain"`
	Info  string `json:"info" form:"info"`
}

type TokensMatchReq struct {
	Chain string `json:"chain" form:"chain"`
	Info  string `json:"info" form:"info"`
}

type AnalysisTransactionsReq struct {
	Chain    string `json:"-" form:"chain"`
	Txhash   string `json:"txhash" form:"txhash"`
	Depth    string `json:"depth" form:"depth"`
	MaxValue string `json:"maxValue" form:"maxValue"`
	MinValue string `json:"minValue" form:"minValue"`
}

type InternalTxReq struct {
	Pagination
	Chain       string `json:"-" form:"chain"`
	Address     string `json:"-" form:"address"`
	FilterValue string `json:"filterValue" form:"filterValue"`
}

type AnalysisAddressesEdgeTxsReq struct {
	Pagination
	Chain     string `json:"-" form:"chain"`
	RequestId string `json:"requestId" form:"requestId"`
	From      string `json:"from" form:"from"`
	To        string `json:"to" form:"to"`
	Sort      string `json:"sort" form:"sort"`
}

type AddressStateReq struct {
	Chain   string `json:"-" form:"chain" binding:"required"`
	Address string `json:"-" form:"address" binding:"required"`
}

type SearchReq struct {
	Chain string `json:"chain" form:"chain"`
	Info  string `json:"info" form:"info"`
}

type AddressstatisticReq struct {
	SourceAddress        string `json:"sourceAddress" form:"sourceAddress" binding:"required"`
	TokenContractAddress string `json:"tokenContractAddress" form:"tokenContractAddress"`
	Type                 string `json:"type" form:"type" binding:"required"`
}
