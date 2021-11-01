package model

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Address struct {
	gorm.Model
	DataSource string          `gorm:"size:16"`
	Project    string          `gorm:"size:32"`
	Type       string          `gorm:"size:16"`
	Tag        string          `gorm:"size:64"`                              //  OKEx|...
	Coin       string          `gorm:"index:idx0_on_address,unique;size:16"` //  btc|eth
	Address    string          `gorm:"index:idx0_on_address,unique;size:64" json:"address"`
	Balance    decimal.Decimal `gorm:"type:decimal(36,12)" json:"balance"`

	TagList []struct {
		Project string `json:"project"`
		Tag     string `json:"tag"`
	} `gorm:"-" json:"tagList"`
}

func (address *Address) BeforeSave(db *gorm.DB) (err error) {
	address.Balance = address.Balance.Truncate(12)
	return
}

func (address *Address) BeforeUpdate(db *gorm.DB) (err error) {
	address.Balance = address.Balance.Truncate(12)
	return
}
