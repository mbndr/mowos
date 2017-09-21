package main

import (
	"os"

	"github.com/mbndr/mowos"
	"github.com/mbndr/mowos/agent"
)

func main() {
	mowos.InitLog("agent ", 0)

	cli := agent.NewCliApp()
	err := cli.Run(os.Args)
	if err != nil {
		mowos.Log.Fatal(err)
	}
}
