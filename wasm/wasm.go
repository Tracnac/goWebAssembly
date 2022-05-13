package main

import (
	"fmt"
	"strconv"
	"syscall/js"
)

func Wasm() js.Func {
	wasm := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 1 {
			return "!Number"
		}
		n, _ := strconv.Atoi(args[0].String())
		return 42 * n
	})
	return wasm
}

func main() {
	fmt.Println("Wasm loaded")
	js.Global().Set("Wasm", Wasm())
	<-make(chan bool)
}
