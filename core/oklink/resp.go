package oklink

import (
	"encoding/json"
	"log"
	"strings"

	"wiki-link/common"
	"wiki-link/core/http"
)

//基本信息
type Base struct {
	Code      int         `json:"code"`
	Msg       string      `json:"msg"`
	DetailMsg string      `json:"detailMsg"`
	Data      interface{} `json:"data"`
}

func (base *Base) GetCode() int {
	return base.Code
}

func (base *Base) GetMsg() string {
	return base.Msg
}

func (base *Base) GetData() interface{} {
	return base.Data
}

type hits struct {
	Symbol               string `json:"symbol"`
	CoinName             string `json:"coinName"`
	TokenContractAddress string `json:"tokenContractAddress"`
	LogoUrl              string `json:"logoUrl"`
	Marketcap            int    `json:"marketcap"`
}

//
type SearchResp struct {
	Base
	Data struct {
		DataType string `json:"dataType"`
		Chain    string `json:"chain"`
	} `json:"data"`
}

//链基本信息
type ChainInfoResp struct {
	Base
	Data struct {
		Id                    int              `json:"id"`          //唯一标识
		CoinType              string           `json:"coinType"`    //币种类型
		Symbol                string           `json:"symbol"`      //符号
		Name                  string           `json:"name"`        //币种名称
		CoinName              string           `json:"coinName"`    //币种名称
		FullName              string           `json:"fullName"`    //币种全名
		WebsiteSlug           string           `json:"websiteSlug"` //货币互联网名称
		IconPath              string           `json:"iconPath"`    //图标链接
		Rank                  int              `json:"rank"`        //排名
		Mineable              bool             `json:"mineable"`    //是否挖矿
		Algorithm             string           `json:"algorithm"`   //挖矿算法
		ProofType             string           `json:"proofType"`   //
		FullyPremined         bool             `json:"fullyPremined"`
		PreMinedValue         string           `json:"preMinedValue"`
		NetHashesPerSecond    int              `json:"netHashesPerSecond"`
		BlockReward           int              `json:"blockReward"`
		BlockPeriodTime       int64            `json:"blockPeriodTime"`
		CirculatingSupply     float64          `json:"circulatingSupply"`
		TotalSupply           float64          `json:"totalSupply"`
		MaxSupply             float64          `json:"maxSupply"`
		TokenAddress          string           `json:"tokenAddress"`
		FirstBlockTime        int64            `json:"firstBlockTime"`
		FirstBlockHeight      int64            `json:"firstBlockHeight"`
		FirstHistoricalData   int64            `json:"firstHistoricalData"`
		LastHistoricalData    int64            `json:"lastHistoricalData"`
		LastSyncTime          int64            `json:"lastSyncTime"`
		CreateTime            int64            `json:"createTime"`
		UpdateTime            int64            `json:"updateTime"`
		IcoPrice              string           `json:"icoPrice"`
		Market                market           `json:"market"`
		Address               address          `json:"address"`
		Block                 block            `json:"block"`
		Fee                   fee              `json:"fee"`
		GlobalDifficulty      globalDifficulty `json:"globalDifficulty"` //全网难度
		Hashes                hashes           `json:"hashes"`           //算力
		Mine                  mine             `json:"mine"`             //挖矿
		ReduceReward          reduceReward     `json:"reduceReward"`     //减半奖励
		Transaction           transaction      `json:"transaction"`      //交易信息
		Usdt                  usdt             `json:"USDT"`
		Okchain               string           `json:"okchain"`
		TrxUsdtTotalSupply    int64            `json:"trxUsdtTotalSupply"`
		EthUsdtTotalSupply    int64            `json:"ethUsdtTotalSupply"`
		CoreAlgorithm         string           `json:"coreAlgorithm"`
		MasterNodeCount       int              `json:"masterNodeCount"`
		DiffChangeRemainBlock int64            `json:"diffChangeRemainBlock"`
		PlatformId            int64            `json:"platformId"`
		Platform              int              `json:"platform"`
		Eth                   eth              `json:"eth"`
	} `json:"data"`
}

