package main

// #include <stdlib.h>
import "C"

import (
	"fmt"
	"github.com/andrescosta/wrkl/sdk"
)

// main is required for TinyGo to compile to Wasm.
func main() {}

// greet prints a greeting to the console.
func greet(name string) {
	sdk.log(fmt.Sprint("wasm >> ", greeting(name)))
}


// greeting gets a greeting for the name.
func greeting(name string) string {
	return fmt.Sprint("Hello, ", name, "!")
}

// _greet is a WebAssembly export that accepts a string pointer (linear memory
// offset) and calls greet.
//
//export greet
func _greet(ptr, size uint32) {
	name := sdk.ptrToString(ptr, size)
	greet(name)
}

// _greeting is a WebAssembly export that accepts a string pointer (linear memory
// offset) and returns a pointer/size pair packed into a uint64.
//
// Note: This uses a uint64 instead of two result values for compatibility with
// WebAssembly 1.0.
//
//export greeting
func _greeting(ptr, size uint32) (ptrSize uint64) {
	name := sdk.ptrToString(ptr, size)
	g := greeting(name)
	ptr, size = sdk.stringToLeakedPtr(g)
	return (uint64(ptr) << uint64(32)) | uint64(size)
}

