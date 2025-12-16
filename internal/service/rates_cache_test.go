package service

import (
	"context"
	"testing"
	"time"

	"github.com/viacheslaev/CurrencyExchanger/model/currency"
)

func TestFetchRatesCached_UsesCache(t *testing.T) {
	// Given
	ctx := context.Background()

	callCount := 0

	// Mock fetchRates
	fetchRates = func(ctx context.Context) (*currency.CBRResponse, error) {
		callCount++
		return &currency.CBRResponse{
			Date: time.Now().Format(time.RFC3339),
			Valute: map[currency.Code]currency.Currency{
				currency.USD: {Nominal: 1, Value: 80},
			},
		}, nil
	}

	// After test return real FetchRates instead mock
	defer func() {
		fetchRates = FetchRates
		cache.data = nil
	}()

	// When
	_, err1 := FetchRatesCached(ctx)
	_, err2 := FetchRatesCached(ctx)

	// Then
	if err1 != nil || err2 != nil {
		t.Fatalf("unexpected error")
	}

	if callCount != 1 {
		t.Fatalf("Want fetchRates to be called once, got %d", callCount)
	}
}
