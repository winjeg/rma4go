package main

import (
	"github.com/winjeg/rma4go/analyzer"
	"github.com/winjeg/rma4go/client"
	"github.com/winjeg/rma4go/cmder"

	"flag"
)

func main() {
	flag.Parse()
	if cmder.ShowHelp() {
		flag.Usage()
		return
	}
	printKeyStat(cmder.GetHost(), cmder.GetAuth(), cmder.GetPort())
}

func printKeyStat(host, auth string, port int) {
	var cli = client.BuildRedisClient(client.ConnInfo{
		Host: host,
		Auth: auth,
		Port: port,
	}, cmder.GetDb())
	stat := analyzer.ScanAllKeys(cli)
	stat.Print()
}
