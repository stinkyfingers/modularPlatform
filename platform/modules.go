package server

import (
	"github.com/stinkyfingers/modularPlatform/config"
)

// GetModulesFromConfig returns the "modules" field from the config
func GetModulesFromConfig() ([]Module, error) {
	configFields := make(map[string][]Module)
	err := config.GetConfig(configFields)
	return configFields["modules"], err
}

func RunModules(modules []Module) error {
	for _, module := range modules {
		err := module.run()
		if err != nil {
			return err
		}
	}
	return nil
}
