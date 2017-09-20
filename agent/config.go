package agent

import (
    "io/ioutil"

    "gopkg.in/yaml.v2"
)

var config *agentConfig

// this is unmarshalled from config file
type agentConfig struct {
    Agent struct {
        ListenIP string `yaml:"listen-ip"`
        ListenPort string `yaml:"listen-port"`
    } `yaml:"agent"`
    // items is a list of maps
    Items []map[string]interface{} `yaml:"items"`
}

func loadConfigFile(path string) error {
    fileContent, err := ioutil.ReadFile(path)
    if err != nil {
        return err
    }

    err = yaml.Unmarshal(fileContent, &config)
    if err != nil {
        return err
    }

    return nil
}
