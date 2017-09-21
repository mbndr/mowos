package main

import (
	"os"

	"github.com/mbndr/mowos"
	"github.com/mbndr/mowos/monitor"
)

func main() {
	mowos.InitLog("monitor ", 0)

	cli := monitor.NewCliApp()
	err := cli.Run(os.Args)
	if err != nil {
		mowos.Log.Fatal(err)
	}
}
