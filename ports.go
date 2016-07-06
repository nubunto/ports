package ports

import (
	"bytes"
	"os/exec"
)

type Program interface {
	Run() ([]byte, error)
}

type baseProgram struct {
	name   string
	params []string
	stdout bytes.Buffer
}

func (b baseProgram) Run() ([]byte, error) {
	cmd := exec.Command(b.name, b.params...)
	cmd.Stdout = &b.stdout
	err := cmd.Start()
	if err != nil {
		return nil, err
	}
	err = cmd.Wait()
	if err != nil {
		return nil, err
	}
	return b.stdout.Bytes(), nil
}

func Launch(name string, params []string) ([]byte, error) {
	b := &baseProgram{name: name, params: params}
	res, err := b.Run()
	if err != nil {
		return nil, err
	}
	return res, nil
}
