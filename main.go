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
	"log/slog"
)

func main() {
	name := "facsert"
    slog.Info(fmt.Sprintf("slog name: %s\n", name))
}
