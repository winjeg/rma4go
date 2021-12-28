// mainly deal the the driver
package client

import (
	"github.com/go-redis/redis/v8"
	"strings"

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
	addr := fmt.Sprintf("%s:%d", info.Host, info.Port)
	var options *redis.Options
	if len(strings.TrimSpace(info.Auth)) == 0 {
		options = &redis.Options{
			Addr:     addr,
			DB:       db,
			PoolSize: defaultPoolSize,
		}
	} else {
		options = &redis.Options{
			Addr:     addr,
			DB:       db,
			Password: info.Auth,
			PoolSize: defaultPoolSize,
		}
	}
	cli := redis.NewClient(options)
	return cli
}

func BuildClusterClient(urls []string, auth string) Client {
	return redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    urls,
		Password: auth,
	})
}
