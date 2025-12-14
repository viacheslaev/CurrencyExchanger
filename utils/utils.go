package utils

import "time"

func FormatCBRDate(rawDate string) string {
	t, err := time.Parse(time.RFC3339, rawDate)
	if err != nil {
		return rawDate
	}

	return t.Format("02 January 2006 15:04")
}
