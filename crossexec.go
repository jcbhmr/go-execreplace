package crossexec

import (
	"os"
)

func CrossExecProcess(name string, argv []string, attr *os.ProcAttr) error {
	return crossExecProcess(name, argv, attr)
}
