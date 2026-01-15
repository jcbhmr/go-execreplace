//go:build unix

package execreplace

import (
	"os"
	"syscall"
)

func execReplace(path string, argv []string, envv []string) error {
	if argv == nil {
		argv = os.Args
	}
	if envv == nil {
		envv = os.Environ()
	}
	
	return syscall.Exec(path, argv, envv)
}
