package log

import (
	"fmt"
)

// Log acts a fake demo logger
type Log struct{}

// New returns an empty instance od logger
func New() (*Log, error) {
	return &Log{}, nil
}

// NewDummyLog pretends to setup Log in a different maner
func NewDummyLog() (*Log, error) {
	return &Log{}, nil
}

// Log pretends it does something
func (l *Log) Log(s string) {
	fmt.Println(s)
	return
}
