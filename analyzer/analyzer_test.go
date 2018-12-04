package analyzer

import "testing"

func TestBuildRedisStat(t *testing.T) {
	a := make(chan KeyMeta)
	meta := KeyMeta{
		Key:"dassdas",
	}
	BuildRedisStat(a)
	a <- meta
}