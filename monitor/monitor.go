package monitor

import (
	"bufio"
	"encoding/json"
	"net"

	"github.com/pkg/errors"
	"github.com/urfave/cli"

	"github.com/mbndr/mowos"
)

// does bootstrapping stuff to reduce complexity in main function
func bootstrapMonitor(c *cli.Context) error {
	mowos.SetLogLevel(c.Bool("verbose"))

	// read config
	configPath := c.String("config")
	mowos.Log.Debug("using ", configPath)
	err := mowos.LoadConfigFile(configPath, &config)
	if err != nil {
		return errors.Wrap(err, "read config file")
	}

	// init web server

	return nil
}

// runs the monitor
func runMonitor(c *cli.Context) error {
	mowos.Log.Infof("starting %s version %s", c.App.Name, c.App.Version)

	// DEBUG: send one request
	addr := config.Hosts[0].IP + ":" + config.Hosts[0].Port
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return err
	}

	// send to agent
	conn.Write([]byte("REQUEST\r\n\r\n"))

	// get raw response
	raw, err := mowos.ReadBytes(bufio.NewReader(conn))
	if err != nil {
		mowos.Log.Error(errors.Wrap(err, "reading"))
	}

	// convert to object
	reply := make(mowos.AgentResponse)
	err = json.Unmarshal(raw, &reply)
	if err != nil {
		mowos.Log.Error(errors.Wrap(err, "converting"))
	}

	mowos.Log.Warnf("%#v", reply)

	return nil
}
