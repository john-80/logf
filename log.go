package logf

import (
	"os"
	"path"
)

type LogFile struct {
	fileFD   *os.File
	fileName string
	fileMode string
	filePath string
	fileSize int64
	level    LEVEL
	prefix   PREFIX
	mode     MODE
	maxSize  int64
}

func NewLogFile(fileName string) *LogFile {
	filePath := path.Dir(fileName)
	os.MkdirAll(filePath, 0777)
	return &LogFile{
		fileName: fileName,
		fileMode: "a",
		filePath: filePath,
		level:    InfoLevel,
		prefix:   InfoPrefix,
		mode:     DayMode,
		maxSize:  1 << 30, // 1GB
	}
}
