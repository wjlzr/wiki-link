package db

import (
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

var (
	client        *redis.Client
	clusterClient *redis.ClusterClient
)

func InitRedis(cfg *viper.Viper) {
	var r struct {
		Addr        string
		Password    string
		Dialtimeout int64
		Poolsize    int
	}
	r.Addr = cfg.GetString("addr")
	r.Password = cfg.GetString("password")
	r.Dialtimeout = cfg.GetInt64("dialtimeout")
	r.Poolsize = cfg.GetInt("poolsize")

	client = redis.NewClient(&redis.Options{
		Addr:        r.Addr,
		Password:    r.Password,
		DialTimeout: time.Second * time.Duration(r.Dialtimeout),
		PoolSize:    r.Poolsize,
	})
	_, err := client.Ping().Result()
	if err != nil {
		panic(fmt.Sprintf("redis connect error :%v", err))
	} else {
		log.Printf("redis connect success!")
	}
}

func RedisClient() *redis.Client {
	return client
}

func CloseRedisClient() error {
	return client.Close()
}

//初始化集群
func NewRedisClusterClient(cs []interface{}, is []int) (err error) {
	if len(cs) != 2 && len(is) != 2 {
		return
	}

	clusterClient = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:       cs[0].([]string),
		Password:    cs[1].(string),
		DialTimeout: time.Second * time.Duration(is[0]),
		PoolSize:    is[1],
	})

	_, err = clusterClient.Ping().Result()
	return
}

//
func RedisClusterClient() *redis.ClusterClient {
	return clusterClient
}

func InitRedisCluster(cfg *viper.Viper) {
	var rc struct {
		Addrs       []string
		Password    string
		Dialtimeout int64
		Poolsize    int
	}
	rc.Addrs = cfg.GetStringSlice("addrs")
	rc.Password = cfg.GetString("password")
	rc.Dialtimeout = cfg.GetInt64("dialtimeout")
	rc.Poolsize = cfg.GetInt("poolsize")

	clusterClient = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:       rc.Addrs,
		Password:    rc.Password,
		DialTimeout: time.Second * time.Duration(rc.Dialtimeout),
		PoolSize:    rc.Poolsize,
	})
	_, err := clusterClient.Ping().Result()
	if err != nil {
		panic("redis connect error")
	} else {
		log.Printf("redis connect success!")
	}
}
