package minio

import (
	"log"
	"os/exec"
)

// ExecCmd - Execute Command Shell
func ExecCmd(command string, arguments []string) (string, error) {
	cmd := exec.Command(command, arguments...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("cmd.Run() failed with %s\n", err)
		return "", err
	}
	return string(out), nil
}
