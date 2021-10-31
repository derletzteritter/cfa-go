package utils

import (
	"os/exec"
)

func CloneRepo(path string) error {
	c := exec.Command("git", "clone", "https://github.com/project-error/cfa-templates")
	c.Dir = path
	err := c.Run()

	return err
}

func HasCommand(cmd string) error {
	_, err := exec.LookPath(cmd)

	return err
}
