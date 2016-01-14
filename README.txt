package log // import "github.com/carlosjhr64/log"

var ErrWriter = os.Stderr
var Format = ""
var OutWriter = os.Stdout
func Err(a ...interface{})
func Out(a ...interface{})
