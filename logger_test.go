package logger_test

import (
	"bufio"
	"os"
	"strings"
	"testing"

	"github.com/Pauloo27/logger"
	"github.com/stretchr/testify/assert"
)

type testLogLevel struct {
	Level    logger.Level
	LogFunc  func(...interface{})
	LogfFunc func(string, ...interface{})
}

func TestLogger(t *testing.T) {
	// change the stdout to something we can manage
	r, w, err := os.Pipe()

	assert.Nil(t, err)
	assert.NotNil(t, r)
	assert.NotNil(t, w)

	logger.Stdout = w

	// create a reader to read the logged lines
	reader := bufio.NewReader(r)

	// a function that checks if the output is the "exepect" one
	assertLog := func(expected string) {
		line, err := reader.ReadString('\n')
		line = strings.TrimSuffix(strings.Join(strings.Split(line, " ")[3:], " "), "\n")
		assert.Nil(t, err)
		assert.Equal(t, expected, line)
	}

	defaultLevelsFunc := []testLogLevel{
		{Level: logger.DEBUG, LogFunc: logger.Debug, LogfFunc: logger.Debugf},
		{Level: logger.SUCCESS, LogFunc: logger.Success, LogfFunc: logger.Successf},
		{Level: logger.INFO, LogFunc: logger.Info, LogfFunc: logger.Infof},
		{Level: logger.WARN, LogFunc: logger.Warn, LogfFunc: logger.Warnf},
	}

	t.Run("log 'hello' in all non-error levels", func(t *testing.T) {
		for _, level := range defaultLevelsFunc {
			level.LogFunc("hello")
			assertLog("hello")
		}
	})

	t.Run("log '10' in all non-error levels", func(t *testing.T) {
		for _, level := range defaultLevelsFunc {
			level.LogFunc(10)
			assertLog("10")
		}
	})

	t.Run("logf 'hi steve'", func(t *testing.T) {
		for _, level := range defaultLevelsFunc {
			level.LogfFunc("hi %s", "steve")
			assertLog("hi steve")
		}
	})

	t.Run("logf 'hi im steve and my favorite number is -127'", func(t *testing.T) {
		for _, level := range defaultLevelsFunc {
			level.LogfFunc("hi im %s and my favorite number is %d", "steve", -127)
			assertLog("hi im steve and my favorite number is -127")
		}
	})

	// TODO: logf
	// TODO: multiple parameters
	// TODO: custom level
	// TODO: error/fatal
}
