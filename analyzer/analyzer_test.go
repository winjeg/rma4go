package analyzer

import (
	"encoding/json"
	"fmt"
	"github.com/winjeg/rma4go/client"
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
	err := ioutil.WriteFile(`E:\desktop\a.json`, d, os.ModeAppend)
	if err != nil {
		t.FailNow()
	}
}
