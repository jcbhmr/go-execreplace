package crossexec

import (
	"fmt"
	"os"
	"runtime"
	"sync"

	"github.com/jcbhmr/go-crossexec/internal/consoleapi"
	"golang.org/x/sys/windows"
)

const windowsTRUE = 1

// const windowsFALSE = 0

var getIgnoreCtrlHandler = sync.OnceValue(func() consoleapi.PHANDLER_ROUTINE {
	return consoleapi.PHANDLER_ROUTINE(windows.NewCallback(func(_ uint32) uintptr {
		return windowsTRUE
	}))
})

var forked sync.Mutex

func crossExecProcess(name string, argv []string, attr *os.ProcAttr) (err error) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	forked.Lock()
	defer forked.Unlock()

	err = consoleapi.SetConsoleCtrlHandler(getIgnoreCtrlHandler(), true)
	if err != nil {
		return err
	}

	process, err := os.StartProcess(name, argv, attr)
	if err != nil {
		return err
	}
	state, err := process.Wait()
	if err != nil {
		fmt.Fprintf(os.Stderr, "crossexec: %v\n", err)
		os.Exit(1)
	}
	os.Exit(state.ExitCode())
	panic("noreturn")
}
