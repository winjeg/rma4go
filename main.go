package main

import (
	"flag"
	"github.com/winjeg/rma4go/analyzer"
	"github.com/winjeg/rma4go/client"
	"github.com/winjeg/rma4go/cmder"
)

func main() {
	flag.Parse()
	if cmder.ShowHelp() {
		flag.Usage()
		return
	}
	printKeyStat()
}


func printKeyStat() {
	h := cmder.GetHost()
	a := cmder.GetAuth()
	p := cmder.GetPort()
	cli := client.BuildRedisClient(client.ConnInfo{
		Host: h,
		Auth: a,
		Port: p,
	}, cmder.GetDb())

	stat := analyzer.ScanAllKeys(cli)
	stat.Print()

}