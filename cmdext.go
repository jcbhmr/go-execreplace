package crossexec

import (
	"os/exec"
)

type CmdExt exec.Cmd

func (c *CmdExt) CrossExec() error {
	return c.crossExec()
}
