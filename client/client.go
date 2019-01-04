// mainly deal the the driver
package client

import (
	"github.com/winjeg/redis"

	"fmt"
)

type Client = redis.UniversalClient

const defaultPoolSize = 3

type ConnInfo struct {
	Host string
	Port int
	Auth string
}

func BuildRedisClient(info ConnInfo, db int) Client {

	// check db
	if db < 0 || db > 15 {
		db = 0
	}

	// check info
	if len(info.Host) == 0 || info.Port < 0 || info.Port > 65535 {
		return nil
	}

	cli := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", info.Host, info.Port),
		Password: info.Auth, // no password set
		DB:       db,
		PoolSize: defaultPoolSize,
	})
	return cli
}

func BuildClusterClient(urls []string, auth string) Client {
	return redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    urls,
		Password: auth,
	})
}
