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
	"learn/common"
	_ "learn/common"
	_ "learn/file"
	_ "learn/flags"
	"learn/logger"
	"time"
	_ "time"
)

func main() {
	// fmt.Printf("flag: %#v\n", flags.Main())
    // common.Title(" first title", 0)
	// common.Title("first title", 5)
	// common.Title(" first title", 2)
	// common.Title(" first title", 3)
    common.Exec("ping -c 3 127.0.0.1", time.Second*10, true)
	logger.Info("info log")
	logger.Error("sum log")

	defer logger.Close()
}
