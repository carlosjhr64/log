package main

import "github.com/carlosjhr64/log"

func main() {
  log.Out("Hello, how are you?")
  log.Err("Ok, and you?")
  a := log.New("I'm %s, %s.\n")
  b := log.New("%s are %s!\n")
  a.Out("fine", "thank you")
  b.Err("You", "welcomed")
  b.Err("We", "good")
}
