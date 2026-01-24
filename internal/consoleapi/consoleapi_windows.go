package consoleapi

import "golang.org/x/sys/windows"

var kernel32 = windows.NewLazySystemDLL("kernel32.dll")

var setConsoleCtrlHandler = kernel32.NewProc("SetConsoleCtrlHandler")

const windowsTRUE = 1
const windowsFALSE = 0

type PHANDLER_ROUTINE uintptr

func SetConsoleCtrlHandler(handler PHANDLER_ROUTINE, add bool) error {
	add2 := func() uintptr {
		if add {
			return windowsTRUE
		} else {
			return windowsFALSE
		}
	}()
	_, _, err := setConsoleCtrlHandler.Call(uintptr(handler), add2)
	if err != nil && err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}