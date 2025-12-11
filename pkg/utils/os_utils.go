package utils

import (
	"fmt"
	"runtime"
)

func PrintGoVersion() {
	fmt.Println("Go version:", runtime.Version())
}
