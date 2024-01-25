// mainly deal the the driver
package client

import (
	"crypto/tls"
	"fmt"
	"github.com/go-redis/redis/v8"
	"strings"
)

type Client = redis.UniversalClient

const defaultPoolSize = 3

type ConnInfo struct {
	Host string
	Port int
	Pass string
	User string
	Tls  bool
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
	if len(strings.TrimSpace(info.Pass)) == 0 {
		options = &redis.Options{
			Addr:     addr,
			DB:       db,
			PoolSize: defaultPoolSize,
		}
	} else {
		options = &redis.Options{
			Addr:     addr,
			DB:       db,
			Username: info.User,
			Password: info.Pass,
			PoolSize: defaultPoolSize,
		}
	}
	if info.Tls {
		options.TLSConfig = &tls.Config{InsecureSkipVerify: true}
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
