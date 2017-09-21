package mowos

import (
	"github.com/mbndr/logo"
	"os"
)

var Log *logo.Logger

// InitLog initiates the logger
func InitLog(prefix string, level int) {
	cliRec := logo.NewReceiver(os.Stderr, prefix)
	cliRec.Color = true
	cliRec.Level = logo.Itol(level)

	Log = logo.NewLogger(cliRec)
}
