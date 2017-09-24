package agent

import (
	"github.com/urfave/cli"

	"github.com/mbndr/mowos"
)

// Version of the agent
var Version string = "0.0.0-dev"

func NewCliApp() *cli.App {
	app := cli.NewApp()
	app.Name = "mowos-agent"
	app.Usage = "sends data to mowos-monitor"
	app.Version = Version
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Value: mowos.DefaultConfigPath() + "/mowos-agent.yml",
			Usage: "config file to load",
		},
		cli.BoolFlag{
			Name:  "verbose, vv",
			Usage: "show debug output",
		},
	}
	app.Commands = []cli.Command{
		mowos.CreateDefaultConfigCommand("mowos-agent.yml"),
	}
	app.Action = func(c *cli.Context) error {
		mowos.Log.SetPrefix("bootstrap")
		err := bootstrapAgent(c)
		if err != nil {
			return err
		}

		mowos.Log.SetPrefix("agent")
		err = runAgent(c)
		if err != nil {
			return err
		}

		return nil
	}
	return app
}
