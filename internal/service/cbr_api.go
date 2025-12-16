package service

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"time"

	"github.com/viacheslaev/CurrencyExchanger/internal/config"
	"github.com/viacheslaev/CurrencyExchanger/model/currency"
)

var client = &http.Client{
	Timeout: 5 * time.Second,
}

func FetchRates(ctx context.Context) (*currency.CBRResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, config.CBRUrl(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("CBR API error: status %d", resp.StatusCode)
	}

	var data currency.CBRResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	normalizeRates(&data)

	return &data, nil
}

// normalizeRates rounds currency rates in CBRResponse ( 79.4494999 -> 79.45 )
func normalizeRates(data *currency.CBRResponse) {
	for code, cur := range data.Valute {
		cur.Value = math.Round(cur.Value*100) / 100
		cur.Previous = math.Round(cur.Previous*100) / 100
		data.Valute[code] = cur
	}
}
