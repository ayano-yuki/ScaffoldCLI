package service

import (
	"fmt"
	"os"
	"path/filepath"
)

type FileWriter interface {
	WriteFile(path string, content string) error
	EnsureDir(dir string) error
}

type fileWriterImpl struct{}

func NewFileWriter() FileWriter {
	return &fileWriterImpl{}
}

// WriteFile 指定パスにファイルを書き出す
func (f *fileWriterImpl) WriteFile(path string, content string) error {
	dir := filepath.Dir(path)
	if err := f.EnsureDir(dir); err != nil {
		return err
	}

	return os.WriteFile(path, []byte(content), 0644)
}

// EnsureDir 存在しないディレクトリを再帰的に作成
func (f *fileWriterImpl) EnsureDir(dir string) error {
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", dir, err)
	}
	return nil
}
