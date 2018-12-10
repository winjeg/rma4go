// data structures
package analyzer

const (
	defaultSize = 128
	compactNum  = 30
	minKeyLen   = 5
	// expire section
	Second = 1000000000
	Minute = Second * 60
	Hour   = Minute * 60
	Day    = Hour * 24
	Week   = Day * 7

	// data size
	B  = 1
	KB = B * 1024
	MB = KB * 1024
)
const (
	typeString = "string"
	typeHash   = "hash"
	typeList   = "list"
	typeSet    = "set"
	typeZSet   = "zset"
	typeOther  = "other"
)

type KeyMeta struct {
	Key      string
	KeySize  int64
	DataSize int64
	Ttl      int64
	Type     string
}

type RedisStat struct {
	All     KeyStat `json:"all"`
	String  KeyStat `json:"string"`
	Hash    KeyStat `json:"hash"`
	Set     KeyStat `json:"set"`
	List    KeyStat `json:"list"`
	ZSet    KeyStat `json:"zset"`
	Other   KeyStat `json:"other"`
	BigKeys KeyStat `json:"bigKeys"`
}

// distributions of keys of all prefixes
type Distribution struct {
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

func (m *Metrics) MergeMeta(meta KeyMeta) {
	m.DataSize += meta.DataSize
	m.KeySize += meta.KeySize
	m.KeyCount ++
	switch {
	case meta.Ttl < 0:
		m.KeyNeverExpire++
	case meta.Ttl >= 0 && meta.Ttl < Hour:
		m.ExpireInHour++
	case meta.Ttl >= Hour && meta.Ttl < Day:
		m.ExpireInDay++
	case meta.Ttl >= Day && meta.Ttl < Week:
		m.ExpireInWeek++
	case meta.Ttl >= Week:
		m.ExpireOutWeek++
	}
}

// total stat and distributions
type KeyStat struct {
	Distribution map[string]Distribution `json:"distribution"`
	Metrics
}

func (stat *RedisStat) Compact() {
	stat.All.compact()
	stat.String.compact()
	stat.BigKeys.compact()
	stat.Other.compact()
	stat.Hash.compact()
	stat.ZSet.compact()
	stat.Set.compact()
	stat.List.compact()
}

func (stat *RedisStat) Merge(meta KeyMeta) {
	stat.All.Merge(meta)
	// big keys
	if meta.DataSize >= 1*MB {
		stat.BigKeys.Merge(meta)
	}
	switch meta.Type {
	case typeString:
		stat.String.Merge(meta)
	case typeList:
		stat.List.Merge(meta)
	case typeHash:
		stat.Hash.Merge(meta)
	case typeSet:
		stat.Set.Merge(meta)
	case typeZSet:
		stat.ZSet.Merge(meta)
	default:
		stat.Other.Merge(meta)
	}
}

func (stat *KeyStat) Merge(meta KeyMeta) {
	stat.MergeMeta(meta)
	dists := stat.Distribution
	if dists == nil {
		dists = make(map[string]Distribution, defaultSize)
	}
	keyLen := len(meta.Key)

	// check for if there are already some key in the collection
	inMap := false
	for i := 0; i < keyLen; i++ {
		x := meta.Key[0 : i+1]
		if v, ok := dists[x]; ok {
			d := Distribution(v)
			d.MergeMeta(meta)
			dists[x] = d
			inMap = true
		}
	}
	//
	if !inMap {
		var d Distribution
		d.MergeMeta(meta)
		dists[meta.Key] = d
	}

	stat.Distribution = dists
}

func (stat *KeyStat) compact() {
	distMap := stat.Distribution
	tmpMap := make(map[string][]string, defaultSize)
	shrinkTo := compactNum
	for k := range distMap {
		compactedKey := k
		if orgks, ok := tmpMap[compactedKey]; ok {
			orgks = append(orgks, k)
			tmpMap[compactedKey] = orgks
		} else {
			ks := make([]string, 0, defaultSize)
			ks = append(ks, k)
			tmpMap[compactedKey] = ks
		}
	}
	shrinkTo--
	for len(tmpMap) > compactNum && shrinkTo >= minKeyLen {
		tnMap := make(map[string][]string, defaultSize)
		for k := range tmpMap {
			// shrink
			if len(k) > shrinkTo {
				compactedKey := k[0:shrinkTo]
				if oik, ok := tnMap[compactedKey]; ok {
					oik = append(oik, tmpMap[k]...)
					tnMap[compactedKey] = oik

				} else {
					ks := make([]string, 0, defaultSize)
					ks = append(ks, tmpMap[k]...)
					tnMap[compactedKey] = ks
				}
			} else {
				tnMap[k] = tmpMap[k]
			}
		}

		// 如果此次shrink 没有使得这个集合的元素数量增加， 就使用原来的key
		for k := range tmpMap {
			if len(k) > shrinkTo {
				ck := k[0:shrinkTo]
				if len(tnMap[ck]) == len(tmpMap[k]) && len(tnMap[ck]) > 1 {
					x := make([]string, 0, defaultSize)
					tnMap[k] = append(x, tnMap[ck]...)
					delete(tnMap, ck)
				}
			}
		}
		tmpMap = tnMap
		shrinkTo --
	}

	dists := make(map[string]Distribution, defaultSize)
	for k, v := range tmpMap {
		if len(v) > 1 {
			var nd Distribution
			for _, dk := range v {
				d := distMap[dk]
				nd.KeyCount += d.KeyCount
				nd.KeySize += d.KeySize
				nd.DataSize += d.DataSize
				nd.ExpireInHour += d.ExpireInHour
				nd.ExpireInWeek += d.ExpireInWeek
				nd.ExpireInDay += d.ExpireInDay
				nd.ExpireOutWeek += d.ExpireOutWeek
				nd.KeyNeverExpire += d.KeyNeverExpire
			}
			dists[k] = nd
		} else {
			for _, dk := range v {
				dists[dk] = distMap[dk]
			}
		}
	}
	stat.Distribution = dists
}