func (cir *ChainInfoResp) SetIcon() {
	url := "https://q.wikibit.com/zorro/zorro.json"
	bs, err := http.GetFromOKLink(url)
	if err != nil {
		log.Println(err)
		return
	}
	var rs []struct {
		BaseTokenName string `json:"BaseTokenName"`
		Name          string `json:"Name"`
		Icon          string `json:"Icon"`
	}
	if err := json.Unmarshal(bs, &rs); err != nil {
		log.Println(err)
		return
	}
	for _, r := range rs {
		if r.BaseTokenName == cir.Data.Symbol {
			cir.Data.IconPath = r.Icon
			for _, coin := range common.AllCoins {
				if strings.ToLower(coin.Symbol) == strings.ToLower(r.BaseTokenName) {
					coin.IconPath = r.Icon
					coin.Name = cir.Data.Name
					coin.Symbol = cir.Data.Symbol
					coin.CoinType = cir.Data.CoinType
					coin.CoinName = cir.Data.CoinName
					coin.FullName = cir.Data.FullName
					coin.WebsiteSlug = cir.Data.WebsiteSlug
				}
			}
		}
	}
}

//
type usdt struct {
	Fee                float64 `json:"fee"`
	OmiUsdtTotalSupply float64 `json:"omiUsdtTotalSupply"`
	TrxUsdtTotalSupply float64 `json:"trxUsdtTotalSupply"`
	EthUsdtTotalSupply float64 `json:"ethUsdtTotalSupply"`
	WeekAddCoin        float64 `json:"weekAddCoin"`
	WeekDestoryCoin    float64 `json:"weekDestoryCoin"`
}

//
type transaction struct {
	PendingTransactionCount       int     `json:"pendingTransactionCount"`       //未确认交易数
	PendingTransactionSize        float64 `json:"pendingTransactionSize"`        //pending 交易大小
	TransactionValue24h           float64 `json:"transactionValue24h"`           //24 小时的链上交易量
	TransactionCount24h           int     `json:"transactionCount24h"`           //24 小时平均交易数量
	TotalTransactionCount         int64   `json:"totalTransactionCount"`         //链上交易宗数
	TranRate                      float64 `json:"tranRate"`                      //50 块平均 tps
	AvgTransactionCount24h        float64 `json:"avgTransactionCount24h"`        //24 小时平均交易数量
	AvgTransactionCount24hPercent float64 `json:"avgTransactionCount24hPercent"` //24 小时平均交易数量涨幅
}

//
type reduceReward struct {
	NextReduceRewardTime   int64 `json:"nextReduceRewardTime"`   //预测下次产量减半时间
	NextReduceRewardHeight int64 `json:"nextReduceRewardHeight"` //预测下次产量减半高度
}

//
type mine struct {
	AvgMineReward24h         float64 `json:"avgMineReward24h"`         //24 小时平均挖矿奖励
	MinerIncomePerUnit       float64 `json:"minerIncomePerUnit"`       //每单位算力收益
	MinerIncomePerUnitAddFee float64 `json:"minerIncomePerUnitAddFee"` //每单位算力收益
	MinerIncomePerUnitCoin   float64 `json:"minerIncomePerUnitCoin"`   //每单位算力收益币数
}

//
type hashes struct {
	GlobalHashes                 string  `json:"globalHashes"`                 //全网算力
	GlobalHashesPercentChange24h float64 `json:"globalHashesPercentChange24h"` //全网算力 24 小时变化比分币
}

