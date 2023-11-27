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

func test(data string) (uint64, string) {
	sdk.Log(sdk.DebugLevel, data)
	sdk.Log(sdk.InfoLevel, data)
	return sdk.NoError, "ok"
}
