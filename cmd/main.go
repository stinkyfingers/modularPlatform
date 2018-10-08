package main

import (
	"log"

	"github.com/stinkyfingers/modularPlatform/platform"
)

func main() {
	modules, err := platform.GetModulesFromConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = platform.RunModules(modules)
	if err != nil {
		log.Fatal(err)
	}
}
