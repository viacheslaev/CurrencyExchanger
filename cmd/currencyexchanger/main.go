package main

import (
	"github.com/viacheslaev/CurrencyExchanger/internal/flags"
	"github.com/viacheslaev/CurrencyExchanger/internal/ui"
)

func init() {
	flags.Handle()
}

func main() {
	ui.Start()
}
