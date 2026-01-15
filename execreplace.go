package execreplace

// Unix has [syscall.Exec], Windows does not. There is no way on Windows to
// assimilate another process into the current one. ExecReplace is a
// cross-platform abstraction for this functionality.
//
//  - path: should be a path to an executable file. No PATH lookup is performed.
//  - argv: will default to os.Args if nil.
//  - envv: will default to os.Environ() if nil.
//
// # Unix
//
// On Unix, ExecReplace quickly calls [syscall.Exec]. This replaces the current
// process with the new process. Its PID remains the same.
//
// # Windows
//
// On Windows, ExecReplace uses [os/exec.Cmd] and normal subprocess creation. It
// then calls [os.Exit] with the exit code of the subprocess. The child will
// have a different PID from the parent.
//
// This behavior is different from [CRT _execve]. CRT _execve
// starts a new process as a child of the old process' parent, and then terminates
// the old process. The new process has a different PID.
//
// [CRT _execve]: https://learn.microsoft.com/en-us/cpp/c-runtime-library/reference/execve-wexecve?view=msvc-170
func ExecReplace(path string, argv []string, envv []string) error {
	return execReplace(path, argv, envv)
}
