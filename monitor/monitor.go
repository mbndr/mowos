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

	return nil
}

// runs the monitor
func runMonitor(c *cli.Context) error {
	mowos.Log.Infof("starting %s version %s", c.App.Name, c.App.Version)

	// init web server
	//return startWebServer()

	// DEBUG: send one request
	host := config.Hosts[0]

	addr := host.IP + ":" + host.Port
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return err
	}

	// CRYPTO
	mowos.UsedCryptor = mowos.NewPSKCryptor(
		[]byte(host.PSK.Key),
		[]byte(host.PSK.Identity),
	)

	// send to agent
	mowos.SendBytes(conn, []byte("REQUEST"))

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
