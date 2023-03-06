package internal

import "os/exec"

func ExecCommand(cmd string, args ...string) error {
	c := exec.Command(cmd, args...)
	return c.Run()
}
