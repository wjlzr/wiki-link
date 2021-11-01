package common

import (
	"github.com/oldfritter/sidekiq-go"
)

type Coin struct {
	Visitable   bool   `json:"-"`
	Index       int    `json:"-"`
	CoinType    string `json:"coinType"`
	Symbol      string `json:"symbol"`
	Name        string `json:"name"`
	CoinName    string `json:"coinName"`
	FullName    string `json:"fullName"`
	WebsiteSlug string `json:"websiteSlug"`
	IconPath    string `json:"iconPath"`
	// Id                  int    `json:"id"`
	// Rank                int    `json:"rank"`
	// Mineable            bool   `json:"mineable"`
	// Algorithm           string `json:"algorithm"`
	// ProofType           string `json:"proofType"`
	// FullyPremined       string `json:"fullyPremined"`
	// PreMinedValue       string `json:"preMinedValue"`
	// NetHashesPerSecond  string `json:"netHashesPerSecond"`
	// BlockReward         string `json:"blockReward"`
	// BlockPeriodTime     string `json:"blockPeriodTime"`
	// CirculatingSupply   int64  `json:"circulatingSupply"`
	// TotalSupply         int64  `json:"totalSupply"`
	// MaxSupply           int64  `json:"maxSupply"`
	// FirstBlockTime      int64  `json:"firstBlockTime"`
	// FirstBlockHeight    int64  `json:"firstBlockHeight"`
	// FirstHistoricalData int64  `json:"firstHistoricalData"`
	// LastHistoricalData  int64  `json:"lastHistoricalData"`
	// LastSyncTime        int64  `json:"lastSyncTime"`
	// CreateTime          int64  `json:"createTime"`
	// UpdateTime          int64  `json:"updateTime"`
}

func (c *Coin) SetCoin(coin *Coin) {
	c.CoinType = coin.CoinType
	c.Symbol = coin.Symbol
	c.Name = coin.Name
	c.CoinName = coin.CoinName
	c.FullName = coin.FullName
	c.WebsiteSlug = coin.WebsiteSlug
	c.IconPath = coin.IconPath
	c.Index = coin.Index
}

var (
	AllWorkers  []sidekiq.Worker
	AllWorkerIs = map[string]func(*sidekiq.Worker) sidekiq.WorkerI{}
	SWI         []sidekiq.WorkerI
	AllCoins    = map[string]*Coin{
		"BTC":        &Coin{Symbol: "BTC", Name: "bitcoin", Visitable: true, Index: 1},
		"ETH":        &Coin{Symbol: "ETH", Name: "ethereum", Visitable: true, Index: 2},
		"DASH":       &Coin{Symbol: "DASH", Name: "dash", Index: 5},
		"LTC":        &Coin{Symbol: "LTC", Name: "litecoin", Index: 3},
		"BCH":        &Coin{Symbol: "BCH", Name: "bitcoin cash", Index: 4},
		"ETC":        &Coin{Symbol: "ETC", Name: "ethereum classic", Index: 6},
		"USDT-OMNI":  &Coin{Symbol: "USDT", Name: "usdt-omni"},
		"USDT-ERC20": &Coin{Symbol: "USDT", Name: "usdt-erc20"},
	}
)

func AllVisitableCoins() []*Coin {
	var coins []*Coin
	for _, coin := range AllCoins {
		if coin.Visitable {
			coins = append(coins, coin)
		}
	}
	return coins
}
