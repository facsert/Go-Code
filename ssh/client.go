package client

import (
	"fmt"
	"io"
	"log/slog"
	"strings"

	"bufio"
	"sync"
	"time"

	"golang.org/x/crypto/ssh"
)

var wg sync.WaitGroup

type Client struct {
	Host     string
	Port     string
	Username string
	Password string

	TimeOut time.Duration

	SSHClient  *ssh.Client
	SSHSession *ssh.Session

	StdinPipe  io.WriteCloser
	StdoutPipe io.Reader

	mu sync.Mutex
}

func NewClient(host, port, username, password string, timeout time.Duration) (*Client, error) {
	client := &Client{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
		TimeOut:  timeout,
	}
	err := client.Connect()
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (c *Client) Connect() error {
	config := &ssh.ClientConfig{
		User: c.Username,
		Auth: []ssh.AuthMethod{
			ssh.Password(c.Password), // 或使用公钥认证
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // 生产环境中应验证 HostKey
		Timeout:         c.TimeOut,
	}

	// 建立 SSH 连接
	conn, err := ssh.Dial("tcp", fmt.Sprintf("%s:%s", c.Host, c.Port), config)
	if err != nil {
		return fmt.Errorf("failed to dial: %s", err)
	}
	c.SSHClient = conn

	// 创建 SSH 会话
	session, err := conn.NewSession()
	if err != nil {
		return fmt.Errorf("failed to create session: %s", err)
	}
	c.SSHSession = session

	// 请求伪终端（PTY）以支持交互式命令（如 sudo）
	modes := ssh.TerminalModes{
		ssh.ECHO:          0,     // 启用输入回显
		ssh.TTY_OP_ISPEED: 14400, // 输入速度
		ssh.TTY_OP_OSPEED: 14400, // 输出速度
	}

	// 设置终端类型和窗口大小（行数、列数）
	if err := session.RequestPty("xterm", 80, 40, modes); err != nil {
		c.Close()
		return fmt.Errorf("failed to request PTY: %s", err)
	}

	stdinPipe, err := session.StdinPipe()
	if err != nil {
		c.Close()
		return fmt.Errorf("failed to get stdin pipe: %s", err)
	}
	c.StdinPipe = stdinPipe

	// 创建管道以捕获输出
	stdoutPipe, err := session.StdoutPipe()
	if err != nil {
		c.Close()
		return fmt.Errorf("failed to get stdout pipe: %w", err)
	}

	stderrPipe, err := session.StderrPipe()
	if err != nil {
		c.Close()
		return fmt.Errorf("failed to get stderr pipe: %w", err)
	}

	c.StdoutPipe = io.MultiReader(stdoutPipe, stderrPipe)

	// 启动远程 Shell
	if err := session.Shell(); err != nil {
		c.Close()
		return fmt.Errorf("failed to start shell: %w", err)
	}

	return nil
}

func (c *Client) Run(cmd string, respTimeout, waitTimeout int, expects, errors []string, view bool) (string, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if view {
		slog.Info(cmd)
	}

	var output strings.Builder
	lines := make(chan string)

	go func() {
		defer close(lines)

		reader := bufio.NewReader(c.StdoutPipe)
		for {
			line, err := reader.ReadString('\n')
			if err != nil || io.EOF == err {
				break
			}
			lines <- line
		}
	}()

	// 发送命令到远程 Shell
	_, err := fmt.Fprintf(c.StdinPipe, "\n%s\n", cmd)
	if err != nil {
		return "", fmt.Errorf("command write failed: %w", err)
	}

	respTimer := time.NewTimer(time.Duration(respTimeout) * time.Second)
	waitTimer := time.NewTimer(time.Duration(waitTimeout) * time.Second)

	for {
		select {
		case line, ok := <-lines:
			output.WriteString(line)
			respTimer.Reset(time.Duration(respTimeout) * time.Second)

			if !ok {
				return output.String(), fmt.Errorf("channel closed")
			}

			if view {
				// slog.Info(fmt.Sprint(line))
				fmt.Print(line)
			}

			for _, expect := range expects {
				if strings.Contains(line, expect) {
					return output.String(), nil
				}
			}
			for _, error := range errors {
				if strings.Contains(line, error) {
					return output.String(), fmt.Errorf("error: %s", error)
				}
			}
		case <-respTimer.C:
			slog.Error("response timeout", "response timeout",respTimeout)
			return output.String(), fmt.Errorf("response timeout")
		case <-waitTimer.C:
			slog.Error("wait timeout", "wait timeout",waitTimeout)
			return output.String(), fmt.Errorf("wait timeout")
		}
	}
}

func (c *Client) Exec(cmd string, waitTimeout int, view bool) (string, error) {
	if view {
		slog.Info(cmd)
	}
	// 创建 SSH 会话
	session, err := c.SSHClient.NewSession()
	if err != nil {
		return "", fmt.Errorf("failed to create session: %s", err)
	}
	defer session.Close()

	stdout, err := session.StdoutPipe()
	// stderr, err = session.StderrPipe()
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	stderr, err := session.StdoutPipe()
	// stderr, err = session.StderrPipe()
	if err != nil {
		fmt.Println(err)
		return "", err
	}
    
	var output strings.Builder

	session.Start(cmd)
	reader := bufio.NewReader(io.MultiReader(stdout, stderr))

	//实时循环读取输出流中的一行内容
	for {
		line, err := reader.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		output.WriteString(line)
		slog.Info(line)
		// fmt.Print(line)
	}

	//阻塞直到该命令执行完成，该命令必须是被Start方法开始执行的
	return output.String(), session.Wait()
}

func (c *Client) Close() {
	if c.SSHSession != nil {
		c.SSHSession.Close()
	}
	if c.SSHClient != nil {
		c.SSHClient.Close()
	}
	if c.StdinPipe != nil {
		c.StdinPipe.Close()
	}
}
