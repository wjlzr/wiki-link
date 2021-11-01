package oklink

import (
	"testing"
	"wiki-link/config"
)

func TestOklink(t *testing.T) {
	config.ReadConfig("D:\\gitlab\\wiki-link\\config")
	// req := TransactionsReq{
	// 	Tag: "okex",
	// }
	// FindByAddress("btc", "3DrATsHzqgsgkbsJB3yEoNhsXniMBQbBBg")
	// SummaryInfo()

	// req := BlocksReq{
	// 	Type: "pool",
	// }
	// s := util.Struct2MapJson(req)
	// fmt.Println(s)
}
