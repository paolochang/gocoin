package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/paolochang/gocoin/explorer"
	"github.com/paolochang/gocoin/rest"
)

func usage() {
	fmt.Printf("Welcome to Go Coin\n\n")
	fmt.Printf("Please us the following flags:\n\n")
	fmt.Printf("-port:		Start the PORT of the server\n")
	fmt.Printf("-mode:		Choose between 'html' and 'rest'\n\n")
	os.Exit(0)
}

func Strat() {
	if len(os.Args) == 1 {
		usage()
	}

	mode := flag.String("mode", "rest", "Choose between 'html' and 'rest'")
	port := flag.Int("port", 4000, "Set port of the server")

	flag.Parse()

	switch *mode {
	case "rest":
		// start rest api
		rest.Start(*port)
	case "html":
		// start html explorer
		explorer.Start(*port)
	default:
		usage()
	}}