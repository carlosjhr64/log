package log // import "github.com/carlosjhr64/log"

Very simple logger with mutex.

var ErrFormat = ""
var ErrWriter = os.Stderr
var OutFormat = ""
var OutWriter = os.Stdout

func Err(a ...interface{})
func Out(a ...interface{})
