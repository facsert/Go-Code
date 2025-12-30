package buff

import (
	"bufio"
	"fmt"
	"io"
	// "log/slog"
	"os"
	"strings"
)

// Reader 缓存读取输出
func Read(file string) (string, error) {
	fs, err := os.OpenFile(file, os.O_RDONLY, 0666)
	if err != nil {
		return "", fmt.Errorf("open file error %w", err)
	}
	defer fs.Close()

	reader, output := bufio.NewReader(fs), make([]string, 0, 100)
	for {
		line, err := reader.ReadString('\n')
        if err == io.EOF { break }
		if err != nil {
			return "", fmt.Errorf("read file error %w", err)
		}

		// slog.Info(line)
		output = append(output, line)
	}
	return strings.Join(output, "\n"), nil
}

// scanner 缓存读取
func Scanner(file string) (string, error) {
	fs, err := os.OpenFile(file, os.O_RDONLY, 0666)
	if err != nil {
		return "", fmt.Errorf("open file error %w", err)
	}
	defer fs.Close()

	scanner, output := bufio.NewScanner(fs), make([]string, 0, 100)
	for scanner.Scan() {
		output = append(output, scanner.Text())
	}
	if scanner.Err() != nil {
		return "", fmt.Errorf("read file error %w", err)
	}

	return strings.Join(output, "\n"), nil
}


// writer 缓存写入
func Writer(file string, s string) error {
	fs, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err!= nil {
		return fmt.Errorf("open file error %w", err)
	}
	defer fs.Close()

	writer := bufio.NewWriter(fs)
	_, err = writer.WriteString(s)
	if err != nil { 
		return fmt.Errorf("write string error %w", err) 
	}
	return nil
}