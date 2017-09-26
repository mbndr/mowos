package monitor

import (
	"github.com/mbndr/mowos"
)

var config *monitorConfig

// this is unmarshalled from config file
type monitorConfig struct {
	Monitor struct {
		ListenIP   string `yaml:"listen-ip"`
		ListenPort string `yaml:"listen-port"`
	} `yaml:"monitor"`
	Hosts []struct {
		Name        string    `yaml:"name"`
		Description string    `yaml:"description"`
		IP          string    `yaml:"ip"`
		Port        string    `yaml:"port"`
		PSK         mowos.PSK `yaml:"psk"`
	} `yaml:"hosts"`
}
