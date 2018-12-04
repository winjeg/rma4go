// data structures
package analyzer

const (
	// expire section
	Second = 1000
	Minute = Second * 60
	Hour   = Minute * 60
	Day    = Hour * 24
	Week   = Day * 7


	// data size
	B	  = 1
	KB	  = B * 1024
	MB	  = KB * 1024
)

type RedisStat struct {
	All     KeyStat `json:"all"`
	String  KeyStat `json:"string"`
	Hash    KeyStat `json:"hash"`
	Set     KeyStat `json:"set"`
	List    KeyStat `json:"list"`
	ZSet    KeyStat `json:"zset"`
	BigKeys BigKeys `json:"bigKeys"`
}

// distributions of keys of all prefixes
type Distribution struct {
	KeyType    string `json:"type"`
	KeyPattern string `json:"pattern"`
	Metrics
}

// basic metrics of a group of key
type Metrics struct {
	KeyCount       int64 `json:"keyCount"`
	KeySize        int64 `json:"keySize"`
	DataSize       int64 `json:"dataSize"`
	KeyNeverExpire int64 `json:"neverExpire"`
	ExpireInHour   int64 `json:"expireInHour"`  // >= 0h < 1h
	ExpireInDay    int64 `json:"expireInDay"`   // >= 1h < 24h
	ExpireInWeek   int64 `json:"expireInWeek"`  // >= 1d < 7d
	ExpireOutWeek  int64 `json:"expireOutWeek"` // >= 7d
}

// total stat and distributions
type KeyStat struct {
	Distribution []Distribution `json:"distribution"`
	Metrics
}

// for big keys
type BigKeys struct {
	GT1MCount   int64    `json:"gt1mCount"` // >= 1m < 2m
	GT1MPattern []string `json:"gt1mPattern"`
	GT2MCount   int      `json:"gt2mCount"` // >= 2m < 5m
	GT2MPattern []string `json:"gt2mPattern"`
	GT5MCount   []int    `json:"gt5mCount"` // >= 5m
	GT5MPattern []string `json:"gt5mPattern"`
}
