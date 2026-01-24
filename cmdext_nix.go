//go:build unix || plan9

package crossexec

import (
	jcbhmrexec "github.com/jcbhmr/go-exec"
)

func (c *CmdExt) crossExec() error {
	return (*jcbhmrexec.CmdExt)(c).Exec()
}
