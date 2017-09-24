package mowos

import (
	"github.com/mbndr/logo"
	"os"
)

// Log is the log object used
var Log *logo.Logger

// InitLog initiates the logger
func InitLog() {
	cliRec := logo.NewReceiver(os.Stderr, "")
	cliRec.Color = true
	cliRec.Level = logo.INFO

	Log = logo.NewLogger(cliRec)
}

// SetLogLevel sets the verbose debug level
func SetLogLevel(verbose bool) {
	if verbose {
		Log.SetLevel(logo.DEBUG)
	}
}
