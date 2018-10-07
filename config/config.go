package config

import (
	"os"

	"github.com/go-yaml/yaml"
)

// GetConfig assigns values from the config to v
func GetConfig(v interface{}) error {
	f, err := os.Open(getConfigFileLocation())
	if err != nil {
		return err
	}

	decoder := yaml.NewDecoder(f)
	return decoder.Decode(v)
}
