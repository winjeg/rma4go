package analyzer

import (
	"context"
	"github.com/go-redis/redis/v8"

	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	redisVersion  = "redis_version"
	serverSection = "Server"
	minVersion    = "4.0.0"
)

type redisInfo map[string]map[string]string

// get parsed redis info
func getRedisMetaInfo(cli redis.UniversalClient) redisInfo {
	infoStr, err := cli.Info(context.Background()).Result()
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	strings.Split(infoStr, "\n")
	d := strings.Split(infoStr, "\n")
	result := make(map[string]map[string]string, defaultSize)
	var lastMap = make(map[string]string, defaultSize)
	var lastSection string
	for _, v := range d {
		if len(strings.TrimSpace(v)) == 0 {
			continue
		}
		if strings.Index(v, "#") == 0 && len(v) > 1 {
			section := strings.TrimSpace(string(v[1:]))
			// 上次的信息存储起来
			if len(lastSection) > 0 {
				result[lastSection] = lastMap
			}
			lastSection = section
			lastMap = make(map[string]string, defaultSize)
		} else {
			kv := strings.Split(v, ":")

			if len(kv) == 2 {
				m := kv[0]
				n := kv[1]
				lastMap[strings.TrimSpace(m)] = strings.TrimSpace(n)
			}
		}
	}
	result[lastSection] = lastMap
	return result
}

func getRedisVersion(cli redis.UniversalClient) string {
	info := getRedisMetaInfo(cli)
	if server, ok := info[serverSection]; ok {
		if version, ok := server[redisVersion]; ok {
			return version
		}
	}
	return ""
}

// check if the redis support 'mem usage' command
func checkSupportMemUsage(cli redis.UniversalClient) bool {
	version := getRedisVersion(cli)
	if len(version) < 1 {
		return false
	}
	i, err := compareVersion(version, minVersion)
	if err != nil {
		return false
	}
	return i > -1
}

// compare version of two strings
// like 2.0.0, 2.0.1, 2.0 etc.
func compareVersion(v1, v2 string) (int, error) {
	p, _ := regexp.Compile(`[\d.]+`)
	illegalVersion := !p.Match([]byte(v1)) || !p.Match([]byte(v2))
	if illegalVersion {
		return 0, errors.New("version illegal")
	}
	v1a := strings.Split(v1, ".")
	v2a := strings.Split(v2, ".")
	v1ia := make([]int, 0, len(v1a))
	v2ia := make([]int, 0, len(v2a))
	for i := range v1a {
		if len(v1a[i]) > 0 {
			v, _ := strconv.Atoi(v1a[i])
			v1ia = append(v1ia, v)
		}
	}
	for i := range v2a {
		if len(v2a[i]) > 0 {
			v, _ := strconv.Atoi(v2a[i])
			v2ia = append(v2ia, v)
		}
	}
	v2Len := len(v2ia)
	for i := range v1ia {
		if i < v2Len {
			if v1ia[i] > v2ia[i] {
				return 1, nil
			} else if v1ia[i] == v2ia[i] {
				continue
			} else {
				return -1, nil
			}
		} else {
			return 1, nil
		}
	}
	if len(v1ia) == len(v2ia) {
		return 0, nil
	}
	return -1, nil
}
