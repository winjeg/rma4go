// mainly deal with the args that is passed to the command line
package cmder

import (
	"flag"
	"github.com/fatih/color"
)

var (
	host    string
	port    int
	auth    string
	db      int
	help    bool
	cluster string
)

func init() {
	flag.BoolVar(&help, "h", false, "help content")
	flag.StringVar(&host, "H", "localhost", "address of a redis")
	flag.StringVar(&host, "r", "localhost", "address of a redis")
	flag.IntVar(&port, "p", 6379, "port of the redis")
	flag.StringVar(&auth, "a", "", "password/auth of the redis")
	flag.IntVar(&db, "d", 0, "db of the redis to analyze")
	flag.StringVar(&cluster, "c", "", "cluster info separated by comma, like localhost:123,localhost:456")

	flag.Usage = usage
}

func usage() {
	color.Cyan("rma4go usage:")
	color.Green("rma4go -r some_host -p 6379 -a password -d 0")
	color.Yellow("======================================================")
	flag.PrintDefaults()
}

func GetHost() string {
	return host
}

func GetPort() int {
	return port
}
func GetAuth() string {
	return auth
}

func GetDb() int {
	return db
}

func GetCluster() string {
	return cluster
}