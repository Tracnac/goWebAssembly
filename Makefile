BINARY_NAME=sandbox
WASM_DIR=wasm
 
all: build

init:
		go get -u
		go mod tidy
		cp `go env GOROOT`/misc/wasm/wasm_exec.js ${WASM_DIR}/wasm_exec.js
		curl https://github.com/picocss/pico/blob/master/css/pico.min.css -o wasm/pico.min.css
		curl https://github.com/picocss/pico/blob/master/css/pico.min.css.map -o wasm/pico.min.css.map
build:
		GOOS="js" GOARCH="wasm" go build -o ${WASM_DIR}/module.wasm ${WASM_DIR}/wasm.go
		go build -o ${BINARY_NAME} ${BINARY_NAME}.go
 
test:
		go test -v ${BINARY_NAME}.go
 
run: build
		go build -o ${BINARY_NAME} main.go
		./${BINARY_NAME}
 
clean:
		go clean
		test -f '${WASM_DIR}/module.wasm' && rm -rf '${WASM_DIR}/module.wasm' || echo "done."
