package logger_test

import (
	"testing"

	"github.com/Pauloo27/logger"
)

func TestLogger(t *testing.T) {
	logger.Debug(10)
	logger.Debug("Hello world")
	logger.Success("Hello world")
	logger.Info("Hello world")
	logger.Warn("Hello world")
	logger.Error("Hello world")
	// do not call fatal or the test fails =(
	//logger.Fatal("Hello world")
}
