package main

import (
	"fmt"
	"runtime"
)

func main() {

	fmt.Println("Hello from CurrencyExchanger")
	PrintGoVersion()
}

func PrintGoVersion() {
	fmt.Println("Go version:", runtime.Version())
}
