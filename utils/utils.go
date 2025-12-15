package utils

import (
	"flag"
	"fmt"
	"os"
	"time"
)

const appVersion = "1.0"

func FormatCBRDate(rawDate string) string {
	t, err := time.Parse(time.RFC3339, rawDate)
	if err != nil {
		return rawDate
	}

	return t.Format("02 January 2006 15:04")
}

func HandleFlags() {
	showVersion := flag.Bool("version", false, "Show application version")
	flag.Parse()

	if *showVersion {
		printVersion()
		os.Exit(0)
	}
}

func printVersion() {
	fmt.Println("version", appVersion)
}
