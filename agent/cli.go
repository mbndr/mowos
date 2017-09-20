package agent

import (
    "github.com/pkg/errors"

    "github.com/urfave/cli"
    "github.com/mbndr/mowos"
)

// TODO compile in
var Version string = "0.0.1-dev"

func NewCliApp() *cli.App {
    app := cli.NewApp()
    app.Name = "mowos-agent"
    app.Usage = "sends data to mowos-monitor"
    app.Version = Version
    app.Flags = []cli.Flag {
        cli.StringFlag{
            Name: "config, c",
            Value: "config/linus-agent.yml",
            Usage: "config file to load",
        },
    }
    app.Action = func(c *cli.Context) error {
        mowos.Log.Infof("starting %s version %s", c.App.Name, c.App.Version)

        // only for dev purpose
        //configPath := "config/linus-agent.yml"
        configPath := c.String("config")

        // read config
        mowos.Log.Debug("using ", configPath)
        err := loadConfigFile(configPath)
        if err != nil {
            return errors.Wrap(err, "read config file")
        }

        // prepare dispatcher
        disp := &dispatcher{}

        // for each item
        for _, i := range config.Items {
            // get type and create item
            item := getItem(i)
            if item == nil {
                mowos.Log.Warn("unknown item type: ", i["type"])
                continue
            }
            disp.items = append(disp.items, item)
        }

        disp.logItems()

        // TO BE CALLED ON SERVER REQUEST
        ret := disp.getValues()
        // prettier print
        for k, s := range ret {
            mowos.Log.Debugf("%s: %s (%d)", k, s.Value, s.Status)
        }
        //mowos.Log.Debugf("%#v", ret)

        // start tcp server

        return nil
    }
    return app
}
