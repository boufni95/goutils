package goutils

import (
	"fmt"
	"strings"
	"time"
	"io"

	"github.com/davecgh/go-spew/spew"
)

var logger loggerType

type loggerType struct {
	level LogType
}

type LogType int

const (
	DEBUG   LogType = 0
	WARNING LogType = 1
	ERROR   LogType = 2
)

func InitLogger(level string) {
	level = strings.ToLower(level)
	var l LogType
	switch level {
	case "error":
		l = ERROR
	case "warning":
		l = WARNING
	case "debug":
		l = DEBUG
	default:
		l = DEBUG
	}
	logger = loggerType{
		level: l,
	}

}

func PrintError(err error) {
	fmt.Println("ERROR at", time.Now())
	fmt.Println(err)
}

func Warn(s string) {
	if logger.level <= WARNING {
		fmt.Println("WARNING at", time.Now())
		fmt.Println(s)
	}
}

func Log(s string) {
	if logger.level == 0 {
		fmt.Println("LOG at", time.Now())
		fmt.Println(s)
	}
}

func Dump(v interface{}) {
	spew.Dump(v)
}

func Fdump(w io.Writer,v interface{}) {
	spew.Fdump(w,v)
}