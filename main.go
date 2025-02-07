package main

import (
	"fmt"
	"log/slog"

	"learn/comm"
)

func init() {
	comm.Init()
}

func main() {
    fmt.Println("learning everyday")
	slog.Info("test logger")
}