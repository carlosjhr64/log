package log

import (
  "io"
  "os"
  "fmt"
  "sync"
)

var mutex = &sync.Mutex{}
var Format = ""
var OutWriter = os.Stdout
var ErrWriter = os.Stderr

func p(s io.Writer, a ...interface{}) {
  mutex.Lock()
  if Format == "" {
    fmt.Fprintln(s, a...)
  } else {
    fmt.Fprintf(s, Format, a...)
  }
  mutex.Unlock()
}

func Out(a ...interface{}) {
  p(OutWriter, a...)
}

func Err(a ...interface{}) {
  p(ErrWriter, a...)
}
