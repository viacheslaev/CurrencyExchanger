package service

import (
	"math"
	"testing"

	"github.com/viacheslaev/CurrencyExchanger/model/currency"
)

// epsilon: max allowed difference when comparing floats (≈ 0.001 RUB)
const epsilon = 1e-3

func TestExchangeUSDToRUB(t *testing.T) {
	// Given
	usd := currency.Currency{
		Nominal: 1,
		Value:   78.10,
	}

	// When
	got := exchangeToRUB(100, usd)

	// Then
	want := 7810.00
	if math.Abs(got-want) > epsilon {
		t.Fatalf("want %.2f, got %.2f", want, got)
	}
}

func TestExchangeToRUB(t *testing.T) {
	// Given
	tests := []struct {
		name   string
		amount float64
		from   currency.Currency
		want   float64
	}{
		{
			name:   "1000 USD",
			amount: 1000,
			from:   currency.Currency{Nominal: 1, Value: 78.10},
			want:   78100.00,
		},
		{
			name:   "0 EUR",
			amount: 0,
			from:   currency.Currency{Nominal: 1, Value: 92.00},
			want:   0.00,
		},
		{
			name:   "100 EUR",
			amount: 100,
			from:   currency.Currency{Nominal: 1, Value: 92.00},
			want:   9200.00,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// When
			got := exchangeToRUB(tt.amount, tt.from)

			// Then
			if math.Abs(got-tt.want) > epsilon {
				t.Fatalf("want %.2f, got %.2f", tt.want, got)
			}
		})
	}
}

func TestExchangeUSDToEUR(t *testing.T) {
	// Given
	tests := []struct {
		name   string
		amount float64
		from   currency.Currency
		to     currency.Currency
		want   float64
	}{
		{
			name:   "1 USD to EUR",
			amount: 1,
			from:   currency.Currency{Nominal: 1, Value: 80.0},
			to:     currency.Currency{Nominal: 1, Value: 100.0},
			want:   0.80,
		},
		{
			name:   "10 USD to EUR",
			amount: 10,
			from:   currency.Currency{Nominal: 1, Value: 80.0},
			to:     currency.Currency{Nominal: 1, Value: 100.0},
			want:   8.00,
		},
		{
			name:   "0 USD to EUR",
			amount: 0,
			from:   currency.Currency{Nominal: 1, Value: 80.0},
			to:     currency.Currency{Nominal: 1, Value: 100.0},
			want:   0.00,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// When
			got := exchange(tt.amount, tt.from, tt.to)

			// Then
			if math.Abs(got-tt.want) > epsilon {
				t.Fatalf("want %.2f, got %.2f", tt.want, got)
			}
		})
	}
}
