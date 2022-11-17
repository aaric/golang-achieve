# golang-achieve

[![License](https://img.shields.io/badge/License-MIT-green.svg?style=flat&logo=github)](https://www.mit-license.org)
[![Go](https://img.shields.io/badge/Go-1.18.8-success.svg?style=flat&logo=go)](https://go.dev)
[![Release](https://img.shields.io/badge/Release-0.2.0-blue.svg)](https://github.com/aaric/golang-achieve/releases)

> GO Lang Learning.

## 1 fyne v2

|No.|Env Key|Pkg Path|Remark|
|:---:|:---:|-----|-----|
|1|`GOPATH`|`~/go`|*default*|
|2|`GOPATH`|`/path/to/go`| |
|3|`PATH`|`$PATH:$GOPATH/bin`|*install binary*|

&emsp;&emsp;*Download to `$GOPATH/pkg/mod/fyne.io`.*

```bash
# init
#go mod init

# download
# go clean -i -n fyne.io/fyne/v2
go get -u fyne.io/fyne/v2

# tidy
go mod tidy

# run
go run fyne.go

# package
go install fyne.io/fyne/v2/cmd/fyne@latest
#fyne package --name fyne --icon Icon.png --appVersion 0.1.0 fyne.go
fyne package --name fyne --appVersion 0.1.0 fyne.go
```
