package initializers

import (
	"github.com/spf13/viper"

	"wiki-link/utils"
)

func initS3() {
	cfg := viper.Sub("s3")
	if cfg == nil {
		panic("config not found s3")
	}
	utils.S3Config["AWS_REGION"] = cfg.GetString("AWS_REGION")
	utils.S3Config["AWS_ACCESS_KEY_ID"] = cfg.GetString("AWS_ACCESS_KEY_ID")
	utils.S3Config["AWS_SECRET_ACCESS_KEY"] = cfg.GetString("AWS_SECRET_ACCESS_KEY")
	utils.S3Config["S3_BACKUP_BUCKET"] = cfg.GetString("S3_BACKUP_BUCKET")

}
