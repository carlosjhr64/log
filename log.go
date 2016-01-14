package log

import (
  "io"
  "os"
  "fmt"
  "sync"
)

var mutex = &sync.Mutex{}
var OutFormat = ""
var OutWriter = os.Stdout
var ErrFormat = ""
var ErrWriter = os.Stderr

func p(s io.Writer, f string, a ...interface{}) {
  mutex.Lock()
  if f == "" {
    fmt.Fprintln(s, a...)
  } else {
    fmt.Fprintf(s, f, a...)
  }
  mutex.Unlock()
}

func Out(a ...interface{}) {
  p(OutWriter, OutFormat, a...)
}

func Err(a ...interface{}) {
  p(ErrWriter, ErrFormat, a...)
}
