package logger

import (
	"fmt"
	"log"
	"os"
	"runtime/debug"
	"time"
)

var (
	stdout = log.New(os.Stdout, "", 0)
	stderr = log.New(os.Stderr, "", 0)
)

// levels
var (
	DEBUG   = Level{"DEBUG", ColorBlue, false, false}
	SUCCESS = Level{"SUCCESS", ColorGreen, false, false}
	INFO    = Level{"INFO", ColorWhite, false, false}
	WARN    = Level{"WARN", ColorYellow, false, false}
	ERROR   = Level{"ERROR", ColorRed, true, false}
	FATAL   = Level{"FATAL", ColorRed, true, true}
)

func Log(level Level, params ...interface{}) {
	now := time.Now().Format("15:04:05")
	formatParams := []interface{}{level.Color, ColorBold, level.Name, now, ColorReset}
	formatParams = append(formatParams, params...)

	formattedMsg := fmt.Sprintf(
		"%s%s[%s @ %s]%s %v",
		formatParams...,
	)
	if level.Error {
		stderr.Println(formattedMsg)
		debug.PrintStack()
	} else {
		stdout.Println(formattedMsg)
	}
	if level.Fatal {
		os.Exit(-1)
	}
}

func Logf(level Level, format string, v ...interface{}) {
	Log(level, fmt.Sprintf(format, v...))
}

func Debug(params ...interface{}) {
	Log(DEBUG, params...)
}

func Success(params ...interface{}) {
	Log(SUCCESS, params...)
}

func Info(params ...interface{}) {
	Log(INFO, params...)
}

func Warn(params ...interface{}) {
	Log(WARN, params...)
}

func Error(params ...interface{}) {
	Log(ERROR, params...)
}

func Fatal(params ...interface{}) {
	Log(FATAL, params...)
}

func Debugf(format string, v ...interface{}) {
	Logf(DEBUG, format, v...)
}

func Successf(format string, v ...interface{}) {
	Logf(SUCCESS, format, v...)
}

func Infof(format string, v ...interface{}) {
	Logf(INFO, format, v...)
}

func Warnf(format string, v ...interface{}) {
	Logf(WARN, format, v...)
}

func Errorf(format string, v ...interface{}) {
	Logf(ERROR, format, v...)
}

func Fatalf(format string, v ...interface{}) {
	Logf(FATAL, format, v...)
}

func HandleFatal(err error, msg string) {
	if err == nil {
		return
	}
	Fatalf("%s: %v", msg, err.Error())
}
