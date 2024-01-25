// mainly deal with the args that is passed to the command line
package cmder

import (
	"flag"
	"github.com/fatih/color"
)

var (
	host  = "127.0.0.1"
	port  = 6379
	pass  = ""
	db    = 0
	help  = false
	match string
	tls   = false
	user  = ""
)

func init() {
	flag.BoolVar(&help, "h", false, "help content")
	flag.StringVar(&host, "H", "localhost", "address of a redis")
	flag.StringVar(&host, "r", "localhost", "address of a redis")
	flag.IntVar(&port, "p", 6379, "port of the redis")
	flag.StringVar(&pass, "a", "", "password/auth of the redis")
	flag.StringVar(&match, "m", "*", "match the pattern to scan, like 'a*'")
	flag.IntVar(&db, "d", 0, "db of the redis to analyze")
	flag.BoolVar(&tls, "tls", false, "enable tls")
	flag.StringVar(&user, "u", "", "user/auth of the redis")
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

func GetPass() string {
	return pass
}

func GetUser() string {
	return user
}

func GetTLS() bool {
	return tls
}

func GetDb() int {
	return db
}

func GetMatch() string {
	return match
}

func ShowHelp() bool {
	return help
}
