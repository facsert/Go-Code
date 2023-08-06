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
	"os/exec"
)

func Run(cmd string) (output string, err error) {
	fmt.Println(cmd)
	c := exec.Command("bash", "-c", cmd)  

	stdout, err := c.StdoutPipe()
	if err != nil { return output, err }
    c.Stderr = c.Stdout

	go func() {
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			line := scanner.Text()
			fmt.Println(line)
			output += line + "\n"
		}
	}()

	if err = c.Start(); err != nil { return output, err }
	err = c.Wait()
	return output, err
}