//
type globalDifficulty struct {
	CurrentDiffculty              string  `json:"currentDiffculty"`              //当前难度
	RealCurrentDiffculty          float64 `json:"realCurrentDiffculty"`          //
	CurrentDiffcultyPercentChange float64 `json:"currentDiffcultyPercentChange"` //上次难度调整百分比
	CurrentDiffcultyChangeTime    int64   `json:"currentDiffcultyChangeTime"`    //上次难度调整时间
	CurrentDiffcultyChangeHeight  int64   `json:"currentDiffcultyChangeHeight"`  //上次难度调整高度
	NextDiffculty                 string  `json:"nextDiffculty"`                 //预测下次难度
	RealNextDiffculty             float64 `json:"realNextDiffculty"`             //
	NextDiffcultyChangeTime       int64   `json:"nextDiffcultyChangeTime"`       //预测下次难度调整时间
	NextDiffcultyChangeHeight     int64   `json:"nextDiffcultyChangeHeight"`     //预测下次难度调整高度
	NextDiffcultyPercentChange    float64 `json:"nextDiffcultyPercentChange"`    //预测下次难度调整百分比
	NextDifficultyChangeBlock     int64   `json:"nextDifficultyChangeBlock"`     //预测下次难度调整区块
}

//
type fee struct {
	BestFeePerKbyte  int64       `json:"bestFeePerKbyte"`
	BestFeePerKwu    interface{} `json:"bestFeePerKwu"`
	BestFeePerKvbyte float64     `json:"bestFeePerKvbyte"`
	BestGasPrice     int64       `json:"bestGasPrice"`
}

//
type block struct {
	Height                      int64   `json:"height"`                      //高度
	FirstExchangeHistoricalTime int64   `json:"firstExchangeHistoricalTime"` //第一次交易时间
	FirstBlockTime              int64   `json:"firstBlockTime"`              //第一个出块时间
	FirstBlockHeight            int64   `json:"firstBlockHeight"`            //第一个区块高度
	AvgBlockInterval            int64   `json:"avgBlockInterval"`            //平均出块时间
	AvgBlockSize24h             float64 `json:"avgBlockSize24h"`             //24 小时平均区块大小
	AvgBlockSize24hPercent      float64 `json:"avgBlockSize24hPercent"`      //24 小时平均区块大小涨幅
	MediaBlockSize              float64 `json:"mediaBlockSize"`              //一周区块中位数大小
	HalveTime                   int64   `json:"halveTime"`                   //减半时间
}

type address struct {
	ValidAddressCount     int64 `json:"validAddressCount"`
	ValidAddressCountDiff int   `json:"validAddressCountDiff"`
	NewAddressCount24h    int   `json:"newAddressCount24h"`
}

//
type market struct {
	Symbol           string  `json:"symbol"`           //符号
	MarketSymbol     string  `json:"marketSymbol"`     //法币符号
	Price            float64 `json:"price"`            //价格
	Volume24h        float64 `json:"volume24h"`        // 24 小时交易量
	PercentChange1h  float64 `json:"percentChange1h"`  //价格 1 小时涨幅
	PercentChange24h float64 `json:"percentChange24h"` //价格 24 小时涨幅
	PercentChange7d  float64 `json:"percentChange7d"`  //价格 7 天涨幅
	MarketCap        float64 `json:"marketCap"`        //市值
	Timestamp        int64   `json:"timestamp"`        //时间戳
}

//
type eth struct {
	ValueCoinAmount           int64   `json:"valueCoinAmount"`
	ContractAmount            int64   `json:"contractAmount"`
	Erc20Amount               int64   `json:"erc20Amount"`
	InternalTransactionAmount int64   `json:"internalTransactionAmount"`
	TokenAmount               int64   `json:"tokenAmount"`
	UncleBlockAmount          int64   `json:"uncleBlockAmount"`
	NewErc20Amount            int64   `json:"newErc20Amount"`
	GasFee                    float64 `json:"gasFee"`
	Erc20Value                float64 `json:"erc20Value"`
}

//
type SummaryInfoResp struct {
	Base
	Data []Summary `json:"data"`
}

type Summary struct {
	common.Coin
	// Symbol                  string  `json:"symbol"`                  //符号
	Price                   float64 `json:"price"`                   //价格
	LastBlockTime           int64   `json:"lastBlockTime"`           //上一个区块时间
	TotalTransacationCount  int64   `json:"totalTransacationCount"`  // 交易总数
	TotalSupply             float64 `json:"totalSupply"`             //流通量
	TotalSupplyPercent      float64 `json:"totalSupplyPercent"`      //流通量占比
	PendingTransactionCount int64   `json:"pendingTransactionCount"` //未确认交易总数
	RunTime                 int64   `json:"runTime"`                 //运行时间
	PricePercentChange24h   float64 `json:"pricePercentChange24h"`   //24 小时价格变化
	Height                  int64   `json:"height"`                  //高度
}

