package logf

import (
	"fmt"
	"log"
	"os"
	"path"
)

var logfile LogFile = LogFile{
	fileSize: 1 << 20, // 1MB
	days:     3,
	level:    InfoLevel,
	prefix:   InfoPrefix,
	mode:     DayMode,
	maxSize:  1 << 30, // 1GB
}

/* func init() {
	logfile.fileFD = os.Stdout
} */

func Config(fileName string, mode MODE, level LEVEL) {
	logfile.filePath = path.Dir(fileName)
	os.MkdirAll(logfile.filePath, os.ModePerm) //0755)

	logfile.fileName = fileName
	SetLevel(level)
	SetMode(mode)

	log.SetOutput(&logfile)
	log.SetFlags(log.Llongfile | log.Ldate | log.Ltime)
}

func SetLevel(level LEVEL) {
	logfile.level = level
}

func SetMode(mode MODE) {
	logfile.mode = mode
}

func SetSize(size int64) {
	if size > logfile.maxSize {
		logfile.fileSize = logfile.maxSize
		return
	}
	logfile.fileSize = size
}

func Debugf(format string, args ...interface{}) {
	if logfile.level >= DebugLevel {
		if logfile.fileFD == os.Stdout {
			log.SetPrefix(blue(DebugPrefix))
		} else {
			log.SetPrefix(DebugPrefix)
		}
		log.Output(2, fmt.Sprintf(format, args...))
	}
}

func Infof(format string, args ...interface{}) {
	if logfile.level >= InfoLevel {
		if logfile.fileFD == os.Stdout {
			log.SetPrefix(green(InfoPrefix))
		} else {
			log.SetPrefix(InfoPrefix)
		}
		log.Output(2, fmt.Sprintf(format, args...))
	}
}

func Warnf(format string, args ...interface{}) {
	if logfile.level >= WarnLevel {
		if logfile.fileFD == os.Stdout {
			log.SetPrefix(yellow(WarnPrefix))
		} else {
			log.SetPrefix(WarnPrefix)
		}
		log.Output(2, fmt.Sprintf(format, args...))
	}
}

func Errorf(format string, args ...interface{}) {
	if logfile.level >= ErrorLevel {
		if logfile.fileFD == os.Stdout {
			log.SetPrefix(red(ErrorPrefix))
		} else {
			log.SetPrefix(ErrorPrefix)
		}
		log.Output(2, fmt.Sprintf(format, args...))
	}
}
