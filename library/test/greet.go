package main

import (
	"github.com/andrescosta/wrkl/sdk"
)

// main is required for TinyGo to compile to Wasm.
func main() {}

//export init
func _init() {
	sdk.OnEvent = test
}

func test(data string) string {
	sdk.Log(sdk.DebugLevel, data)
	sdk.Log(sdk.InfoLevel, data)
	return 1, "ok"
}
