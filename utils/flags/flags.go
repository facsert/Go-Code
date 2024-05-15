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

func FlagInit() (paramMap map[string]string) {
	params := make(map[string]*string, len(paramList))
	for _, param := range paramList {
		value := flag.String(param.Name, param.Default, param.Help)
		params[param.Name] = value
	}
    
	flag.Parse()
	
	paramMap = make(map[string]string, len(paramList))
	for name, value := range params {
		paramMap[name] = *value
	}
	return paramMap
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
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nExample:\n")
		fmt.Fprintf(os.Stderr, "  %s -dir=/home/log  \"set dir\"\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s -output=/home/log -list=1,2,3  \"set output and list\"\n", os.Args[0])
	}

	flag.Parse()
    
    return &Args{
        Host: host,
        Port: *port,
		List: *list,
	}
}
