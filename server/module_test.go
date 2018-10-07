package server

import (
	"go/build"
	"path/filepath"
	"testing"
)

func TestRunJSModule(t *testing.T) {
	m := &Module{
		Name:       "test_js_module",
		Location:   filepath.Join(build.Default.GOPATH, "src", "github.com", "stinkyfingers", "netanal", "modules", "js_example", "client.js"),
		RunCommand: "node",
		Port:       "10000",
	}

	err := m.run()
	if err != nil {
		t.Error(err)
	}
}

func TestRunGoModule(t *testing.T) {
	m := &Module{
		Name:       "test_go_module",
		Location:   filepath.Join(build.Default.GOPATH, "src", "github.com", "stinkyfingers", "netanal", "modules", "go_example", "cmd", "main.go"),
		RunCommand: "go run",
		Port:       "9999",
	}

	err := m.run()
	if err != nil {
		t.Error(err)
	}
}
