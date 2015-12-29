package logalert

import (
	"fmt"
	"io"
	"log"
)

// Logalert logger
type Logalert struct {
	infoLogger   *log.Logger
	errLogger    *log.Logger
	alertSenders []AlertSender
}

// NewLogalert init a new logalert
func NewLogalert(infoWriter, errWritter io.Writer, alertSenders []AlertSender) *Logalert {
	logger := Logalert{
		infoLogger:   log.New(infoWriter, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		errLogger:    log.New(errWritter, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
		alertSenders: alertSenders,
	}
	return &logger
}

// SendAlert send alerts via alertSenders
func (l *Logalert) SendAlert(v ...interface{}) {
	for _, sender := range l.alertSenders {
		sender.Send(fmt.Sprint(v...))
	}
}

// Info log at INFO level
func (l *Logalert) Info(v ...interface{}) {
	l.infoLogger.Println(v...)
	return
}

// InfoAlert log at info level and send alert
func (l *Logalert) InfoAlert(v ...interface{}) {
	l.Info(v...)
	l.SendAlert(v...)
	return
}

// Error write error log
func (l *Logalert) Error(v ...interface{}) {
	l.errLogger.Println(v...)
	return
}

// ErrorAlert write error log and send alert
func (l *Logalert) ErrorAlert(v ...interface{}) {
	l.Error(v...)
	l.SendAlert(v...)
	return
}
