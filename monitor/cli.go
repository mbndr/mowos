package monitor

import (
	"github.com/urfave/cli"

	"github.com/mbndr/mowos"
)

// TODO compile in
var Version string = "0.0.1-dev"

func NewCliApp() *cli.App {
	app := cli.NewApp()
	app.Name = "mowos-monitor"
	app.Usage = "get data from mowos-agent and serve a web page"
	app.Version = Version
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Value: mowos.DefaultConfigPath() + "/mowos-monitor.yml",
			Usage: "config file to load",
		},
		cli.BoolFlag{
			Name:  "verbose, vv",
			Usage: "show debug output",
		},
	}
	app.Commands = []cli.Command{
		mowos.CreateDefaultConfigCommand("mowos-monitor.yml"),
	}
	app.Action = func(c *cli.Context) error {
		mowos.Log.SetPrefix("bootstrap")
		err := bootstrapMonitor(c)
		if err != nil {
			return err
		}

		mowos.Log.SetPrefix("monitor")
		err = runMonitor(c)
		if err != nil {
			return err
		}

		return nil
	}
	return app
}
