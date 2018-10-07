package server

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

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