//
type StatisticCommonResp struct {
	Base
	Data []statisticCommonData `json:"data"`
}

type statisticCommonData struct {
	Day                     int64  `json:"day"`                     //时间
	ActiveAddressCount      int    `json:"activeAddressCount"`      //活跃地址树
	ValidAddressCount       int    `json:"validAddressCount"`       //持币地址数
	NewAddressCount         int    `json:"newAddressCount"`         //新增地址数
	TransactionCount        int    `json:"transactionCount"`        //交易数
	TransactionValue        string `json:"transactionValue"`        //交易币数量
	TransactionValuePercent string `json:"transactionValuePercent"` //交易币数量 24 小时占比
	ActiveAddressPercent    string `json:"activeAddressPercent"`    //活跃地址数 24 小时占比
	ValidAddressPercent     string `json:"validAddressPercent"`     //持币地址数 24 小时占比
	NewAddressPercent       string `json:"newAddressPercent"`       //新增地址数 24 小时占比
	TransactionCountPercent string `json:"transactionCountPercent"` //交易数 24 小时占比
	PricePercent            string `json:"pricePercent"`            //价格 24 小时占比
	Market                  market `json:"market"`                  //行情信息
	Price                   string `json:"price"`                   //价格
}

//
type StatisticLargeTransferResp struct {
	Base
	Data []StatisticLargeTransfeData `json:"data"`
}

type StatisticLargeTransfeData struct {
	Datetime             int64   `json:"datetime"`             //时间
	TransferRangeFrom    int     `json:"transferRangeFrom"`    //转账数量起点
	TransferRangeTo      int     `json:"transferRangeTo"`      //转账数量终点
	TransferCount        int     `json:"transferCount"`        //转帐笔数
	TransferValue        float64 `json:"transferValue"`        //转账总币数
	TransferAddressCount int     `json:"transferAddressCount"` //转账的地址数
	Market               market  `json:"market"`               //行情信息
}

//
type StatisticRichersStatResp struct {
	Base
	Data []StatisticRichersStatData `json:"data"`
}

type StatisticRichersStatData struct {
	HolderRangeFrom int     `json:"holderRangeFrom"` //排名起点
	HolderRangeTo   int     `json:"holderRangeTo"`   //排名终点
	AddressCount    int     `json:"addressCount"`    //地址总数
	HoldersValue    float64 `json:"holdersValue"`    //持币总数
	PercentRate     float64 `json:"percentRate"`     //占比
}

//
type BlocksResp struct {
	Base
	Data struct {
		Total  int64      `json:"total"`
		Hits   []blockHit `json:"hits"`
		Extend string     `json:"extend"`
	} `json:"data"`
}

