package mowos

import (
	"io/ioutil"

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
