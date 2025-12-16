package flags

import (
	"flag"
	"fmt"
	"os"
)

const appVersion = "1.0"

// Handle parses command-line flags, for example -version.
func Handle() {
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
