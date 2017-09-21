package monitor

import (
    "net"
    "bufio"
    "github.com/pkg/errors"
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
    app.Flags = []cli.Flag {
        cli.StringFlag{
            Name: "config, c",
            Value: "config/laptop-monitor.yml",
            Usage: "config file to load",
        },
    }
    app.Action = func(c *cli.Context) error {
        mowos.Log.Infof("starting %s version %s", c.App.Name, c.App.Version)

        // read config
        configPath := c.String("config")
        mowos.Log.Debug("using ", configPath)
        err := loadConfigFile(configPath)
        if err != nil {
            return errors.Wrap(err, "read config file")
        }

        // start web server

        // DEBUG: send one request
        addr := config.Hosts[0].IP+":"+config.Hosts[0].Port
        conn, err := net.Dial("tcp", addr)
        if err != nil {
            return err
        }

        conn.Write([]byte("REQUEST\r\n\r\n"))

        reply, err := mowos.ReadBytes(bufio.NewReader(conn))
        if err != nil {
            mowos.Log.Error(errors.Wrap(err, "error reading"))
        }

        mowos.Log.Debugf("%#v", string(reply))

        return nil
    }
    return app
}
