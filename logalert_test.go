package logalert

import (
	"os"
	"testing"
)

// TestNewLogalert test NewLogAlert
func TestLogNoalert(t *testing.T) {
	var err error
	var logger *Logalert

	if logger = NewLogalert(os.Stdout, os.Stderr, []AlertSender{}); logger == nil {
		t.Error("expected '*Logalert', got 'nil'")
		return
	}
	logger.Info("test info", err)

}
