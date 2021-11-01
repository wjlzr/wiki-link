package oklink

import (
	"encoding/json"
	"strings"
	"wiki-link/core/response"

	"wiki-link/common"
	"wiki-link/core/http"
	"wiki-link/core/util"
)

//根据输入的内容对 token 进行模糊匹配
//chain  待查询区块链的符号小写
//q      模糊搜索的代币符号,代币名称，合约地 址
func MatchTokens(chain, q string) (*Base, error) {
	chain = strings.ToLower(chain)
	var (
		result Base
		err    error
		query  = map[string]interface{}{
			"q": q,
		}
	)
	if chain != "" {
		if err := http.OKLinkInvoke("search", &result, query, chain); err != nil {
			return &result, err
		}
	} else {
		var re []interface{}
		for _, coin := range common.AllCoins {
			if err := http.OKLinkInvoke("search", &result, query, strings.ToLower(coin.Symbol)); err == nil {
				re = append(re, result.Data)
			}
		}
		result.Data = re
	}
	return &result, err
}

//根据输入的高度/交易 Hash/地址等数据返回查询的类型
//chain     待查询区块链的符号小写
//q			待查询的高度/交易 Hash/地址
func Search(chain, q string) (*SearchResp, error) {
	chain = strings.ToLower(chain)
	var (
		result SearchResp
		query  = map[string]interface{}{
			"q": q,
		}
	)
	err := http.OKLinkInvoke("chainSearch", &result, query, chain)
	return &result, err

}

//查询一个区块链的详情信息
//chain     待查询区块链的符号小写
func ChainInfo(chain string) (*ChainInfoResp, error) {
	var result ChainInfoResp
	chain = strings.ToLower(chain)
	err := http.OKLinkInvoke("chainInfo", &result, nil, chain)
	return &result, err
}

//查询支持的区块链的汇总信息
func SummaryInfo() (*SummaryInfoResp, error) {
	var result SummaryInfoResp
	err := http.OKLinkInvoke("summary", &result, nil)
	return &result, err
}

//查询一个区块链的基本统计信息。
func StatisticCommon(chain string, req *ChainPagination) (result Base, err error) {
	chain = strings.ToLower(chain)
	var query map[string]interface{}
	if req != nil {
		req.InitPagination()
		query = util.Paged2MapJson(*req)
	}
	err = http.OKLinkInvoke("statisticCommon", &result, query, chain)
	return
}

//查询一个区块链的 24 小时大额转账信息
func StatisticLargeTransfer(chain string) (result Base, err error) {
	chain = strings.ToLower(chain)
	err = http.OKLinkInvoke("statisticLargeTransferLastest", &result, nil, chain)
	return
}

//查询一个区块链的主链币或代币的持仓分布统计。
func StatisticRichersStat(chain, tokenId string) (result Base, err error) {
	chain = strings.ToLower(chain)
	var query map[string]interface{}
	if tokenId != "" {
		query = make(map[string]interface{})
		query["tokenId"] = tokenId
	}
	err = http.OKLinkInvoke("statisticRichersStat", &result, query, chain)
	return
}

//查询一个区块链的区块列表,区块的排序默认为区块高度倒序
//sort 排序，规则为 field:desc|asc 支持的排序字 段为 height、timestamp，默认 height 倒 序
//offset  起始位置
//limit  返回条数
func Blocks(chain string, req *BlocksReq) (result Base, err error) {
	chain = strings.ToLower(chain)
	var query map[string]interface{}
	if req != nil {
		req.InitPagination()
		query = util.Paged2MapJson(*req)
	}
	err = http.OKLinkInvoke("blocks", &result, query, chain)
	return
}

//查询一个区块链的区块详情
//chain 待查询区块链的符号小写
//param 区块高度或区块 hash，为 lastest 获取最新区块
func GetBlockInfo(chain string, req *BlockInfoReq, result http.Result) (err error) {
	chain = strings.ToLower(chain)
	err = http.OKLinkInvoke("blocksInfo", result, nil, chain, req.Info)
	return
}

//查询一个区块链的交易列表
func Transactions(chain string, req *TransactionsReq) (result TransactionResp, err error) {
	var (
		ethResult  EthTransactionResp
		btcResullt BtcTransactionResp
		query      map[string]interface{}
	)
	chain = strings.ToLower(chain)
	if req != nil {
		req.InitPagination()
		query = util.Paged2MapJson(*req)
	}
	switch chain {
	case "btc":
		err = http.OKLinkInvoke("transactions", &btcResullt, query, chain)
		result = btcResullt
	case "eth":
		err = http.OKLinkInvoke("transactions", &ethResult, query, chain)
		result = ethResult
	}

	return result, err
}

