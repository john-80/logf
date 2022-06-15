package logf

import "os"

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
