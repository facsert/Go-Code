/*
 * @Author: facsert
 * @Date: 2023-08-01 21:41:41
 * @LastEditTime: 2023-08-02 22:45:37
 * @LastEditors: facsert
 * @Description:
 */

package main

import (
	_ "fmt"

	"learn/flags"
	"learn/logger"
)

func main() {
	defer logger.Close()
	param := flags.Main()
	logger.Info("param: %#v\n", param)
	logger.Display(true, "Get param success")
}
