package logger_test

import (
	"bufio"
	"os"
	"strings"
	"testing"

	"github.com/Pauloo27/logger"
	"github.com/stretchr/testify/assert"
)

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

	defaultLevelsFunc := []func(...interface{}){logger.Debug, logger.Success, logger.Info, logger.Warn}

	t.Run("simple 'hello' log in all non-error levels", func(t *testing.T) {
		for _, level := range defaultLevelsFunc {
			level("hello")
			assertLog("hello")
		}
	})

	t.Run("simple '10' log in all non-error levels", func(t *testing.T) {
		for _, level := range defaultLevelsFunc {
			level(10)
			assertLog("10")
		}
	})

	// TODO: logf
	// TODO: multiple parameters
	// TODO: custom level
	// TODO: error/fatal
}
