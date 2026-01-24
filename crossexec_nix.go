//go:build unix || plan9

package crossexec

import (
	"os"

	jcbhmrexec "github.com/jcbhmr/go-exec"
)

func crossExecProcess(name string, argv []string, attr *os.ProcAttr) error {
	return jcbhmrexec.ExecProcess(name, argv, attr)
}
