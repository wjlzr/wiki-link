package response

// 查询一个高度的区块的交易列表
type TransactionNoRestrictResp struct {
	Code      int    `json:"code"`
	Msg       string `json:"msg"`
	DetailMsg string `json:"detailMsg"`
	Data      struct {
		Total  int64                           `json:"total"`
		Extend interface{}                     `json:"extend"`
		Hits   []TransactionNoRestrictDataHits `json:"hits"`
	} `json:"data"`
}

//
type TransactionNoRestrictDataHits struct {
	Hash                 string      `json:"hash"`
	BlockTime            int64       `json:"blocktime"`
	LegalRate            float64     `json:"legalRate"`
	Index                int64       `json:"index"`
	BlockHash            string      `json:"blockHash"`
	BlockHeight          int64       `json:"blockHeight"`
	Coinbase             bool        `json:"coinbase"`
	Size                 int64       `json:"size"`
	Version              int64       `json:"version"`
	DoubleSpend          bool        `json:"doubleSpend"`
	Fee                  float64     `json:"fee"`
	FeePerKbyte          int64       `json:"feePerKbyte"`
	InputsCount          int64       `json:"inputsCount"`
	OutputsCount         int64       `json:"outputsCount"`
	InputsValue          float64     `json:"inputsValue"`
	OutputsValue         float64     `json:"outputsValue"`
	RealTransferValue    float64     `json:"realTransferValue"`
	InputsValueSat       int64       `json:"inputsValueSat"`
	OutputsValueSat      int64       `json:"outputsValueSat"`
	RealTransferValueSat int64       `json:"realTransferValueSat"`
	Inputs               interface{} `json:"inputs"`
	PrevAddresses        []string    `json:"prevAddresses"`
	ScriptType           string      `json:"scriptType"`
	VinIndex             int64       `json:"vinIndex"`
	PrevBlockHeight      int64       `json:"prevBlockHeight"`
	PrevTxhash           string      `json:"prevTxhash"`
	PrevVoutIndex        int64       `json:"prevVoutIndex"`
	PrevValueSat         int64       `json:"prevValueSat"`
	PrevValue            float64     `json:"prevValue"`
	PrevBlocktime        int64       `json:"prevBlocktime"`
	PrevScriptType       string      `json:"prevScriptType"`
	ScriptHex            string      `json:"scriptHex"`
	ScriptData           string      `json:"scriptData"`
	Sequence             int64       `json:"sequence"`
	Lifespan             int64       `json:"lifespan"`
	CoindaysDestroyed    int64       `json:"coindaysDestroyed"`
	Witness              []string    `json:"witness"`
	Outputs              interface{} `json:"outputs"`
	Addresses            []string    `json:"addresses"`
	VoutIndex            int64       `json:"voutIndex"`
	ValueSat             int64       `json:"valueSat"`
	Value                float64     `json:"value"`
	FromCoinbase         bool        `json:"v"`
	Spent                bool        `json:"spent"`
	SpentBlockHeight     int64       `json:"spentBlockHeight"`
	SpentBlockHash       string      `json:"spentBlockHash"`
	SpentTxHash          string      `json:"spentTxhash"`
	SpentVinIndex        string      `json:"spentVinIndex"`
	SpentBlockTime       int64       `json:"spentBlocktime"`
	ScriptAsm            string      `json:"scriptAsm"`
	OuputType            string      `json:"ouputType"`
	LockTime             int64       `json:"lockTime"`
	Sigops               int64       `json:"sigops"`
	StrippedSize         int64       `json:"strippedSize"`
	VirtualSize          int64       `json:"virtualSize"`
	Weight               int64       `json:"weight"`
	HasWitness           bool        `json:"hasWitness"`
	WitnessHash          string      `json:"witnessHash"`
	FeePerKwu            int64       `json:"feePerKwu"`
	FeePerKvbyte         int64       `json:"feePerKvbyte"`
	Confirm              int64       `json:"confirm"`
	RealAddressBalance   float64     `json:"realAddressBalance"`
	TexTend              interface{} `json:"textend"`
}