func TransactionsNewEdition(chain string, req *TransactionsReq, result http.Result) (err error) {
	var query map[string]interface{}
	chain = strings.ToLower(chain)
	if req != nil {
		req.InitPagination()
		query = util.Paged2MapJson(*req)
	}
	err = http.OKLinkInvoke("transactions", result, query, chain)
	return
}

//查询一个区块链的交易详情
//chain  待查询区块链的符号小写
//hash   交易 hash
func GetTransactionsByHash(chain, hash string, result http.Result) (err error) {
	chain = strings.ToLower(chain)
	err = http.OKLinkInvoke("transactionsByHash", result, nil, chain, hash)
	return
}

//查询一个区块链的未确认交易列表，
func TransactionsTrade(chain string, req *TransactionsReq) (result Base, err error) {
	chain = strings.ToLower(chain)
	var query map[string]interface{}
	if req != nil {
		req.InitPagination()
		query = util.Paged2MapJson(*req)
	}
	err = http.OKLinkInvoke("transactionsTrade", &result, query, chain)
	return
}

//查询一个地址的基本信息
//chain    待查询区块链的符号小写
//address
func FindByAddress(chain, address string, result http.Result) (err error) {
	chain = strings.ToLower(chain)
	err = http.OKLinkInvoke("findByAddress", result, nil, chain, address)
	return
}

//根据地址查询地址的余额
//chain    待查询区块链的符号小写
//address
func FindBalanceByAddress(chain, address string) (*BalanceByAddressResp, error) {
	var result BalanceByAddressResp
	chain = strings.ToLower(chain)
	err := http.OKLinkInvoke("findBalanceByAddress", &result, nil, chain, address)
	return &result, err
}

//  查询一个地址的交易列表
//  chain    待查询区块链的符号小写
//  address
func FindTransactionsByAddress(chain string, req *TransactionsReq) (result Base, err error) {
	chain = strings.ToLower(chain)
	var query map[string]interface{}
	if req != nil {
		req.InitPagination()
		query = util.Paged2MapJson(*req)
	}
	err = http.OKLinkInvoke("findTransactionsByAddress", &result, query, chain, req.Address)
	return
}

//  查询一个区块链的富豪地址排行信息
//  chain    待查询区块链的符号小写
func FindRichers(chain string, req *ChainPagination) (result Base, err error) {
	chain = strings.ToLower(chain)
	var query map[string]interface{}
	if req != nil {
		req.InitPagination()
		query = util.Paged2MapJson(*req)
	}
	err = http.OKLinkInvoke("findRichers", &result, query, chain)
	return
}

//  查询一个地址的分析信息
//  chain    待查询区块链的符号小写
func AnalysisByAddress(chain string, req *AnalysisReq) (result Base, err error) {
	chain = strings.ToLower(chain)
	var query map[string]interface{}
	if req != nil {
		query = util.Struct2MapJson(*req)
	}
	err = http.OKLinkPost("analysis", &result, query, chain)
	return
}

// 普通转账
func Transfers(chain string, req *TransfersReq) (result Base, err error) {
	chain = strings.ToLower(chain)
	var query map[string]interface{}
	if req != nil {
		req.InitPagination()
		query = util.Paged2MapJson(*req)
	}
	err = http.OKLinkInvoke("transfers", &result, query, chain)
	return
}

// token 持仓
func TokensHolders(chain string, req *PageReq) (result Base, err error) {
	chain = strings.ToLower(chain)
	var query map[string]interface{}
	if req != nil {
		req.InitPagination()
		query = util.Paged2MapJson(*req)
	}
	err = http.OKLinkInvoke("tokensHolders", &result, query, chain, req.Address)
	return
}

// token列表
func Tokens(chain string, req *PageReq) (result Base, err error) {
	chain = strings.ToLower(chain)
	var query map[string]interface{}
	if req != nil {
		req.InitPagination()
		query = util.Paged2MapJson(*req)
	}
	err = http.OKLinkInvoke("tokens", &result, query, chain, req.Address)
	return
}

// 地址持仓
func AddressHolders(chain string, req *PageReq, result http.Result) (err error) {
	chain = strings.ToLower(chain)
	var query map[string]interface{}
	if req != nil {
		req.InitPagination()
		query = util.Paged2MapJson(*req)
	}
	err = http.OKLinkInvoke("addressHolders", result, query, chain, req.Address)
	return
}

