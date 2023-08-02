/*
 * @Author: facsert
 * @Date: 2023-08-01 21:46:50
 * @LastEditTime: 2023-08-02 22:49:00
 * @LastEditors: facsert
 * @Description: package flags
 */

package flags

import (
	"flag"
)

var paramList = []map[string]string{
	{
		"name":    "host",
		"default": "127.0.0.1",
		"help":    "server host",
	},
	{
		"name":    "username",
		"default": "root",
		"help":    "server username",
	},
	{
		"name":    "password",
		"default": "admin",
		"help":    "server password",
	},
}

func Main() (paramMap map[string]string) {
	params := make(map[string]*string, len(paramList))
	for _, param := range paramList {
		var value string
		flag.StringVar(&value, param["name"], param["default"], param["help"])
		params[param["name"]] = &value
	}

	flag.Parse()

	paramMap = make(map[string]string, len(paramList))
	for name, value := range params {
		paramMap[name] = *value
	}
	return paramMap
}
