package logalert

// AlertSender interafce for sending alert
type AlertSender interface {
	Send(msg string) error
}
