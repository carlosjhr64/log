package log // import "github.com/carlosjhr64/log"

Very simple logger with mutex.
const VERSION = "0.1.0"
var ErrWriter = os.Stderr
var OutWriter = os.Stdout
func Err(a ...interface{})
func Out(a ...interface{})
func New(format string) *Format
type Format string
