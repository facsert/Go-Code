/*
 * @Author: facsert
 * @Date: 2023-08-01 21:41:41
 * @LastEditTime: 2023-08-16 22:57:27
 * @LastEditors: facsert
 * @Description:
 */

package main

import (
	"fmt"
	"learn/pkg/comm"
	"log/slog"
)

func main() {
	name := "John"
    slog.Info(fmt.Sprintf("slog name: %s\n", name))
    slog.Info(comm.AbsPath())
	// for index, file := range comm.ListDir("pkg") {
	// 	slog.Info(fmt.Sprintf("%d: %s\n", index, file))
	// }
}
