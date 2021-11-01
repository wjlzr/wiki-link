package address

import (
	"fmt"
	"strings"
	"wiki-link/common"

	"github.com/StrongSquirrel/crptwav"
	"github.com/cpacia/bchutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/ripemd160"
)

const addressChecksumLen = 4

//检查地址属于哪条公链
//address 			地址
func Validation(chain, address string) bool {
	chain = strings.ToLower(chain)
	address = strings.Trim(address, " ")

	var isRight bool = false

	switch chain {
	case common.ETH:
		if ethAddressValidate(address) == strings.ToLower(address) {
			isRight = true
		}
	case common.BTC:
		return btcAddressValidate(address)
	case common.DASH:
		return dashAddressValidate(address)
	case common.BCH:
		return bchAddressValidate(address)
	case common.LTC:
		return ltcAddressValidate(address)
	}
	return isRight
}

//检查地址属于哪条公链
//address 			地址
func ValidationWithoutChain(address string) string {
	address = strings.Trim(address, " ")
	for _, chain := range common.AllCoins {
		switch strings.ToLower(chain.Symbol) {
		case common.ETH:
			if ethAddressValidate(address) == strings.ToLower(address) {
				return common.ETH
			}
		case common.BTC:
			if btcAddressValidate(address) {
				return common.BTC
			}
		case common.DASH:
			if dashAddressValidate(address) {
				return common.DASH
			}
		case common.LTC:
			if ltcAddressValidate(address) {
				return common.LTC
			}
		case common.BCH:
			if bchAddressValidate(address) {
				return common.BCH
			}

		}
	}
	return ""
}

//eth地址检测
func ethAddressValidate(address string) string {
	if len([]byte(address)) != 42 {
		return ""
	}
	addrLowerStr := strings.ToLower(address)
	if strings.HasPrefix(addrLowerStr, "0x") {
		addrLowerStr = addrLowerStr[2:]
		address = address[2:]
	}
	var binaryStr string
	addrBytes := []byte(addrLowerStr)
	hash256 := crypto.Keccak256Hash([]byte(addrLowerStr))

	for i, e := range addrLowerStr {
		//如果是数字则跳过
		if e >= '0' && e <= '9' {
			continue
		} else {
			binaryStr = fmt.Sprintf("%08b", hash256[i/2]) //注意，这里一定要填充0
			if binaryStr[4*(i%2)] == '1' {
				addrBytes[i] -= 32
			}
		}
	}
	return "0x" + strings.ToLower(string(addrBytes))
}

//btc地址检测
func btcAddressValidate(address string) bool {
	// addr, err := btcsuitebtcutil.DecodeAddress(address, nil)
	// if err != nil || addr == nil {
	// 	return false
	// }
	// return true

	len := len(address)
	if len < 25 {
		return false
	}

	if strings.HasPrefix(address, "1") && len >= 26 && len <= 34 {
		return true
	}
	if (strings.HasPrefix(address, "3") && len == 34) || (strings.HasPrefix(address, "bc1") && len > 34) {
		return true
	}

	return false
}

//dash地址检测
func dashAddressValidate(address string) bool {
	return crptwav.IsValidAddress(address, "dash")
}

//bch地址检测
func bchAddressValidate(address string) bool {
	if !btcAddressValidate(address) {
		decoded, _, typ, err := bchutil.CheckDecodeCashAddress("bitcoincash:" + address)
		if err != nil {
			return false
		}

		switch len(decoded) {
		case ripemd160.Size:
			switch typ {
			case bchutil.P2PKH, bchutil.P2SH:
				return true
			default:
				return false
			}
		}
	}
	return true
}

//ltc地址检测
func ltcAddressValidate(address string) bool {
	if !crptwav.IsValidAddress(address, "LTC") {
		if strings.HasPrefix(address, "ltc1") && len(address) == 43 {
			return true
		}
		return false
	}
	return true
}
