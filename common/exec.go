/*
* @Author: facsert
* @Date: 2023-08-06 10:39:37
 * @LastEditTime: 2023-08-06 22:11:16
 * @LastEditors: facsert
* @Description:
*/
package common

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"time"
)

func Exec(cmd string, timeout time.Duration, view bool) (string, error) {
	fmt.Println(cmd)
	proc := exec.Command("bash", "-c", cmd)

	stdout, err := proc.StdoutPipe()
	proc.Stderr = proc.Stdout
	if err != nil { 
		fmt.Println("Get stdout failed:", err)
		return "", err
	}

	err = proc.Start()
	if err != nil {
		fmt.Println("Error starting command:", err)
		return "", err
	}

	reader, output := bufio.NewReader(stdout), ""
    done := make(chan error)
    
	go func() {
		defer close(done)
		for {
			line, err := reader.ReadString('\n')
			if view { fmt.Print(line) }
			output += line
	
			if err != nil || io.EOF == err {
				done <- err
				break
			}
		}
	}()

	select {
	case err := <-done:
		if err != nil { output += "\nCommand Run error" }
		return output, nil
	case <-time.After(timeout):
		output += "\nTimeout!!"
		fmt.Print("\nTimeout!!\n")
		proc.Process.Kill()
		return output, nil
	}
}




