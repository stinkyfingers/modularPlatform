package config

import (
	"os"
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
