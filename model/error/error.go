package error

//
type Error struct {
	Code int64  `gorm:"int" json:"code"`
	ZhCN string `gorm:"varchar(300) 'zh-CN'" json:"zh-CN"`
	ZhHK string `gorm:"varchar(300) 'zh-HK'" json:"zh-HK"`
	ZhTW string `gorm:"varchar(300) 'zh-TW'" json:"zh-TW"`
	En   string `gorm:"varchar(300)" json:"en"`
	Vi   string `gorm:"varchar(300)" json:"vi"`
	Th   string `gorm:"varchar(300)" json:"th"`
	Fr   string `gorm:"varchar(1024)" json:"fr"`
	Id   string `gorm:"varchar(1024)" json:"id"`
	Es   string `gorm:"varchar(1024)" json:"es"`
	Ru   string `gorm:"varchar(1024)" json:"ru"`
	De   string `gorm:"varchar(1024)" json:"de"`
	Fil  string `gorm:"varchar(1024)" json:"en-PH"`
	It   string `gorm:"varchar(1024)" json:"it"`
	Hi   string `gorm:"varchar(1024)" json:"hi"`
	Ja   string `gorm:"varchar(1024)" json:"ja"`
}
