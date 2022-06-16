package main

import (
	"fmt"
	"github/john-80/logf"
	"log"
)

func main() {
	// do something
	logf.Config("./example/log.log", logf.DayMode, logf.DebugLevel)
	logf.Debugf("debug")
	logf.Infof("info")
	logf.Warnf("warn")
	logf.Errorf("error")
	log.Println("print log test")
	fmt.Println("end")
}
