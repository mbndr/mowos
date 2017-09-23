package agent

import (
	"bufio"
	"net"

	"github.com/mbndr/mowos"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
)

// holds and runs all items
var disp *dispatcher

// does bootstrapping stuff to reduce complexity in main function
func bootstrapAgent(c *cli.Context) error {
	mowos.SetLogLevel(c.Bool("verbose"))

	// read config
	configPath := c.String("config")
	mowos.Log.Debug("using ", configPath)
	err := mowos.LoadConfigFile(configPath, &config)
	if err != nil {
		return errors.Wrap(err, "read config file")
	}

	disp = &dispatcher{}

	// fill dispatcher
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

	return nil
}

// runs the agent
func runAgent(c *cli.Context) error {
	mowos.Log.SetPrefix("agent")
	mowos.Log.Infof("starting %s version %s", c.App.Name, c.App.Version)

	addr := config.Agent.ListenIP + ":" + config.Agent.ListenPort

	// start tcp server
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	defer l.Close()

	mowos.Log.Info("listening on tcp ", addr)

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
