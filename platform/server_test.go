package platform

import (
	"go/build"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestStart(t *testing.T) {
	p := NewServer("10000")
	err := p.Start()
	if err != nil {
		t.Error(err)
	}

	cmd := exec.Command("node", filepath.Join(build.Default.GOPATH, "src", "github.com", "stinkyfingers", "modularPlatform", "modules", "js_example", "client.js"))
	out, err := cmd.Output()
	if err != nil {
		t.Error(err)
	}
	t.Log(string(out))

	err = p.Stop()
	if err != nil {
		t.Error(err)
	}

	return
}
