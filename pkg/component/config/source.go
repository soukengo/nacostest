package config

import (
	"fmt"
	"io"
	"os"
	"path"
	"strings"
)

type Source interface {
	From() io.Reader
	ConfigType() string
	Error() error
	Close() error
}

type FileSource struct {
	err        error
	configType string
	file       *os.File
}

func NewFileSource(filePath string) Source {
	configType := strings.TrimPrefix(path.Ext(filePath), ".")
	s := &FileSource{configType: configType}
	fi, err := os.Stat(filePath)
	if err != nil {
		s.err = err
		return s
	}
	if fi.IsDir() {
		err = fmt.Errorf("%s is not a directory", filePath)
		return s
	}
	file, err := os.Open(filePath)
	if err != nil {
		s.err = err
		return s
	}
	s.file = file
	return s
}

func (s *FileSource) From() io.Reader {
	return s.file
}

func (s *FileSource) ConfigType() string {
	return s.configType
}

func (s *FileSource) Error() error {
	return s.err
}
func (s *FileSource) Close() error {
	if s.file != nil {
		return s.file.Close()
	}
	return nil
}
