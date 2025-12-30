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


type Command struct {
    Command string
	Timeout time.Duration
	View    bool
}

func NewCmd(cmd string) *Command {
	return &Command{
		Command: cmd,
		Timeout: 10 * time.Second,
		View:    true,
	}
}

func (c *Command) SetTimeout(timeout time.Duration) *Command {
	c.Timeout = timeout
	return c
}

func (c *Command) SetView(view bool) *Command {
	c.View = view
	return c
}

func (c *Command) read(ctx context.Context, reader io.Reader, ch chan string) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		select {
		case <-ctx.Done():
			return
		default:
			line := scanner.Text()
			if c.View {
				slog.Info(line)
			}
			ch <- line
		}
	}
}

func (c *Command) Run() (string, error) {
	slog.Info(c.Command)
	ctx, cancel := context.WithTimeout(context.Background(), c.Timeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, "bash", "-c", c.Command)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", fmt.Errorf("stdout pipe: %w", err)
	}
    stderr, err := cmd.StderrPipe()
    if err != nil {
        return "", fmt.Errorf("stderr pipe: %w", err)
    }
	pipe := io.MultiReader(stdout, stderr)
	if err := cmd.Start(); err != nil {
		return "", fmt.Errorf("start command: %w", err)
	}
    
	ch := make(chan string, 100)
	var output strings.Builder
	scanner := bufio.NewScanner(pipe)
	
	go func() {
		defer close(ch)
		for scanner.Scan() {
			line := scanner.Text()
			select {
			case <-ctx.Done():
				return
			case ch <- line:
				if c.View {
					slog.Info(line)
				}
			}
		}
	}()

    
	for {
		select {
		case line, ok := <-ch:
			if !ok {
				goto finish
			}
			output.WriteString(line)
		case <-ctx.Done():
			slog.Error(fmt.Sprintf("command timed out error: %v\n", ctx.Err()))
			goto finish
		}
	}
	
	finish:
		for line := range ch {
			output.WriteString(line)
		}

		if err := scanner.Err(); err != nil {
			slog.Error("scanner error", "error", err)
		}

		if err := cmd.Wait(); err != nil {
			return output.String(), fmt.Errorf("wait: %w", err)
		}
		return output.String(), nil
}



func Exec(cmdStr string, timeout time.Duration, view bool) (string, error) {
	slog.Info(cmdStr)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, "bash", "-c", cmdStr)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", fmt.Errorf("stdout pipe: %w", err)
	}
	
    stderr, err := cmd.StderrPipe()
    if err != nil {
        return "", fmt.Errorf("stderr pipe: %w", err)
    }
	pipe := io.MultiReader(stdout, stderr)

	if err := cmd.Start(); err != nil {
		return "", fmt.Errorf("start command: %w", err)
	}

	var output strings.Builder
	lines := make(chan string)
	scanDone := make(chan struct{})

	// 启动读取goroutine
	go func() {
		defer close(lines)
		reader := bufio.NewReader(pipe)
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