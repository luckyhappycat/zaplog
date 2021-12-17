package zlog

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestSampleLogger(t *testing.T) {
	Info("root msg")
	Debug("debug info")

	l := CreateLogger("mod1")
	l.Info("mod1 msg")
	l.Debug("mod1 debug")
}

func Test_getLoggerLevelByName(t *testing.T) {
	_ = os.Setenv("LOG_LEVEL", "info")
	assert.Equal(t, zap.InfoLevel, getLoggerLevelByName(""))
	assert.Equal(t, zap.InfoLevel, getLoggerLevelByName("mod1"))

	_ = os.Setenv("LOG_LEVEL", "debug")
	assert.Equal(t, zap.DebugLevel, getLoggerLevelByName(""))
	assert.Equal(t, zap.DebugLevel, getLoggerLevelByName("mod1"))

	_ = os.Setenv("LOG_LEVEL", "debug,mod1=error")
	assert.Equal(t, zap.DebugLevel, getLoggerLevelByName(""))
	assert.Equal(t, zap.ErrorLevel, getLoggerLevelByName("mod1"))
	assert.Equal(t, zap.DebugLevel, getLoggerLevelByName("mod2"))

	_ = os.Setenv("LOG_LEVEL", "mod1=error")
	assert.Equal(t, zap.InfoLevel, getLoggerLevelByName(""))
	assert.Equal(t, zap.ErrorLevel, getLoggerLevelByName("mod1"))
	assert.Equal(t, zap.InfoLevel, getLoggerLevelByName("mod2"))
}
