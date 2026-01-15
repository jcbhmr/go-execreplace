package execreplace

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"sync"

	"github.com/jcbhmr/go-execreplace/internal/errorsastype"
	"golang.org/x/sys/windows"
)

var kernel32 = windows.NewLazySystemDLL("kernel32.dll")
var setConsoleCtrlHandler = kernel32.NewProc("SetConsoleCtrlHandler")
const windowsTRUE = 1
const windowsFALSE = 0

var executing sync.Mutex

var getCtrlcHandler = sync.OnceValue(func() uintptr {
	return windows.NewCallback(func(ctrlType uint32) uintptr {
		return windowsTRUE
	})
})

func execReplace(path string, argv []string, envv []string) error {
	executing.Lock()
	defer executing.Unlock()

	if runtime.NumGoroutine() > 1 {
		// TODO: Do something?
	}

	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	defer runtime.GOMAXPROCS(runtime.GOMAXPROCS(0))
	runtime.GOMAXPROCS(1)

	if argv == nil {
		argv = os.Args
	}

	cmd := exec.Cmd{
		Path:   path,
		Args:   argv,
		Env:    envv,
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}

	_, _, err := setConsoleCtrlHandler.Call(getCtrlcHandler(), windowsTRUE)
	if err != nil {
		if err == windows.ERROR_SUCCESS {
			// continue
		} else {
			return fmt.Errorf("could not set Ctrl+C handler: %w", err)
		}
	}
	defer func() {
		_, _, err := setConsoleCtrlHandler.Call(getCtrlcHandler(), windowsFALSE)
		if err != nil {
			if err == windows.ERROR_SUCCESS {
				// continue
			} else {
				err = fmt.Errorf("could not unset Ctrl+C handler: %w", err)
				panic(err)
			}
		}
	}()

	err = cmd.Run()
	if err != nil {
		if _, ok := errorsastype.AsType[*exec.ExitError](err); ok {
			// continue
		} else {
			return err
		}
	}
	os.Exit(cmd.ProcessState.ExitCode())
	panic("unreachable: os.Exit should never return")
}
