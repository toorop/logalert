package logalert

import (
	"fmt"
	"io"
	"log"
	"sync"
	"time"
)

// Logalert logger
type Logalert struct {
	*sync.Mutex
	infoLogger    *log.Logger
	errLogger     *log.Logger
	alertSenders  []AlertSender
	gracePeriod   time.Duration
	lastAlertSent time.Time
}

// NewLogalert init a new logalert
func NewLogalert(infoWriter, errWritter io.Writer, alertSenders []AlertSender, gracePeriod time.Duration) *Logalert {
	logger := Logalert{
		&sync.Mutex{},
		log.New(infoWriter, "INFO: ", log.Ldate|log.Ltime),
		log.New(errWritter, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
		alertSenders,
		gracePeriod,
	}
	return &logger
}

// SetAlertSenders set alert senders
func (l *Logalert) SetAlertSenders(as []AlertSender) {
	l.Lock()
	defer l.Unlock()
	l.alertSenders = as
}

// SendAlert send alerts via alertSenders
func (l *Logalert) SendAlert(v ...interface{}) {
	l.Lock()
	defer l.Unlock()
	if time.Since(l.lastAlertSent) > l.gracePeriod {
		for _, sender := range l.alertSenders {
			sender.Send(fmt.Sprint(v...))
		}
		l.lastAlertSent = time.Now()
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
