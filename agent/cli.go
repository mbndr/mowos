package agent

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

        // read config
        configPath := c.String("config")
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

        // start tcp server
        l, err := net.Listen("tcp", config.Agent.ListenIP+":"+config.Agent.ListenPort)
        if err != nil {
            return err
        }
        defer l.Close()

        mowos.Log.Infof("listening on tcp %s:%s", config.Agent.ListenIP, config.Agent.ListenPort)

        // save memory
        config = nil

        // accept connections
        for {
            conn, err := l.Accept()
            if err != nil {
                mowos.Log.Error("error accepting: ", err.Error())
            }
            go handleRequest(conn, disp)
        }

        return nil
    }
    return app
}

// handle an incoming request from a mowos-monitor
func handleRequest(conn net.Conn, disp *dispatcher) {
    defer func() {
        mowos.Log.Debug("closing connection to " + conn.RemoteAddr().String())
        conn.Close()
    }()
    mowos.Log.Debug("new connection to " + conn.RemoteAddr().String())

    // read message
    msg, err := mowos.ReadBytes(bufio.NewReader(conn))
    if err != nil {
        mowos.Log.Error(errors.Wrap(err, "error reading"))
    }

    mowos.Log.Debugf("%#v", string(msg))

    // send reply
    if string(msg) == "REQUEST" {
        itemResp, err := disp.getItemResponsesBytes()
        if err != nil {
            mowos.Log.Error(errors.Wrap(err, "error getting response"))
        }

        mowos.SendBytes(conn, itemResp)
    }
}