// 地址普通交易
func TransactionsByClassfy(chain string, req *PageReq) (result Base, err error) {
	chain = strings.ToLower(chain)
	var query map[string]interface{}
	if req != nil {
		req.InitPagination()
		query = util.Paged2MapJson(*req)
	}
	err = http.OKLinkInvoke("transactionsByClassfy", &result, query, chain, req.Address)
	return
}

// 块的普通交易
func TransactionsNoRestrict(chain string, req *PageReq) (r *response.TransactionNoRestrictResp, err error) {
	chain = strings.ToLower(chain)
	var query map[string]interface{}
	if req != nil {
		req.InitPagination()
		query = util.Paged2MapJson(*req)
	}
	var result Base
	err = http.OKLinkInvoke("transactionsNoRestrict", &result, query, chain)
	res, _ := json.Marshal(result)
	_ = json.Unmarshal(res, &r)
	return
}

// 合约地址
func ContractByAddress(chain, address string) (result Base, err error) {
	chain = strings.ToLower(chain)
	err = http.OKLinkInvoke("contract", &result, nil, chain, address)
	return
}

// 合约地址的ERC20（ERC721）交易
func TransfersByAddress(chain string, req *AddressTransfers) (result Base, err error) {
	chain = strings.ToLower(chain)
	var query map[string]interface{}
	if req != nil {
		req.InitPagination()
		query = util.Paged2MapJson(*req)
	}
	err = http.OKLinkInvoke("addressTransfers", &result, query, chain, req.Address)
	return
}

// 块的内部交易
func TransactionsInternal(chain string, req *InternalTransactions) (result Base, err error) {
	chain = strings.ToLower(chain)
	var query map[string]interface{}
	if req != nil {
		req.InitPagination()
		query = util.Paged2MapJson(*req)
	}
	err = http.OKLinkInvoke("internalTransactions", &result, query, chain)
	return
}

// ETH叔块
func UncleBlocks(chain string, req *UncleBlocksReq) (result Base, err error) {
	chain = strings.ToLower(chain)
	var query map[string]interface{}
	if req != nil {
		req.InitPagination()
		query = util.Paged2MapJson(*req)
	}
	err = http.OKLinkInvoke("uncleBlocks", &result, query, chain)
	return
}

// 叔块详情
func UncleBlock(chain string, req *UncleBlockReq) (result Base, err error) {
	chain = strings.ToLower(chain)
	err = http.OKLinkInvoke("uncleBlock", &result, nil, chain, req.Info)
	return
}

func TransactionsLogs(chain string, req *TransactionsLogsReq) (result Base, err error) {
	chain = strings.ToLower(chain)
	err = http.OKLinkInvoke("transactionsLogs", &result, nil, chain, req.Info)
	return
}

// 匹配token
func TokensMatch(chain string, req *TokensMatchReq) (result Base, err error) {
	chain = strings.ToLower(chain)
	err = http.OKLinkInvoke("tokensMatch", &result, nil, chain, req.Info)
	return
}

// 交易图谱
func AnalysisTransactions(chain string, req *AnalysisTransactionsReq) (result Base, err error) {
	chain = strings.ToLower(chain)
	var query map[string]interface{}
	if req != nil {
		query = util.Struct2MapJson(*req)
	}
	err = http.OKLinkPost("analysisTransactions", &result, query, chain)
	return
}

// 地址的内部交易
func InternalTx(chain string, req *InternalTxReq) (result Base, err error) {
	chain = strings.ToLower(chain)
	var query map[string]interface{}
	if req != nil {
		req.InitPagination()
		query = util.Paged2MapJson(*req)
	}
	err = http.OKLinkInvoke("internalTx", &result, query, chain, req.Address)
	return
}

// 地址间的交易信息
func AnalysisAddressesEdgeTxs(chain string, req *AnalysisAddressesEdgeTxsReq) (result Base, err error) {
	chain = strings.ToLower(chain)
	var query map[string]interface{}
	if req != nil {
		req.InitPagination()
		query = util.Paged2MapJson(*req)
	}
	err = http.OKLinkInvoke("analysisAddressesEdgeTxs", &result, query, chain)
	return
}

// 查询地址的各种最大值
func AddressState(chain string, req *AddressStateReq) (result Base, err error) {
	chain = strings.ToLower(chain)
	err = http.OKLinkInvoke("addressState", &result, nil, chain, req.Address)
	return
}

//
func Addressstatistic(req *AddressstatisticReq) (result Base, err error) {
	var query map[string]interface{}
	if req != nil {
		query = util.Struct2MapJson(*req)
	}
	err = http.OKLinkPost("addressStatistic", &result, query)
	return
}
