# `ExecReplace()` for Go

🏃 Use [`syscall.Exec`](https://pkg.go.dev/syscall#Exec) on Unix or an [`execve(2)`-like shim](https://github.com/jcbhmr/go-execreplace/blob/main/execreplace_windows.go) on Windows

## Installation

```sh
go get github.com/jcbhmr/go-execreplace
```

## Usage

```go
package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/jcbhmr/go-execreplace"
)

func main() {
	fmt.Println("Hello from before ExecReplace!")
	goPath, err := exec.LookPath("go")
	if err != nil {
		log.Fatalf("LookPath %q failed: %v", "go", err)
	}
	execreplace.ExecReplace(goPath, []string{"go", "version"}, os.Environ())
}
```

```sh
go run ./main.go
```

```
Hello from before ExecReplace!
go version go1.25.4 windows/amd64
```