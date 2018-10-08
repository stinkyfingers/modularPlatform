package config

import (
	"os"

	"github.com/go-yaml/yaml"
)

var (
	defaultConfigLocation = "/etc/netanal.conf"
	CONFIG_LOCATION       = "CONFIG_LOCATION"
)

// GetConfigLocation returns envvar CONFIG_LOCATION if set, otherwise the defaultConfigLocation
func getConfigFileLocation() string {
	if os.Getenv("CONFIG_LOCATION") == "" {
		return defaultConfigLocation
	}
	return os.Getenv("CONFIG_LOCATION")
}

// GetConfig assigns values from the config to v
func GetConfig(v interface{}) error {
	f, err := os.Open(getConfigFileLocation())
	if err != nil {
		return err
	}

	decoder := yaml.NewDecoder(f)
	return decoder.Decode(v)
}
