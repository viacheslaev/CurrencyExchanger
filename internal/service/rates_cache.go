package service

import (
	"context"
	"sync"
	"time"

	"github.com/viacheslaev/CurrencyExchanger/model/currency"
)

type ratesCache struct {
	mu        sync.Mutex
	data      *currency.CBRResponse
	expiresAt time.Time
}

var cache = &ratesCache{}

// fetchRates holds FetchRates function to allow replacement in tests
var fetchRates = FetchRates

// FetchRatesCached fetch currency rates from cache or gets from CBR Api if cache expired
func FetchRatesCached(ctx context.Context) (*currency.CBRResponse, error) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	// Return cache if is valid
	if cache.data != nil && time.Now().Before(cache.expiresAt) {
		return cache.data, nil
	}

	// Fetch fresh data
	data, err := fetchRates(ctx)
	if err != nil {
		return nil, err
	}

	// Cache until next day
	cache.data = data
	cache.expiresAt = nextMidnight()

	return data, nil
}

func nextMidnight() time.Time {
	now := time.Now()
	return time.Date(
		now.Year(),
		now.Month(),
		now.Day()+1,
		0, 0, 0, 0,
		now.Location(),
	)
}
