package main

import (
	"fmt"
	"learn/utils/datetime"
	"learn/utils/logger"
	"learn/utils/comm"
)


func init() {
    logger.Init()
}

func TestDateTime() {
	dateTime.Test()
}

func main() {
    // fmt.Println(comm.AbsPath(""))
	for _, file := range comm.ListDir(comm.AbsPath("")) {
		fmt.Println(file)
	}
}
