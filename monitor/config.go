package monitor

import (
    "io/ioutil"

    "gopkg.in/yaml.v2"
)

var config *monitorConfig

// this is unmarshalled from config file
type monitorConfig struct {
    Hosts []struct {
        Name string `yaml:"name"`
        Description string `yaml:"description"`
        IP string `yaml:"ip"`
        Port string `yaml:"port"`
    } `yaml:"hosts"`
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
