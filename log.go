/*
Very simple logger with mutex.
*/
package log

import (
  "os"
  "fmt"
  "sync"
)

const VERSION = "0.1.0"

type Format string

var mutex = &sync.Mutex{}
var OutWriter = os.Stdout
var ErrWriter = os.Stderr

func Out(a ...interface{}) {
  mutex.Lock()
  fmt.Fprintln(OutWriter, a...)
  mutex.Unlock()
}

func Err(a ...interface{}) {
  mutex.Lock()
  fmt.Fprintln(ErrWriter, a...)
  mutex.Unlock()
}

func New(format string) *Format {
  f := Format(format)
  return &f
}

func (f *Format) Out(a ...interface{}) {
  mutex.Lock()
  fmt.Fprintf(OutWriter, string(*f), a...)
  mutex.Unlock()
}

func (f *Format) Err(a ...interface{}) {
  mutex.Lock()
  fmt.Fprintf(ErrWriter, string(*f), a...)
  mutex.Unlock()
}
