package ui

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ShowMenu() string {
	fmt.Println("\n=== Currency Exchanger ===")
	fmt.Println("1 - USD")
	fmt.Println("2 - EUR")
	fmt.Println("0 - Exit")
	fmt.Print("Choose option: ")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}
