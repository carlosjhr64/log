package main

import "github.com/carlosjhr64/log"

func main() {
  log.Out("OK", "This is to Stdout")
  log.Err("Oops", "This is to Stderr")

  log.ErrFormat = "Ooops: %s.\n"
  log.OutFormat = "OK: %s.\n"

  log.Out("This is to Stdout")
  log.Err("This is to Stderr")
}
