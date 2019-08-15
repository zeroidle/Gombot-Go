package main

import (
	"fmt"
	"runtime"
)

var DEBUG = false

func main() {
	setup()
}

func setup() {
	if runtime.GOOS == "darwin" {
		DEBUG = true
		fmt.Println(runtime.GOOS)
		fmt.Println(runtime.GOARCH)
	} else {
		DEBUG = false
		fmt.Println(runtime.GOOS)
		fmt.Println(runtime.GOARCH)
	}
}
