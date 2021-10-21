package logger

import (
	"fmt"
	"os"
	"runtime/debug"
	"strings"
	"time"
)

var (
	Stdout = os.Stdout
	Stderr = os.Stderr
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
	var message strings.Builder

	for _, listener := range listeners {
		listener(level, params)
	}

	now := time.Now().Format("15:04:05")

	// write the prefix
	message.WriteString(fmt.Sprintf("%s%s[%s @ %s]%s ",
		level.Color, ColorBold, level.Name, now, ColorReset,
	))

	message.WriteString(fmt.Sprintln(params...))

	if level.Error {
		fmt.Fprint(Stderr, message.String())
		debug.PrintStack()
	} else {
		fmt.Fprint(Stdout, message.String())
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
