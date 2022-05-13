package main

import (
	"fmt"
	"strconv"
	"syscall/js"
)

func Wasm() js.Func {
	wasm := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		n, err := strconv.Atoi(js.Global().Get("document").Call("getElementById", "input").Get("value").String()) // Dom manipulation to get value
		if err != nil {
			js.Global().Get("document").Call("getElementById", "output").Set("value", "! Number")
		}
		js.Global().Get("document").Call("getElementById", "output").Set("value", 42*n) // Dom manipulation to set value
		return nil
	})
	return wasm
}
func WasmJS() js.Func {
	wasm := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 1 {
			return "! Number"
		}
		n, _ := strconv.Atoi(args[0].String())
		return 42 * n
	})
	return wasm
}

func main() {
	fmt.Println("Wasm loaded")
	js.Global().Set("Wasm", Wasm())     // Export function to JS
	js.Global().Set("WasmJS", WasmJS()) // Export function to JS

	js.Global().Get("document").Call("getElementById", "input").Set("value", 1) // Dom manipulation to set value

	var wasm js.Func = Wasm()
	js.Global().Get("document").Call("getElementById", "submit").Call("addEventListener", "click", wasm)

	<-make(chan bool)
}
