package mowos

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
)

// LoadConfigFile loads unmarshalls a file into an object (pointer)
func LoadConfigFile(path string, obj interface{}) error {
	fileContent, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(fileContent, obj)
	if err != nil {
		return err
	}

	return nil
}

// DefaultConfigPath returns the path of the default config directory
func DefaultConfigPath() string {
	path := os.Getenv("XDG_CONFIG_HOME")

	if path == "" {
		path = os.Getenv("HOME") + "/.config"
	}

	return filepath.Join(path, "mowos")
}

// CreateDefaultConfig creates a new configuration if it doesn't exist.
// file is the subfile (f.e. "mowos-agent.yml")
func CreateDefaultConfig(file string) error {
	dir := DefaultConfigPath()
	fileFull := filepath.Join(dir, file)

	// config file already exists
	if _, err := os.Stat(fileFull); !os.IsNotExist(err) {
		return errors.New("config file already exists")
	}

	// create the directory
	Log.Info("create " + dir)
	err := os.MkdirAll(dir, 0770)
	if err != nil {
		return err
	}
	// write the config file
	data, err := Asset("config/" + file)
	if err != nil {
		return err
	}
	Log.Info("write " + fileFull)
	err = ioutil.WriteFile(fileFull, data, 0660)
	return err
}

// CreateDefaultConfigCommand returns a command for creating the default configuration.
// just for decreasing redundancy a bit
func CreateDefaultConfigCommand(file string) cli.Command {
	return cli.Command{
		Name:  "init",
		Usage: "create the default configuration",
		Action: func(c *cli.Context) error {
			return CreateDefaultConfig(file)
		},
	}
}
