package initializers

import (
	"wiki-link/config"
	"wiki-link/db"
)

func InitAllResources() {
	config.ReadConfig(config.ConfigFilePath)
	initMainDb()
	initRedis()
	initWorkers()
	initSmtp()
	initQiniu()
	initS3()
}

func CloseResources() {
	db.CloseDB()
	db.CloseRedisClient()
}
