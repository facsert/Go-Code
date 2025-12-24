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
    go func() {
		defer close(ch)
		scanner := bufio.NewScanner(pipe)
		line := scanner.Text()
		for scanner.Scan() {
			select {
			case <-ctx.Done():
				return
			case ch <- line:
				if c.View {
					slog.Info(line)
				}
			}
		}
		if err := scanner.Err(); err != nil {
			slog.Error("scanner error", "error", err)
		}
	}()
    
	var output strings.Builder
	for {
			select {
			case line, ok := <-ch:
				if ok {
					output.WriteString(line + "\n")
					continue
				}
				goto output
			case <-ctx.Done():
				return output.String(), fmt.Errorf("command timed out")
			}
		}
    
	output:
		if err := cmd.Wait(); err != nil {
			return output.String(), fmt.Errorf("wait: %w", err)
		}
		return output.String(), nil
}