//
type blockHit struct {
	Hash                  string      `json:"hash"`   //块哈希
	Height                int         `json:"height"` //高度
	TransactionCount      int         `json:"transactionCount"`
	PreviousBlockHash     string      `json:"previousBlockHash"`
	NextBlockHash         string      `json:"nextBlockHash"`
	Blocktime             int64       `json:"blocktime"`
	LegalRate             interface{} `json:"legalRate"`
	TotalTransactionCount int         `json:"totalTransactionCount"`
	Size                  int         `json:"size"`
	Version               int64       `json:"version"`
	MerkleRoot            string      `json:"merkleRoot"`
	InputsCount           int         `json:"inputsCount"`
	OutputsCount          int         `json:"outputsCount"`
	InputsValue           float64     `json:"inputsValue"`
	OutputsValue          float64     `json:"outputsValue"`
	Nonce                 string      `json:"nonce"`
	MinerHash             string      `json:"minerHash"`
	GuessedMiner          string      `json:"guessedMiner"`
	Reward                float64     `json:"reward"`
	Fee                   float64     `json:"fee"`
	FeePerKbyte           int         `json:"feePerKbyte"`
	BlockReward           float64     `json:"blockReward"`
	Aux                   bool        `json:"aux"`
	MedianTime            int64       `json:"medianTime"`
	Bits                  string      `json:"bits"`
	Difficulty            string      `json:"difficulty"`
	MineDifficulty        string      `json:"mineDifficulty"`
	Chainwork             string      `json:"chainwork"`
	CoinbaseDataHex       string      `json:"coinbaseDataHex"`
	InputsValueSat        int64       `json:"inputsValueSat"`
	OutputsValueSat       int64       `json:"outputsValueSat"`
	CoindaysDestroyed     int64       `json:"coindaysDestroyed"`
	FeePerKwu             int         `json:"feePerKwu"`
	WitnessCount          int         `json:"witnessCount"`
	Weight                int         `json:"weight"`
	StrippedSize          int         `json:"strippedSize"`
	MinepoolName          string      `json:"minepoolName"` //矿池名称
	MinepoolCode          string      `json:"minepoolCode"` //矿池 url
	MinepoolUrl           string      `json:"minepoolUrl"`
	MinepoolLogoUrl       string      `json:"minepoolLogoUrl"` // 矿池图标的地址
	Confirm               int         `json:"confirm"`
}

//
type BlockInfoResp struct {
	Base
	Data blockHit `json:"data"`
}

type TransactionResp interface {
	Size() int
	GetData() interface{}
}

//btc交易
type BtcTransactionResp struct {
	Base
	Data struct {
		Total  int64            `json:"total"`
		Hits   []BtcTransaction `json:"hits"`
		Extend string           `json:"extend"`
	} `json:"data"`
}

func (e BtcTransactionResp) Size() int {
	return len(e.Data.Hits)
}

func (e BtcTransactionResp) GetData() interface{} {
	return e.Data
}

//btc交易
type BtcTransaction struct {
	Hash                 string   `json:"hash"`
	Blocktime            int64    `json:"blocktime"`
	LegalRate            float64  `json:"legalRate"`
	BlockHash            string   `json:"blockHash"`
	BlockHeight          int      `json:"blockHeight"`
	TotalIndex           int64    `json:"totalIndex"`
	Hour                 int      `json:"hour"`
	Coinbase             bool     `json:"coinbase"`
	Size                 int      `json:"size"`
	Version              int      `json:"version"`
	DoubleSpend          bool     `json:"doubleSpend"`
	Fee                  int      `json:"fee"`
	FeePerKbyte          int      `json:"feePerKbyte"`
	InputsCount          int      `json:"inputsCount"`
	OutputsCount         int      `json:"outputsCount"`
	InputsValue          float64  `json:"inputsValue"`
	OutputsValue         float64  `json:"outputsValue"`
	RealTransferValue    float64  `json:"realTransferValue"`
	InputsValueSat       int      `json:"inputsValueSat"`
	OutputsValueSat      int      `json:"outputsValueSat"`
	RealTransferValueSat int      `json:"realTransferValueSat"`
	Inputs               []input  `json:"inputs"`
	Outputs              []output `json:"outputs"`
	LockTime             int64    `json:"lockTime"`
	CoindaysDestroyed    int      `json:"coindaysDestroyed"`
	Sigops               int      `json:"sigops"`
	CoinString           string   `json:"coinString"`
	StrippedSize         int      `json:"strippedSize"`
	VirtualSize          int      `json:"virtualSize"`
	Weight               int      `json:"weight"`
	HasWitness           bool     `json:"hasWitness"`
	WitnessHash          string   `json:"witnessHash"`
	FeePerKwu            int      `json:"feePerKwu"`
	FeePerKvbyte         int      `json:"feePerKvbyte"`
	Confirm              int      `json:"confirm"`
	RealAddressBalance   int      `json:"realAddressBalance"`
}

