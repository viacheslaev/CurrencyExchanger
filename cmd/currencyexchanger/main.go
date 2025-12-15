package main

import (
	"github.com/viacheslaev/CurrencyExchanger/ui"
	"github.com/viacheslaev/CurrencyExchanger/utils"
)

func init() {
	utils.HandleFlags()
}

func main() {
	ui.Start()
}
