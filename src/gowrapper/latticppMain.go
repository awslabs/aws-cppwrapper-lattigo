package main

import "C"

import (
	_ "lattigo-cpp/ckks"
	_ "lattigo-cpp/marshal"
	_ "lattigo-cpp/utils"
	_ "lattigo-cpp/ring"
)

// required by cgo
func main() {}
