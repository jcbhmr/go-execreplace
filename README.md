# Cross-platform `Exec` for Go

⚡️ \*nix [`CmdExt.Exec`](https://pkg.go.dev/github.com/jcbhmr/go-exec#CmdExt.Exec) but works on Windows

<table align=center><td>

```go
cmd := exec.Command("go", "version")
log.Fatal((*crossexec.CmdExt)(cmd).CrossExec())
```

</table>

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

Inspired by [`exec_replace`](https://docs.rs/cargo-util/latest/cargo_util/struct.ProcessBuilder.html#method.exec_replace) from [`cargo-util`](https://docs.rs/cargo-util/latest/cargo_util/).
