package main

import "github.com/carlosjhr64/log"

func main() {
  log.Out("This is to Stdout", "OK")
  log.Err("This is to Stderr", "OK")

  log.Format = "%s, %s.\n"

  log.Out("This is to Stdout", "OK")
  log.Err("This is to Stderr", "OK")
}
