// the algorithms to deal with the data
package analyzer

import "fmt"

const (
	typeString = "string"
	typeHash   = "hash"
	typeList   = "list"
	typeSet    = "set"
	typeZSet   = "zset"
)

type KeyMeta struct {
	Key      string
	KeySize  int64
	DataSize int64
	Ttl      int64
	Type     string
}


func BuildRedisStat(meta chan KeyMeta) *RedisStat {
	var a RedisStat
	go func() {
		c := <- meta
		fmt.Println(c)
	}()

	return &a
}


