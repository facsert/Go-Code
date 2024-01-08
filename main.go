/*
 * @Author: facsert
 * @Date: 2023-08-01 21:41:41
 * @LastEditTime: 2023-08-16 22:57:27
 * @LastEditors: facsert
 * @Description:
 */

package main

import (
	"bufio"
	_ "fmt"
	"log/slog"
	"os"
)

func main() {
	fs, _ := os.OpenFile("process.log", os.O_RDWR|os.O_CREATE, 0666)
	scanner := bufio.NewScanner(fs)
	output := ""
	for scanner.Scan() {
		output += scanner.Text() + "\n"
		slog.Info(scanner.Text())
	}
	slog.Info("next line")
	slog.Info("\n" + output)
}
