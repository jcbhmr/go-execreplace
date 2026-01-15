package execreplace_test

import (
	"os"
	"os/exec"
	"testing"

	"github.com/jcbhmr/go-execreplace/internal/errorsastype"
)

func goRunEval(t *testing.T, code string) (output string) {
	err := os.MkdirAll(".test/go-run-eval", 0755)
	if err != nil {
		t.Fatalf("MkdirAll %q failed: %v", ".test/go-run-eval", err)
	}
	mainFile, err := os.CreateTemp(".test/go-run-eval", "main-*.go")
	if err != nil {
		t.Fatalf("CreateTemp %q in %q failed: %v", "main-*.go", ".test/go-run-eval", err)
	}
	_, err = mainFile.WriteString(code)
	if err != nil {
		t.Fatalf("WriteString %q to %q failed: %v", code, mainFile.Name(), err)
	}
	err = mainFile.Close()
	if err != nil {
		t.Fatalf("Close %q failed: %v", mainFile.Name(), err)
	}

	cmd := exec.Command("go", "run", mainFile.Name())
	t.Logf("Running %q", cmd)
	outputBytes, err := cmd.CombinedOutput()
	if err != nil {
		if _, ok := errorsastype.AsType[*exec.ExitError](err); ok {
			// continue
		} else {
			t.Fatalf("exec.Command failed: %v", err)
		}
	}
	return string(outputBytes)
}

func TestExecReplace(t *testing.T) {
	cmd := exec.Command("go", "version")
	t.Logf("Running %q", cmd)
	goVersionOutput, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("exec.Command failed: %v", err)
	}
	expectedOutput := "Hello from TestExecReplace\n" + string(goVersionOutput)

	actualOutput := goRunEval(t, `//go:build ignore

package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/jcbhmr/go-execreplace"
)

func main() {
	fmt.Println("Hello from TestExecReplace")
	goPath, err := exec.LookPath("go")
	if err != nil {
		log.Fatalf("LookPath %q failed: %v", "go", err)
	}
	execreplace.ExecReplace(goPath, []string{"go", "version"}, os.Environ())
}
`)

	if actualOutput != expectedOutput {
		t.Fatalf("expected %q, got %q", expectedOutput, actualOutput)
	}
}
