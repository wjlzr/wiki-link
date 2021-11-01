package model

type BlockHeight struct {
	ID     uint   `gorm:"primaryKey"`
	Coin   string `gorm:"index:idx0_on_block_height,unique;size:3;"` //  btc|eth
	Number int    `gorm:"index:idx0_on_block_height,unique"`         //  blockHeight
}
