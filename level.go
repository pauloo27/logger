package logger

import "runtime"

type Level struct {
	Name, Color string
	// Error = print to stderr and print stacktrace
	// Fatal = exit the program
	Error, Fatal bool
}

var (
	DEBUG = Level{
		Name: "DEBUG", Error: false, Fatal: false,
	}
	SUCCESS = Level{Name: "SUCCESS", Error: false, Fatal: false}
	INFO    = Level{Name: "INFO", Error: false, Fatal: false}
	WARN    = Level{Name: "WARN", Error: false, Fatal: false}
	ERROR   = Level{Name: "ERROR", Error: true, Fatal: false}
	FATAL   = Level{Name: "FATAL", Error: true, Fatal: true}
)

func init() {
	if runtime.GOOS == "windows" {
		// colors are unix by default
		overwriteColorsToWindows()
	}

	// set colors at runtime, so they are compatible with the running os.
	DEBUG.Color = ColorBlue
	SUCCESS.Color = ColorGreen
	INFO.Color = ColorWhite
	WARN.Color = ColorYellow
	ERROR.Color = ColorRed
	FATAL.Color = ColorRed
}