//
type input struct {
	PrevBlockHash     string             `json:"prevBlockHash"`
	PrevAddresses     []string           `json:"prevAddresses"`
	ScriptType        string             `json:"scriptType"`
	PreAddressesTags  []preAddressesTags `json:"preAddressesTags"`
	VinIndex          int                `json:"vinIndex"`
	PrevBlockHeight   int                `json:"prevBlockHeight"`
	PrevTxhash        string             `json:"prevTxhash"`
	PrevVoutIndex     int                `json:"prevVoutIndex"`
	PrevValueSat      int                `json:"prevValueSat"`
	PrevValue         float64            `json:"prevValue"`
	PrevBlocktime     int64              `json:"prevBlocktime"`
	PrevScriptType    string             `json:"prevScriptType"`
	ScriptHex         string             `json:"scriptHex"`
	ScriptData        string             `json:"scriptData"`
	Sequence          int64              `json:"sequence"`
	Lifespan          int                `json:"lifespan"`
	CoindaysDestroyed int                `json:"coindaysDestroyed"`
	Witness           []string           `json:"witness"`
}

//
type output struct {
	Addresses        []string           `json:"addresses"`
	ScriptType       string             `json:"scriptType"`
	AddressesTags    []preAddressesTags `json:"addressesTags"`
	VoutIndex        int                `json:"voutIndex"`
	ValueSat         int                `json:"valueSat"`
	Value            float64            `json:"value"`
	FromCoinbase     bool               `json:"fromCoinbase"`
	Spent            bool               `json:"spent"`
	SpentBlockHeight int                `json:"spentBlockHeight"`
	SpentBlockHash   string             `json:"spentBlockHash"`
	SpentTxhash      string             `json:"spentTxhash"`
	SpentVinIndex    int                `json:"spentVinIndex"`
	SpentBlocktime   int                `json:"spentBlocktime"`
	ScriptAsm        string             `json:"scriptAsm"`
	ScriptHex        string             `json:"scriptHex"`
	OutputType       string             `json:"outputType"`
}

type preAddressesTags struct {
	Address  string      `json:"address"`
	TagLogos interface{} `json:"tagLogos"`
}

//eth交易
type EthTransactionResp struct {
	Base
	Data struct {
		Total  int64            `json:"total"`
		Hits   []EthTransaction `json:"hits"`
		Extend string           `json:"extend"`
	} `json:"data"`
}

func (e EthTransactionResp) Size() int {
	return len(e.Data.Hits)
}
func (e EthTransactionResp) GetData() interface{} {
	return e.Data
}

//
type EthTransaction struct {
	Hash                     string      `json:"hash"`
	Blocktime                int64       `json:"blocktime"`
	Index                    int         `json:"index"`
	BlockHash                string      `json:"blockHash"`
	BlockHeight              int64       `json:"blockHeight"`
	TotalIndex               int64       `json:"totalIndex"`
	Hour                     int         `json:"hour"`
	From                     string      `json:"from"`
	To                       string      `json:"to"`
	Fee                      float64     `json:"fee"`
	Value                    float64     `json:"value"`
	IsContractCall           bool        `json:"isContractCall"`
	Status                   string      `json:"status"`
	IsFromContract           bool        `json:"isFromContract"`
	IsToContract             bool        `json:"isToContract"`
	GasLimit                 int         `json:"gasLimit"`
	GasUsed                  int         `json:"gasUsed"`
	GasPrice                 int64       `json:"gasPrice"`
	CumulativeGasUsed        int64       `json:"cumulativeGasUsed"`
	Nonce                    int         `json:"nonce"`
	InputHex                 string      `json:"inputHex"`
	V                        int         `json:"v"`
	R                        string      `json:"r"`
	S                        string      `json:"s"`
	Erc20TokenTransferCount  int         `json:"erc20TokenTransferCount"`
	Erc721TokenTransferCount int         `json:"erc721TokenTransferCount"`
	InternalTranCount        int         `json:"internalTranCount"`
	ValueTotal               float64     `json:"valueTotal"`
	InternalValueTotal       float64     `json:"internalValueTotal"`
	Confirm                  int         `json:"confirm"`
	RealValue                float64     `json:"realValue"`
	FromTag                  interface{} `json:"fromTag"`
	ToTag                    interface{} `json:"toTag"`
	FromTokenUrl             string      `json:"fromTokenUrl"`
	ToTokenUrl               string      `json:"toTokenUrl"`
	LegalRate                float64     `json:"legalRate"`
}

