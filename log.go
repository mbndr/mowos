package mowos

import (
    "os"
    "github.com/mbndr/logo"
)

var Log *logo.Logger

// InitLog initiates the logger
func InitLog(prefix string, level int) {
    cliRec := logo.NewReceiver(os.Stderr, prefix)
    cliRec.Color = true
    cliRec.Level = logo.Itol(level)

    Log = logo.NewLogger(cliRec)
}
