package buff

import (
	"bufio"
	"fmt"
	"io"
	"log/slog"
	"os"
)

// Reader 按字节读取
func ReadByte(file string) ([]byte, error) {
	fs, err := os.OpenFile(file, os.O_RDONLY, 0666)
	if err!= nil {
		slog.Info(fmt.Sprintf("open file %s error %v\n", file, err))
		return []byte{}, err
	}
	reader, output := bufio.NewReader(fs), []byte{}
	for {
        s, err := reader.ReadBytes('\n')
		if err == io.EOF { break }
		if err != nil {
			slog.Info(fmt.Sprintf("read file %s error %v\n", file, err))
			return []byte{}, err
		}
		fmt.Println(string(s))
		output = append(output, s...)
	}
    return output, nil
}

// Reader 缓存读取输出
func Read(file string) (string, error) {
	fs, err := os.OpenFile(file, os.O_RDONLY, 0666)
	if err != nil {
		slog.Info(fmt.Sprintf("open file %s error %v\n", file, err))
		return fmt.Sprintf("open file %s error %v", file, err), err
	}
	reader, output := bufio.NewReader(fs), ""
	for {
		line, err := reader.ReadString('\n')
        if err == io.EOF { break}
		if err != nil {
			slog.Info(fmt.Sprintf("read file %s error %v\n", file, err))
			return fmt.Sprintf("read file %s error %v", file, err), err
		}
		slog.Info(line)
		output += line + "\n"
	}
	return output, nil
}

// scanner 缓存读取
func Scanner(file string) (string, error) {
	fs, err := os.OpenFile(file, os.O_RDONLY, 0666)
	if err != nil {
		slog.Info(fmt.Sprintf("open file %s error %v\n", file, err))
		return fmt.Sprintf("open file %s error %v", file, err), err
	}
	scanner, output := bufio.NewScanner(fs), ""
	for scanner.Scan() {
		slog.Info(scanner.Text())
        output += scanner.Text() + "\n"
	}
	if scanner.Err() != nil {
		slog.Info(fmt.Sprintf("read file %s error %v\n", file, err))
		return fmt.Sprintf("read file %s error %v", file, err), err
	}
	return output, nil
}


// writer 缓存写入
func Writer(file string, s string) error {
	fs, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err!= nil {
		slog.Info(fmt.Sprintf("open file %s error %v\n", file, err))
	}
	writer := bufio.NewWriter(fs)
	_, err = writer.WriteString("insert string")
	if err != nil { return err }
	return nil
}