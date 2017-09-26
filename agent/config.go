package agent

import (
	"github.com/mbndr/mowos"
)

var config *agentConfig

// this is unmarshalled from config file
type agentConfig struct {
	Agent struct {
		ListenIP   string `yaml:"listen-ip"`
		ListenPort string `yaml:"listen-port"`
	} `yaml:"agent"`
	PSK mowos.PSK `yaml:"psk"`
	// items is a list of maps
	Items []map[string]interface{} `yaml:"items"`
}
