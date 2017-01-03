package discorddgo

import (
	"io"
	"os"

	"github.com/Sirupsen/logrus"
)

var log = logrus.New()
var out = io.Writer(os.Stdout)

func init() {
	log.Level = logrus.WarnLevel
	log.Out = nil
}

// EnableLogging enables default Level logging.
// The Logging Level is not reset
func EnableLogging() {
	log.Out = out
}

// ResetLogLevel returns the logger to Warn messages or higher.
func ResetLogLevel() {
	log.Level = logrus.WarnLevel
}

// EnableDebug will call EnableLogging and then set the highest
// logging level. Use this to see what might go wrong under the
// hood.
func EnableDebug() {
	EnableLogging()
	log.Level = logrus.DebugLevel
}

// SetLogOutput will set the log output to the given
// io.Writer. It will not enable logging and if it was
// already enabled, EnableLogging **MUST** be called again
// otherwise it will continue to output on the previous io.Writer.
func SetLogOutput(output io.Writer) {
	out = output
}
