package platform

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

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

// run runs a grpc server for a single Module and executes the binary (or compiles/interprets+runs the code) for a Module
func (m *Module) run() error {
	p := NewServer(m.Port)
	err := p.Start()
	if err != nil {
		return err
	}

	runCommandArr := strings.Split(m.RunCommand, " ")
	commandArr := append(runCommandArr, m.Location)
	cmd := exec.Command(commandArr[0], commandArr[1:]...)

	// write cmd errs to stdOut
	err = pipeErr(cmd, os.Stdout)
	if err != nil {
		return err
	}

	out, err := cmd.Output()
	if err != nil {
		return err
	}
	fmt.Println("module output: ", string(out))

	err = p.Stop()
	if err != nil {
		return err
	}

	return nil
}

// pipeErr writes cmd stdErr to w
func pipeErr(cmd *exec.Cmd, w io.Writer) error {
	errPipe, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(errPipe)
	go func() {
		for scanner.Scan() {
			_, err = w.Write(scanner.Bytes())
			if err != nil {
				return
			}
		}
	}()
	return nil
}
