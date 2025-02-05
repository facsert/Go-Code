package main

import (
	"fmt"
	"time"

	"learn/exec"
)


func main() {
    // fmt.Println("learning everyday")
	// num := 0o77
	// fmt.Printf("The value of num is %d\n", num)
	// fmt.Printf("The value of num is %o\n", num)

	// fmt.Printf("The value of num is %.3f\n", 3.1415926)
    // t1 := time.Now()

	// for i := 0; i < 5000; i++ {
		// output, err := exec.Exec2("ping -c 3 127.0.0.1", 6 * time.Second, true)
		// exec.Exec2("uname -a", 6 * time.Second, false)
		// fmt.Println(err)
		// fmt.Println(output)
	// }
	// fmt.Println("Time taken:", time.Since(t1))

	output, err := exec.Exec("uuu", 6 * time.Second, false)
	fmt.Println(err)
	fmt.Println(output)
	// exec.Exec2("ping -c 5 127.0.0.1", 3 * time.Second, true)
}