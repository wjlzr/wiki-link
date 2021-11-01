package initializers

import (
	"github.com/spf13/viper"

	"wiki-link/baseServices/email"
)

func initSmtp() {
	smtp := viper.Sub("smtp")
	if smtp == nil {
		panic("config not found redis")
	}
	email.InitSmtp(smtp)
}
