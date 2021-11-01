package initializers

import (
	"github.com/spf13/viper"

	"wiki-link/db"
)

//初始化redis
func initRedis() {
	switch viper.Sub("base").GetInt("redis") {
	case 1:
		initSingleRedis()
	case 2:
		initRedisCluster()
	}
}

//初始化Redis单节点
func initSingleRedis() {
	cfgRedis := viper.Sub("redis")
	if cfgRedis == nil {
		panic("config not found redis")
	}
	db.InitRedis(cfgRedis)
}

//初始化redis集群
func initRedisCluster() {
	cfgRedisCluster := viper.Sub("rediscluster")
	if cfgRedisCluster == nil {
		panic("config not found redisCluster")
	}
	db.InitRedisCluster(cfgRedisCluster)
}
