package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/viacheslaev/CurrencyExchanger/model"
)

const cbrURL = "https://www.cbr-xml-daily.ru/daily_json.js"

func FetchRates(ctx context.Context) (*model.CBRResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, cbrURL, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("CBR API error: status %d", resp.StatusCode)
	}

	var data model.CBRResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return &data, nil
}
