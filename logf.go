package logf

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

type LogFile struct {
	fileFD     *os.File
	fileName   string
	filePath   string
	fileSize   int64
	createDate time.Time
	days       int64
	level      LEVEL
	prefix     string
	mode       MODE
	maxSize    int64
}

// Write(p []byte) (n int, err error)
func (l *LogFile) Write(buf []byte) (n int, err error) {
	if l.fileName == "" {
		fmt.Printf("%s\n", buf)
		return len(buf), nil
	}

	switch l.mode {
	case SizeMode: // file size
		fileInfo, err := os.Stat(l.fileName)
		if err != nil {
			l.createLogFile()
			break
		}
		fileSize := fileInfo.Size()
		if l.fileFD == nil || fileSize > l.fileSize {
			l.createLogFile()
		}
	default: // days
		if l.fileFD == nil {
			l.createLogFile()
		} else if l.isAfterDate() {
			l.fileFD.Sync()
			l.fileFD.Close()
			l.createLogFile()
		}
	}

	if l.fileFD == nil {
		fmt.Println("log fileFD is nil!")
		return len(buf), nil
	}

	return l.fileFD.Write(buf)
}

func (l *LogFile) createLogFile() {
	now := time.Now().UTC().Local()
	filename := fmt.Sprintf("%s_%04d%02d%02d_%02d%02d",
		l.fileName, now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute())
	if err := os.Rename(l.fileName, filename); err == nil {
		go l.tarAndDel(filename)
	}

	for index := 0; index < 10; index++ {
		fd, err := os.OpenFile(l.fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModeExclusive)
		if err == nil {
			l.fileFD.Sync()
			l.fileFD.Close()
			l.fileFD = fd
			break
		}

		fmt.Println("open log file error! err:", err)
		l.fileFD = nil
	}
}

func (l *LogFile) tarAndDel(filename string) {
	tarCmd := exec.Command("tar", "-zcf", filename+".tar.gz", filename, "--remove-files")
	tarCmd.Run()

	delCmd := exec.Command("/bin/bash", "-c",
		"find "+l.filePath+" -type f -mtime +"+fmt.Sprint("%ld", l.days)+` -exec -rm {} \;`)
	delCmd.Run()
}

// check now date is after create date
// is yes, return true; is no, return false
func (l *LogFile) isAfterDate() bool {
	now := time.Now().UTC().Local()

	if l.createDate.IsZero() {
		l.createDate = now
		return false
	}

	if now.Year() > l.createDate.Year() {
		return true
	}
	if now.Year() < l.createDate.Year() {
		return false
	}
	if now.Month() > l.createDate.Month() {
		return true
	}
	if now.Month() < l.createDate.Month() {
		return false
	}
	if now.Day() > l.createDate.Day() {
		return true
	}
	if now.Day() < l.createDate.Day() {
		return false
	}
	return false
}