type EthTransactionByHashResp struct {
	Base
	Data EthTransaction `json:"data"`
}

type BtcTransactionByHashResp struct {
	Base
	Data BtcTransaction `json:"data"`
}

type BtcFindByAddressResp struct {
	Base
	Data btcFindByAddress `json:"data"`
}

type Richer interface {
	GetData() interface{}
}

type BtcRicherAddressResp struct {
	Base
	Data btcRichers `json:"data"`
}

func (b BtcRicherAddressResp) GetData() interface{} {
	return b.Data
}

type btcRichers struct {
	Totol int                `json:"total"`
	Hits  []btcFindByAddress `json:"hits"`
}

type btcFindByAddress struct {
	Address                string    `json:"address"`
	Balance                float64   `json:"balance"`
	LegalRate              float64   `json:"legalRate"`
	BalanceSat             int64     `json:"balanceSat"`
	TotalRecievedSat       int64     `json:"totalRecievedSat"`
	TotalRecieved          float64   `json:"totalRecieved"`
	TotalSentSat           int64     `json:"totalSentSat"`
	TotalSent              float64   `json:"totalSent"`
	TxCount                int       `json:"txCount"`
	UnconfirmedTxCount     int       `json:"unconfirmedTxCount"`
	UnconfirmedReceivedSat int       `json:"unconfirmedReceivedSat"`
	UnconfirmedSentSat     int       `json:"unconfirmedSentSat"`
	UnspentTxCount         int       `json:"unspentTxCount"`
	FirstTransactionTime   int64     `json:"firstTransactionTime"`
	LastTransactionTime    int64     `json:"lastTransactionTime"`
	TagList                []tagList `json:"tagList"`
	UsdtBalance            float64   `json:"usdtBalance"`
}

type tagList struct {
	Tag  string `json:"tag"`
	Logo string `json:"logo"`
	Item string `json:"item"`
}

type EthFindByAddressResp struct {
	Base
	Data ethFindByAddress `json:"data"`
}

type EthRicherAddressResp struct {
	Base
	Data ethRichers `json:"data"`
}

func (b EthRicherAddressResp) GetData() interface{} {
	return b.Data
}

type ethRichers struct {
	Totol int                `json:"total"`
	Hits  []ethFindByAddress `json:"hits"`
}
type ethFindByAddress struct {
	Address                string        `json:"address"`
	Balance                float64       `json:"balance"`
	LegalRate              float64       `json:"legalRate"`
	TotalRecieved          float64       `json:"totalRecieved"`
	TotalSent              float64       `json:"totalSent"`
	TxCount                int           `json:"txCount"`
	UnconfirmedTxCount     int           `json:"unconfirmedTxCount"`
	UnconfirmedReceived    float64       `json:"unconfirmedReceived"`
	UnconfirmedSent        float64       `json:"unconfirmedSent"`
	FirstTransactionTime   int64         `json:"firstTransactionTime"`
	LastTransactionTime    int64         `json:"lastTransactionTime"`
	Type                   string        `json:"type"`
	BalanceWei             int           `json:"balanceWei"`
	TotalRecievedWei       int           `json:"totalRecievedWei"`
	TotalSentWei           int           `json:"totalSentWei"`
	HasInternalTransaction bool          `json:"hasInternalTransaction"`
	HasErc20Transaction    bool          `json:"hasErc20Transaction"`
	HasErc721Transaction   bool          `json:"hasErc721Transaction"`
	Symbol                 string        `json:"symbol"`
	CoinName               string        `json:"coinName"`
	LogoUrl                string        `json:"logoUrl"`
	TokenContractAddress   string        `json:"tokenContractAddress"`
	TokenCount             int           `json:"tokenCount"`
	TabList                []string      `json:"tabList"`
	TagList                []interface{} `json:"tagList"` // [{"tag":"OKEx","logo":"","project":"Exchange","item":"","type":"User"}]
}

