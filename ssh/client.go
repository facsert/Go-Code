package client

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
    "os"
	"log/slog"
	"time"

	"golang.org/x/crypto/ssh"
)

// go get -u "golang.org/x/crypto/ssh"

type Client struct {
	Host string
	Port int
	Username string
	Password string

    client *ssh.Client
}

func NewClient(host string, port int, username string, password string, timeout time.Duration) (*Client, error) {
    config := &ssh.ClientConfig{
        User:            username,
        Auth:            []ssh.AuthMethod{ssh.Password(password)},
        Timeout:         timeout,
        HostKeyCallback: ssh.InsecureIgnoreHostKey(),
    }
    sshClient, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", host, port), config)
    if err != nil {
        return nil, fmt.Errorf("connect host %s:%d error: %v", host, port, err)
    }
    return &Client{host, port, username, password, sshClient}, nil
}

func (c *Client) Exec(command string, timeout time.Duration) (string, error) {
    slog.Info(command)
    session, err := c.client.NewSession()
	if err != nil {
        return "", fmt.Errorf("create session error: %v", err)
	}
    defer session.Close()

    ctx, cancel := context.WithTimeout(context.Background(), timeout)
    defer cancel()
    
    done := make(chan error)
    defer close(done)
    var output []byte
    go func() {
        output, err = session.CombinedOutput(command)
        done <- err
    }()

    select {
    case err := <- done:
        slog.Info(string(output))
        return string(output), err
    case <- ctx.Done():
        slog.Info(string(output))
        slog.Error(fmt.Sprintf("Run %s timeout", command))
        return string(output), fmt.Errorf("Run %s timeout", command)
    }
}

func (c *Client) Run(command string, timeout time.Duration) (string, error) {
    slog.Info(command)
    session, err := c.client.NewSession()
	if err != nil { return "", fmt.Errorf("create session error: %v", err) }
	defer session.Close()

	stdout, err := session.StdoutPipe()
	if err != nil { return "", fmt.Errorf("failed to get stdout: %v", err) }    
	stderr, err := session.StderrPipe()
	if err != nil { return "", fmt.Errorf("failed to get stderr: %v", err) }

    ctx, cancel := context.WithTimeout(context.Background(), timeout)
    defer cancel()

	if err := session.Start(command); err != nil { 
        return "", fmt.Errorf("start command failed: %v", err)
    }

    var buff bytes.Buffer
    done := make(chan error, 1)
	go func() {
        stdoutReader := io.TeeReader(stdout, &buff)
        stderrReader := io.TeeReader(stderr, &buff)

        go io.Copy(os.Stdout, stdoutReader)
        go io.Copy(os.Stderr, stderrReader)
        done <- session.Wait()
	}()

    select {
    case err := <- done:
        return buff.String(), err
    case <- ctx.Done():
        slog.Error(fmt.Sprintf("Run %s timeout", command))
        return buff.String(), errors.New("timeout")
    }
}