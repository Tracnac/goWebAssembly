# Windows
# $Env:GOOS="js" ; $Env:GOARCH="wasm" ; go build -o .\wasm\module.wasm .\wasm\wasm.go
# $Env:GOOS="windows" ; $Env:GOARCH="amd64" ; go build

# Linux
GOOS="js" GOARCH="wasm" go build -o ./wasm/module.wasm ./wasm/wasm.go
GOOS="linux" GOARCH="amd64"  go build
