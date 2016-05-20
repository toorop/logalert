package logalert

import (
	"fmt"
	"io"
	"log"
	"sync"
	"time"
)

// Logger is the logalert logger
type Logger struct {
	*sync.Mutex
	infoLogger    *log.Logger
	errLogger     *log.Logger
	alertSenders  []AlertSender
	gracePeriod   time.Duration
	lastAlertSent time.Time
}

// NewLogalert init a new logalert
func NewLogger(infoWriter, errWritter io.Writer, alertSenders []AlertSender, gracePeriodSeconds int) *Logger {
	logger := Logger{
		&sync.Mutex{},
		log.New(infoWriter, "INFO: ", log.Ldate|log.Ltime),
		log.New(errWritter, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
		alertSenders,
		time.Duration(gracePeriodSeconds) * time.Second,
		time.Unix(0, 0),
	}
	return &logger
}

// SetAlertSenders set alert senders
func (l *Logger) SetAlertSenders(as []AlertSender) {
	l.Lock()
	defer l.Unlock()
	l.alertSenders = as
}

// SendAlert send alerts via alertSenders
func (l *Logger) SendAlert(v ...interface{}) {
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
func (l *Logger) Info(v ...interface{}) {
	l.infoLogger.Println(v...)
	return
}

// InfoAlert log at info level and send alert
func (l *Logger) InfoAlert(v ...interface{}) {
	l.Info(v...)
	l.SendAlert(v...)
	return
}

// Error write error log
func (l *Logger) Error(v ...interface{}) {
	l.errLogger.Println(v...)
	return
}

// ErrorAlert write error log and send alert
func (l *Logger) ErrorAlert(v ...interface{}) {
	l.Error(v...)
	l.SendAlert(v...)
	return
}
