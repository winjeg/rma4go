package analyzer

import (
	"github.com/winjeg/rma4go/client"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)


var cli = client.BuildRedisClient(client.ConnInfo{
	Host: "localhost",
	Port: 6379,
	Auth: "",
}, 0)

func TestBuildRedisStat(t *testing.T) {
	s := ScanAllKeys(cli)
	d, er := json.Marshal(s)
	if er != nil {
		fmt.Println(er.Error())
	}
	err := ioutil.WriteFile(`a.json`, d, os.ModeAppend)
	if err != nil {
		t.FailNow()
	}
}

func TestGetVersion(t *testing.T) {
	version := getRedisVersion(cli)
	if len(version) < 1 {
		t.FailNow()
	}
}

func TestCompVersion(t *testing.T) {

	r1, err := compareVersion("2.0.0", "2.1")
	if err != nil || r1 != -1 {
		t.FailNow()
	}
	r2, err := compareVersion("2.0.2", "2.0.1")
	if err != nil || r2 !=1  {
		t.FailNow()
	}
	r3, err := compareVersion("2.2", "2.0.1")
	if err != nil || r3 != 1 {
		t.FailNow()
	}
	r4, err := compareVersion("3", "2.0.1")
	if err != nil || r4 != 1 {
		t.FailNow()
	}
	r5, err := compareVersion("2.0.1", "2.0.1")
	if err != nil || r5 != 0 {
		t.FailNow()
	}
}
