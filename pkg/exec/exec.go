/*
* @Author: facsert
* @Date: 2023-08-06 10:39:37
 * @LastEditTime: 2023-08-06 22:11:16
 * @LastEditors: facsert
* @Description:
*/
package exec

import (
	"bufio"
	"log"
	"io"
	"os/exec"
	"time"
	"errors"
	"strings"
)


func Exec(cmd string, timeout time.Duration, view bool) (string, error) {
	log.Println(cmd)
	var done = make(chan error)
	proc := exec.Command("bash", "-c", cmd)

	stdout, err := proc.StdoutPipe()
	proc.Stderr = proc.Stdout
	if err != nil { 
		log.Println("Get stdout failed:", err)
		return "", err
	}

	err = proc.Start()
	if err != nil {
		log.Println("Error starting command:", err)
		return "", err
	}

    var output strings.Builder
	go func() {
		reader := bufio.NewReader(stdout)
		defer close(done)
		for {
			line, err := reader.ReadString('\n')
			if err == io.EOF {
				done <- proc.Wait()
				break
			}
	
			if err != nil {
				done <- err
				break
			}

			if view { log.Print(line) }
			output.WriteString(line)
		}
	}()

	select {
	case err := <-done:
		if err != nil { 
			output.WriteString("\nCommand Run error")
		}
		return output.String(), err
	case <-time.After(timeout):
		output.WriteString("\nTimeout!!")
		log.Print("\nTimeout!!\n")
		proc.Process.Kill()
		return output.String(), errors.New("Timeout!!")
	}
}