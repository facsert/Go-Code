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


func Exec(cmdStr string, timeout time.Duration, view bool) (string, error) {
	slog.Info(cmdStr)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, "bash", "-c", cmdStr)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", fmt.Errorf("stdout pipe: %w", err)
	}
	cmd.Stderr = cmd.Stdout // 合并标准错误到标准输出

	if err := cmd.Start(); err != nil {
		return "", fmt.Errorf("start command: %w", err)
	}

	var output strings.Builder
	lines := make(chan string)
	scanDone := make(chan struct{})

	// 启动读取goroutine
	go func() {
		defer close(lines)
		reader := bufio.NewReader(stdout)
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				if err == io.EOF && line != "" {
					lines <- line
				}
				break
			}
			lines <- line
		}
		close(scanDone)
	}()

	var cmdErr error
	loop:
		for {
			select {
			case line, ok := <-lines:
				if !ok {
					break loop
				}
				if view {
					slog.Info(strings.TrimSuffix(line, "\n"))
				}
				output.WriteString(line)
			case <-scanDone:
				break loop
			case <-ctx.Done():
				cmdErr = ctx.Err()
				slog.Error("command timed out")
				break loop
			}
		}

	// 处理剩余数据
	for line := range lines {
		if view {
			slog.Info(strings.TrimSuffix(line, "\n"))
		}
		output.WriteString(line)
	}

	// 获取最终命令状态
	err = cmd.Wait()
	if cmdErr == nil {
		cmdErr = err
	}

	// 优先返回超时错误
	if ctx.Err() == context.DeadlineExceeded {
		return output.String(), fmt.Errorf("timeout: %w", ctx.Err())
	}

	if cmdErr != nil {
		return output.String(), fmt.Errorf("command failed: %w", cmdErr)
	}
	return output.String(), nil
}