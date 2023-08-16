/*
 * @Author: facsert
 * @Date: 2023-08-01 21:41:41
 * @LastEditTime: 2023-08-16 22:57:27
 * @LastEditors: facsert
 * @Description:
 */

package main

import (
	_ "fmt"
    _ "learn/common"
	_ "learn/flags"
	"learn/logger"
	_ "learn/file"
	_ "time"
)

func main() {
	// fmt.Printf("flag: %#v\n", flags.Main())
    // common.Title(" first title", 0)
	// common.Title("first title", 5)
	// common.Title(" first title", 2)
	// common.Title(" first title", 3)

	logger.Info("info log")
	logger.Error("error log")

	logger.SetLogOutput("summary.log")

	logger.Info("summary log")
	logger.Error("sum log")

	defer logger.Close()
}
