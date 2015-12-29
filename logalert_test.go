package logalert

import (
	"os"
	"testing"
)

// TestNewLogalert test NewLogAlert
func TestLogNoalert(t *testing.T) {
	var logger *Logalert

	if logger = NewLogalert(os.Stdout, os.Stderr, []AlertSender{}); logger == nil {
		t.Error("expected '*Logalert', got 'nil'")
		return
	}
	logger.Info("test info")
}

func TestLogalertPusher(t *testing.T) {
	var logger *Logalert
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
	if logger = NewLogalert(os.Stdout, os.Stderr, []AlertSender{asPushover}); logger == nil {
		t.Error("expected '*Logalert', got 'nil'")
		return
	}
	logger.ErrorAlert("alert test")

}
