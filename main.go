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
	printKeyStat(cmder.GetHost(), cmder.GetUser(), cmder.GetPass(), cmder.GetPort(), cmder.GetTLS())
}

func printKeyStat(host, user, pass string, port int, tls bool) {
	var cli = client.BuildRedisClient(client.ConnInfo{
		Host: host,
		User: user,
		Pass: pass,
		Port: port,
		Tls:  tls,
	}, cmder.GetDb())
	stat := analyzer.ScanAllKeys(cli)
	stat.Print()
}
