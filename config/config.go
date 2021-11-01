package config

import (
	"github.com/oldfritter/sidekiq-go"
	"github.com/spf13/viper"
)

var (
	conf           *Config = new(Config)
	ConfigFilePath         = "config/config.yml"
)

type Config struct {
	OkLink oklink
	MySQL  mysql
	// RedisCluster rediscluster
	// Redis        redis
	Workers []sidekiq.Worker
}

type oklink struct {
	ApiKey string
}

//mysql配置
type mysql struct {
	DriverName   string
	Dsn          string
	MaxOpenConns int
	MaxIdleConns int
}

// type rediscluster struct {
//   Addrs       []string
//   Password    string
//   DialTimeout int
//   PoolSize    int
// }
//
// type redis struct {
//   Addr        string
//   Password    string
//   DialTimeout int
//   PoolSize    int
// }

//读取配置文件
func ReadConfig(path string) error {
	viper.SetConfigFile(path)

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	conf.OkLink.ApiKey = viper.GetString("oklink.api_key")

	return nil
}

//
func Conf() *Config {
	return conf
}

//获取api私钥
func (c *Config) GetOkLinkApiKey() string {
	return c.OkLink.ApiKey
}

// //获取redis集群配置
// func (c *Config) GetRedisClusterConfig() (cs []interface{}, is []int) {
//   cs = []interface{}{
//     c.RedisCluster.Addrs,
//     c.RedisCluster.Password,
//   }
//   is = []int{
//     c.RedisCluster.DialTimeout,
//     c.RedisCluster.PoolSize,
//   }
//   return
// }
