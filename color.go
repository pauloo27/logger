package logger

var (
	ColorBold   = "\033[1m"
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorWhite  = "\033[39m"
)

func overwriteColorsToWindows() {
	ColorBold = "" // windows does not support bold
	ColorReset = "\033[0m"
	ColorRed = "\033[31m"
	ColorGreen = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue = "\033[34m"
	ColorWhite = "\033[37m"
}
