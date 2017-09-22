package monitor

var config *monitorConfig

// this is unmarshalled from config file
type monitorConfig struct {
	Hosts []struct {
		Name        string `yaml:"name"`
		Description string `yaml:"description"`
		IP          string `yaml:"ip"`
		Port        string `yaml:"port"`
	} `yaml:"hosts"`
}
