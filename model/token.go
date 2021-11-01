package model

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Token struct {
	gorm.Model
	AddressId uint            `gorm:"index:idx0_on_address_and_coin,unique" json:"address_id"`
	Coin      string          `gorm:"index:idx0_on_address_and_coin,unique;size:16"` //  usdt-omni|usdt-erc20
	Balance   decimal.Decimal `gorm:"type:decimal(32,8)" json:"balance"`
}
