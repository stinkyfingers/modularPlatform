package main

import (
	"log"

	"github.com/stinkyfingers/modularPlatform/server"
)

func main() {
	modules, err := server.GetModulesFromConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = server.RunModules(modules)
	if err != nil {
		log.Fatal(err)
	}
}
