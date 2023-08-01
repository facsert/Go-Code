package main

import (
	"fmt"
	"learn/flags"
)

func main() {
	fmt.Println("hello world")
	param := flags.Main()
	fmt.Printf("%#v\n", param)
}