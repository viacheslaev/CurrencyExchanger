package main

import (
	"context"
	"fmt"
	"time"

	"github.com/viacheslaev/CurrencyExchanger/api"
	"github.com/viacheslaev/CurrencyExchanger/ui"
)

func main() {
	for {
		switch ui.ShowMenu() {
		case "1":
			showCurrency("USD")
		case "2":
			showCurrency("EUR")
		case "0":
			fmt.Println("Bye 👋")
			return
		default:
			fmt.Println("Invalid option")
		}
	}
}

func showCurrency(code string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	data, err := api.FetchRates(ctx)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	currency, ok := data.Valute[code]
	if !ok {
		fmt.Println("Currency not found")
		return
	}

	fmt.Printf("\n%s (%s)\n", currency.Name, currency.CharCode)
	fmt.Printf("Nominal: %d\n", currency.Nominal)
	fmt.Printf("Rate: %.4f RUB\n", currency.Value)
	fmt.Printf("Previous: %.4f RUB\n", currency.Previous)
	fmt.Printf("Date: %s\n", data.Date)
}
