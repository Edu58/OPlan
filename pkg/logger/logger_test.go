package logger

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func captureOutput(f func()) string {
	// Save original output
	oldStdout := os.Stdout

	// Create pipe to capture output
	r, w, _ := os.Pipe()
	os.Stdout = w

	outputChan := make(chan string)

	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outputChan <- buf.String()
	}()

	// Run func to generate output
	f()

	w.Close()
	os.Stdout = oldStdout

	output := <-outputChan
	return output
}

func TestNewLogger(t *testing.T) {
	logger := NewLogger(os.Stdout)
	assert.NotNil(t, logger)
	assert.IsType(t, &ZeroLogger{}, logger)
}

func TestNewLoggerWithLevel(t *testing.T) {
	level := "info"
	logger := NewLoggerWithLevel(level, os.Stdout)

	assert.NotNil(t, logger)
	assert.IsType(t, &ZeroLogger{}, logger)
	assert.Implements(t, (*Logger)(nil), logger)
	assert.Equal(t, level, logger.GetLevel())
}
