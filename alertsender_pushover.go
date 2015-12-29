package logalert

import "github.com/thorduri/pushover"

// AlertSenderPushover - pushover
type AlertSenderPushover struct {
	UserToken string
	AppToken  string
}

// Send an alert via pushover
func (a AlertSenderPushover) Send(msg string) error {
	po, err := pushover.NewPushover(a.AppToken, a.UserToken)
	if err != nil {
		return err
	}
	return po.Message(msg)
}
