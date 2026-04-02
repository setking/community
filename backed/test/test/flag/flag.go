package main

import "github.com/spf13/pflag"

var (
	cfg  = pflag.StringP("config", "c", "", "configuration file")
	help = pflag.BoolP("help", "h", false, "print help message")
)

func main() {
	pflag.Parse()

	if *help {
		pflag.Usage()
		return
	}

	if *cfg == "" {
		pflag.PrintDefaults()
		return
	}

	// Use the loaded configuration file
	// ...
}
