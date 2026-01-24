# Cross-platform `Exec` for Go

⚡️ \*nix `CmdExt.Exec` but works on Windows

<table align=center><td>

<div>💻 <code>syscall.Exec</code>-like</div>

```go
goPath, _ := exec.LookPath("go")
_ = crossexec.CrossExec(goPath, []string{"go", "version"}, nil)
```

<div>🚀 <code>exec.Cmd</code>-based</div>

```go
cmd := exec.Command("go", "version")
// Cross-platform
_ = (*crossexec.CmdExt)(cmd).CrossExec()
```

</table>

🐧 Uses [`syscall.Exec`](https://pkg.go.dev/syscall#Exec) on Unix \
🟦 Emulates process replacement on Windows

## Installation

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=Go&logoColor=FFFFFF)

```sh
go get github.com/jcbhmr/go-crossexec
```

## Usage

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=Go&logoColor=FFFFFF)

[📚 See the docs](https://pkg.go.dev/github.com/jcbhmr/go-crossexec)

## Development

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=Go&logoColor=FFFFFF)
![Windows](https://img.shields.io/badge/Windows-00BFFF?style=for-the-badge)

Inspired by [`exec_replace`](https://docs.rs/cargo-util/latest/cargo_util/struct.ProcessBuilder.html#method.exec_replace) from [`cargo-util`](https://docs.rs/cargo-util/latest/cargo_util/).
