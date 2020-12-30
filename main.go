package main

import (
	"fmt"
	"github.com/winjeg/redis"
	"github.com/winjeg/rma4go/analyzer"
	"github.com/winjeg/rma4go/client"
	"github.com/winjeg/rma4go/cmder"
	"strconv"

	"flag"
	// _ "net/http/pprof"
	"strings"
)

type Client = redis.UniversalClient

func main() {
	//go func() {
	//	http.ListenAndServe("0.0.0.0:8899", nil)
	//}()

	flag.Parse()
	if cmder.ShowHelp() {
		flag.Usage()
		return
	}
	printKeyStat()
}

func printKeyStat() {
	var cli Client
	cluster := cmder.GetCluster()
	if len(cluster) > 0 {
		urls := strings.Split(cluster, ",")
		for _, v := range urls {
			hp := strings.Split(v, ":")
			port, _ := strconv.Atoi(hp[1])
			cli = client.BuildRedisClient(client.ConnInfo{
				Host: hp[0],
				Auth: cmder.GetAuth(),
				Port: port,
			}, cmder.GetDb())
			stat := analyzer.ScanAllKeys(cli)
			fmt.Printf("analysis result for:%s is as follows\n", v)
			stat.Print()
		}

	} else {
		cli = client.BuildRedisClient(client.ConnInfo{
			Host: cmder.GetHost(),
			Auth: cmder.GetAuth(),
			Port: cmder.GetPort(),
		}, cmder.GetDb())
		stat := analyzer.ScanAllKeys(cli)
		stat.Print()
	}
}
