package logalert

// AlertSender interface for sending alert
type AlertSender interface {
	Send(msg string) error
}