type AnalysisResp interface {
	GetData() interface{}
}

func CreateAnalysisResp(chain string) AnalysisResp {
	switch chain {
	case "btc":
		return BtcAnalysisResp{}
	case "eth":
		return EthAnalysisResp{}
	default:
		return nil
	}
}

type EthAnalysisResp struct {
	Base
	Data struct {
		Edges     []EthAnalysisItem `json:"edges"`
		EdgeNum   int               `json:"edgeNum"`
		VertexNum int               `json:"vertexNum"`
		MaxValue  float64           `json:"maxValue"`
	} `json:"data"`
}

func (e EthAnalysisResp) GetData() interface{} {
	return e.Data
}

type EthAnalysisItem struct {
	From         string  `json:"from"`
	To           string  `json:"to"`
	TotalValue   float64 `json:"totalValue"`
	Direction    int     `json:"direction"`
	Txnum        int     `json:"txnum"`
	RequestId    string  `json:"requestId"`
	UpdateTime   int64   `json:"updateTime"`
	RequestIndex int     `json:"requestIndex"`
}

type BtcAnalysisResp struct {
	Base
	Data struct {
		Edges     []BtcAnalysisItem `json:"edges"`
		EdgeNum   int               `json:"edgeNum"`
		VertexNum int               `json:"vertexNum"`
		MaxValue  float64           `json:"maxValue"`
	} `json:"data"`
}

func (e BtcAnalysisResp) GetData() interface{} {
	return e.Data
}

type BtcAnalysisItem struct {
	EthAnalysisItem
	EdgeId string `json:"edgeId"`
}

type SimpleBlockInfoResp struct {
	Base
	Data SimpleBlockInfo `json:"data"`
}

type SimpleBlockInfo struct {
	common.Coin
	Blocktime             int64   `json:"blocktime"`
	Confirm               int64   `json:"confirm"`
	Difficulty            string  `json:"difficulty"`
	Fee                   float64 `json:"fee"`
	Height                int64   `json:"height"`
	MinepoolCode          string  `json:"minepoolCode"`
	MinepoolLogoUrl       string  `json:"minepoolLogoUrl"`
	MinepoolName          string  `json:"minepoolName"`
	MinepoolUrl           string  `json:"minepoolUrl"`
	MinerHash             string  `json:"minerHash"`
	Nonce                 string  `json:"nonce"`
	Size                  int64   `json:"size"`
	Hash                  string  `json:"hash"`
	PreviousBlockHash     string  `json:"previousBlockHash"`
	TotalTransactionCount int64   `json:"totalTransactionCount"`
	TransactionCount      int64   `json:"transactionCount"`
}

type SimpleTransactionsByHashResp struct {
	Base
	Data SimpleTransactionsByHash `json:"data"`
}

type SimpleTransactionsByHash struct {
	common.Coin
	BlockHash         string  `json:"blockHash"`
	BlockHeight       int64   `json:"blockHeight"`
	Blocktime         int64   `json:"blocktime"`
	Confirm           int64   `json:"confirm"`
	Fee               float64 `json:"fee"`
	Hash              string  `json:"hash"`
	Index             int64   `json:"index"`
	RealTransferValue float64 `json:"realTransferValue"`
	Value             float64 `json:"value"`
}

type SimpleAddressResp struct {
	Base
	Data SimpleAddress `json:"data"`
}

type SimpleAddress struct {
	common.Coin
	Address            string  `json:"address"`
	Balance            float64 `json:"balance"`
	LegalRate          float64 `json:"legalRate"`
	TotalRecieved      float64 `json:"totalRecieved"`
	TotalSent          float64 `json:"totalSent"`
	TxCount            int     `json:"txCount"`
	UnconfirmedTxCount int     `json:"unconfirmedTxCount"`
}
