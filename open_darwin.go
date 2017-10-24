package system

import "os/exec"

func open(name string) *exec.Cmd {
	return exec.Command("open", name)
}
