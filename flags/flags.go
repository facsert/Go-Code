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

func Main() (paramMap map[string]string) {
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
