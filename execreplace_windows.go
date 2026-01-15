package execreplace

import (
	"os"
	"os/exec"

	"github.com/jcbhmr/go-execreplace/internal/errorsastype"
)

func execReplace(argv0 string, argv []string, envv []string) error {
	cmd := exec.Cmd{
		Path:   argv0,
		Args:   argv,
		Env:    envv,
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
	os.Stdin = nil
	os.Stdout = nil
	os.Stderr = nil
	err := cmd.Run()
	if err != nil {
		// TODO: Use new errors.AsType() when widely available.
		// https://antonz.org/accepted/errors-astype/
		if _, ok := errorsastype.AsType[*exec.ExitError](err); ok {
			// continue
		} else {
			return err
		}
	}
	os.Exit(cmd.ProcessState.ExitCode())
	panic("unreachable: os.Exit should never return")
}
