package service

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/viacheslaev/CurrencyExchanger/model/currency"
)

func ExchangeCurrency() {
	sourceCurrency, sourceAmount := readSourceCurrencyAndAmount()
	targetCurrency := readTargetCurrency()

	data, err := fetchRatesWithTimeout()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// --- CASE 1: target currency code is RUB ---
	if targetCurrency == currency.RUB {
		from, ok := data.Valute[sourceCurrency]
		if !ok {
			fmt.Println("Unknown currency:", sourceCurrency)
			return
		}

		result := exchangeToRUB(sourceAmount, from)

		fmt.Printf("\nResult: %.2f RUB\n", result)
		return
	}

	// --- CASE 2: currency -> all other currency codes ---
	from, ok := data.Valute[sourceCurrency]
	if !ok {
		fmt.Println("Unknown currency:", sourceCurrency)
		return
	}

	to, ok := data.Valute[targetCurrency]
	if !ok {
		fmt.Println("Unknown currency:", targetCurrency)
		return
	}

	result := exchange(sourceAmount, from, to)

	fmt.Printf("\nResult: %.2f %s\n", result, targetCurrency)
}

func exchange(sourceAmount float64, from currency.Currency, to currency.Currency) float64 {
	result := sourceAmount * from.Value * float64(to.Nominal) /
		(float64(from.Nominal) * to.Value)
	return result
}

func exchangeToRUB(sourceAmount float64, from currency.Currency) float64 {
	result := sourceAmount * from.Value / float64(from.Nominal)
	return result
}

func readSourceCurrencyAndAmount() (currency.Code, float64) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("\nEnter amount and currency (e.g. 500 USD ): ")
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		parts := strings.Fields(line)
		if len(parts) != 2 {
			fmt.Println("Invalid format")
			continue
		}

		amount, err := strconv.ParseFloat(parts[0], 64)
		if err != nil || amount <= 0 {
			fmt.Println("Invalid amount")
			continue
		}

		var currencyCode = parts[1]
		return currency.Code(strings.ToUpper(currencyCode)), amount
	}
}

func readTargetCurrency() currency.Code {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("\nEnter target currency: ")
	line, _ := reader.ReadString('\n')

	return currency.Code(strings.ToUpper(strings.TrimSpace(line)))
}

func fetchRatesWithTimeout() (*currency.CBRResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return FetchRatesCached(ctx)
}
