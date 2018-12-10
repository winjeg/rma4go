package analyzer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestBuildRedisStat(t *testing.T) {
	s := ScanAllKeys(Client)
	d, er := json.Marshal(s)
	if er != nil {
		fmt.Println(er.Error())
	}
	ioutil.WriteFile(`E:\desktop\a.json`, d, os.ModeAppend)
}
