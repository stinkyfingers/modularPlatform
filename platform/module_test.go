package platform

import (
	"go/build"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestGetModulesFromConfig(t *testing.T) {
	config := `
modules:
  - name: go_example
    runcommand: go run
    port: 9999
    location: /Users/johnshenk/go/src/github.com/stinkyfingers/modularPlatform/modules/go_example/cmd/main.go
  - name: js_example
    runCommand: node
    port: 100000
    location: /Users/johnshenk/go/src/github.com/stinkyfingers/modularPlatform/modules/js_example/client.js
`
	currentConfigLocation := os.Getenv("CONFIG_LOCATION")
	os.Setenv("CONFIG_LOCATION", "test.conf")
	defer os.Setenv("CONFIG_LOCATION", currentConfigLocation)
	err := ioutil.WriteFile("test.conf", []byte(config), os.ModePerm)
	if err != nil {
		t.Error(err)
	}
	defer os.Remove("test.conf")

	modules, err := GetModulesFromConfig()
	if err != nil {
		t.Error(err)
	}
	if len(modules) < 2 {
		t.Errorf("expected 2 modules, got %d", len(modules))
	}
	if modules[0].Name != "go_example" {
		t.Errorf("expected module runCommand to be 'go run', got '%s'", modules[0].Name)
	}
	if modules[0].RunCommand != "go run" {
		t.Errorf("expected module runCommand to be 'go run', got '%s'", modules[0].RunCommand)
	}
	if modules[0].Port != "9999" {
		t.Errorf("expected module runCommand to be '9999', got '%s'", modules[0].Port)
	}
	if modules[0].Location != "/Users/johnshenk/go/src/github.com/stinkyfingers/modularPlatform/modules/go_example/cmd/main.go" {
		t.Errorf("expected module runCommand to be '/Users/johnshenk/go/src/github.com/stinkyfingers/modularPlatform/modules/go_example/cmd/main.go', got '%s'", modules[0].Location)
	}
}

func TestRunJSModule(t *testing.T) {
	m := &Module{
		Name:       "test_js_module",
		Location:   filepath.Join(build.Default.GOPATH, "src", "github.com", "stinkyfingers", "modularPlatform", "modules", "js_example", "client.js"),
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
		Location:   filepath.Join(build.Default.GOPATH, "src", "github.com", "stinkyfingers", "modularPlatform", "modules", "go_example", "cmd", "main.go"),
		RunCommand: "go run",
		Port:       "9999",
	}

	err := m.run()
	if err != nil {
		t.Error(err)
	}
}
