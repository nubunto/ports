package ports

import (
	"bytes"
	"log"
	"os/exec"
)

func Launch(program string, outputChan chan []byte) error {
	var stdout bytes.Buffer
	cmd := exec.Cmd{
		Path:   program,
		Stdout: &stdout,
		Stderr: &stdout,
	}
	err := cmd.Start()
	if err != nil {
		return err
	}
	go func() {
		err := cmd.Wait()
		if err != nil {
			// what to do if command fails? no way to access the top goroutine, so... ?
			// decide what to do with it later.
			// maybe decouple Launch from doing everything...
			log.Println(err)
		}
		outputChan <- stdout.Bytes()
		close(outputChan)
	}()
	return nil
}
