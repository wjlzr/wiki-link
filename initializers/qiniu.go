package initializers

import (
	"github.com/spf13/viper"

	"wiki-link/utils"
)

func initQiniu() {
	cfg := viper.Sub("qiniu")
	if cfg == nil {
		panic("config not found qiniu")
	}
	utils.QiniuConfig["access_key"] = cfg.GetString("access_key")
	utils.QiniuConfig["secret_key"] = cfg.GetString("secret_key")
	utils.QiniuConfig["bucket"] = cfg.GetString("bucket")
}
