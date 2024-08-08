/*
* @Author: facsert
* @Date: 2023-08-06 10:39:37
 * @LastEditTime: 2023-08-06 22:11:16
 * @LastEditors: facsert
* @Description:
*/
package exec

import (
	"fmt"
	"bufio"
	"io"
	"log/slog"
	"os/exec"
	"strings"
	"time"
	"context"
)


func Exec(cmd string, timeout time.Duration, view bool) (string, error) {
	slog.Info(cmd)
	proc := exec.Command("bash", "-c", cmd)

	stdout, err := proc.StdoutPipe()
	if err != nil { 
		return "", fmt.Errorf("get stdout failed: %w", err)
	}

	proc.Stderr = proc.Stdout

	err = proc.Start()
	if err != nil {
		return "", fmt.Errorf("start command error: %w", err)
	}

	ctx, cancle := context.WithTimeout(context.Background(), timeout)
	defer cancle()
	
	var output strings.Builder
	var done = make(chan error, 1)
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

			if view { slog.Info(line) }
			output.WriteString(line)
		}
	}()

	select {
	case err := <-done:
		if err != nil { 
			slog.Error(fmt.Sprintf("%s Run error\n", cmd))
		}
		return output.String(), err
	case <- ctx.Done():
		slog.Error("Timeout!!\n")
		proc.Process.Kill()
		return output.String(), fmt.Errorf("timeout")
	}
}