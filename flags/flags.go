/*
 * @Author: facsert
 * @Date: 2023-08-01 21:46:50
 * @LastEditTime: 2023-08-16 22:53:33
 * @LastEditors: facsert
 * @Description: package flags
 */

package flags

import (
	"flag"
	"fmt"
	"os"
)


type Param struct {
	Name    string
	Default string
	Help    string
}

var paramList = []Param{
	{
		Name:    "host",
		Default: "127.0.0.1",
		Help:    "server host",
	},
	{
		Name:    "username",
		Default: "root",
		Help:    "server username",
	},
	{
		Name:    "password",
		Default: "admin",
		Help:    "server password",
	},
}

var USAGE = `
Usage of %s:
    %s
	-host string
		set host (default "127.0.0.1")
	-password string
		set password (default "admin")
	-username string
		set username (default "root")
`

func FlagInit() Param {
    param := Param{}
	flag.StringVar(&param.Name, "name", "default_name", "set name")
	flag.StringVar(&param.Default, "default", "default_default", "set default")
	flag.StringVar(&param.Help, "help", "set help", "set help")

	flag.Usage = func() {
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, USAGE, os.Args[0])
	}

	flag.Parse()
	
	return param
}

type Args struct {
	Host    string
	Port int
	List string
}

func NormalFlag() *Args {
    var host string
	flag.StringVar(&host, "host", "", "set host")
	port := flag.Int("port", 0, "set port")
	list := flag.String("list", "0,1,2,3,4,5,6,7", "num list")

	flag.Usage = func() {
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, USAGE, os.Args[0])
	}

	flag.Parse()
    
    return &Args{
        Host: host,
        Port: *port,
		List: *list,
	}
}
