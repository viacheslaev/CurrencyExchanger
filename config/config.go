package config

import "os"

const defaultCBRURL = "https://www.cbr-xml-daily.ru/daily_json.js"

func CBRUrl() string {
	if url := os.Getenv("CBR_URL"); url != "" {
		return url
	}
	return defaultCBRURL
}
