package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	log "github.com/sirupsen/logrus"
)

func main() {
	cmd := "ls"
	result, err := ExecCmd(cmd)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Printf("Result: %v\n", string(result))
}

// ExecCmd execute command on host and return the result to client.
func ExecCmd(command string) ([]byte, error) {
	cmd := exec.Command("/bin/sh", "-c", command)

	stdout, err := cmd.StdoutPipe() // Obtain cmd output pipe
	if err != nil {
		log.Errorf("Error: obtain stdout pipe")
		return nil, err
	}

	err = cmd.Start() // execute command
	if err != nil {
		log.Errorf("Error: execute command")
		return nil, err
	}

	bytes, err := ioutil.ReadAll(stdout) // read result
	if err != nil {
		log.Errorf("Error: read result")
		return nil, err
	}

	err = cmd.Wait()
	if err != nil {
		log.Errorf("Error: wait to exit")
		return nil, err
	}

	return bytes, nil
}
