package execreplace

func ExecReplace(argv0 string, argv []string, envv []string) error {
	return execReplace(argv0, argv, envv)
}
