package logger

import (
	"fmt"
	"strconv"
)

func AsHexRGB(hex string) string {
	r, _ := strconv.ParseInt(hex[0:2], 16, 64)
	g, _ := strconv.ParseInt(hex[2:4], 16, 64)
	b, _ := strconv.ParseInt(hex[4:6], 16, 64)
	return AsRGB(r, g, b)
}

func AsRGB(r, g, b int64) string {
	return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
}
