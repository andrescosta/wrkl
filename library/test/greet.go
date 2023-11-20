package main

import (
	"fmt"

	"github.com/andrescosta/wrkl/sdk"
)

// main is required for TinyGo to compile to Wasm.
func main() {}

//export init
func _init() {
	sdk.OnEvent = test
}

func test(data string) string {
	sdk.Log(sdk.Error, fmt.Sprint("error wasm >> ", data))
	sdk.Log(sdk.Info, fmt.Sprint("info wasm >> ", data))
	return "1"
}
