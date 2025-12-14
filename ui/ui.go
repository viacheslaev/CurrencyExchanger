package ui

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/viacheslaev/CurrencyExchanger/api"
	"github.com/viacheslaev/CurrencyExchanger/model"
	"github.com/viacheslaev/CurrencyExchanger/utils"
)

func Start() {
	for {
		switch mainMenu() {
		case "1":
			currentRatesTable()
		case "2":
			// TODO: exchange currency
			fmt.Println("Currency exchange is not implemented yet 🚧")
		case "0":
			fmt.Println("Bye 👋")
			return
		default:
			fmt.Println("Invalid option")
		}
	}
}

func mainMenu() string {
	fmt.Println("\n=== Currency Exchanger ===")
	fmt.Println("1 - Show currency rates")
	fmt.Println("2 - Exchange")
	fmt.Println("0 - Exit")
	fmt.Print("Choose option: ")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func currentRatesTable() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	data, err := api.FetchRates(ctx)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("\nКурсы валют по данным ЦБ РФ,", utils.FormatCBRDate(data.Date))
	fmt.Println("----------------------------------------------------------")
	fmt.Printf("%-5s %-20s %-8s %-8s %-8s\n", "CODE", "NAME", "NOMINAL", "TODAY", "YESTERDAY")
	fmt.Println("----------------------------------------------------------")

	for _, code := range model.DefaultCodes() {
		if currency, ok := data.Valute[code]; ok {
			fmt.Printf(
				"%-5s %-20s %-8d %-8.2f %-8.2f\n",
				currency.CharCode,
				currency.Name,
				currency.Nominal,
				currency.Value,
				currency.Previous,
			)
		}
	}

	fmt.Println("----------------------------------------------------------")
}
