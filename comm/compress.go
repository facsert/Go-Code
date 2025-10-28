package comm

import (
	"archive/zip"
	"fmt"
	"io"
    "path/filepath"
	"os"
)

// 递归搜索 src 路径下文件打包生成 dst
func CompressZip(src, dst string) error {
    zipfile, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("create zipfile failed: %w", err)
	}
	defer zipfile.Close()

	zw := zip.NewWriter(zipfile)
	defer zw.Close()

    return filepath.WalkDir(src, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		r, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
        if r == "." {
			return nil
		}
		r = filepath.ToSlash(r)
        
		info , err := d.Info()
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return fmt.Errorf("create file header failed: %w", err)
		}
		header.Name = r

		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}
        
		w, err := zw.CreateHeader(header)
		if err != nil {
			return fmt.Errorf("create zip head failed: %w", err)
		}
        
		f, err := os.Open(path)
		if err != nil {
			return fmt.Errorf("open %s failed: %w", path, err)
		}
		f.Close()
        
		buffer := make([]byte, 32*1024)
		_, err = io.CopyBuffer(w, f, buffer)
		if err != nil {
			return fmt.Errorf("write zip failed: %w", err)
		}

		return nil
	})
}