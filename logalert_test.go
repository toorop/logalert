package logalert

import (
	"os"
	"testing"
)

// TestNewLogalert test NewLogAlert
func TestLogNoalert(t *testing.T) {
	var logger *Logger

	if logger = NewLogger(os.Stdout, os.Stderr, []AlertSender{}, 15); logger == nil {
		t.Error("expected '*Logger', got 'nil'")
		return
	}
	logger.Info("test info")
}

func TestLogalertPusher(t *testing.T) {
	var logger *Logger
	// check env for pushover token
	appToken := os.Getenv("POAPPTOKEN")
	if appToken == "" {
		panic("POAPPTOKEN not set")
	}
	userToken := os.Getenv("POUSERTOKEN")
	if userToken == "" {
		panic("POUSERTOKEN not set")
	}
	asPushover := AlertSenderPushover{
		UserToken: userToken,
		AppToken:  appToken,
	}
	if logger = NewLogger(os.Stdout, os.Stderr, []AlertSender{asPushover}, 15); logger == nil {
		t.Error("expected '*Logger', got 'nil'")
		return
	}
	logger.ErrorAlert("alert test")

}
