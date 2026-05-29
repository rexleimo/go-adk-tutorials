package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("Go Version: %s\n", runtime.Version())
	fmt.Printf("GOOS: %s, GOARCH: %s\n", runtime.GOOS, runtime.GOARCH)
}