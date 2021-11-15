package main

import (
	"errors"
	"fmt"

	"github.com/getsentry/sentry-go"
)

type Event struct {
	Error map[string]interface{}
	// *Error
	*DSN
}

func (event *Event) setDsn(dsn string) {
	event.DSN = NewDSN(dsn)
	if event.DSN == nil {
		sentry.CaptureException(errors.New("null DSN"))
		fmt.Println("null DSN")
	}
}
