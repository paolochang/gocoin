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
	fmt.Printf("Please use the following flags:\n\n")
	fmt.Printf("-mode:		Select the mode from 'all', 'rest', or 'html' (default 'all')\n")
	fmt.Printf("-port:		Start the PORT of the server (default 4000)\n\n")
	os.Exit(0)
}

func Strat() {

	mode := flag.String("mode", "all", "Select the mode from 'html', 'rest', or 'all'")
	port := flag.Int("port", 4000, "Set port of the server")

	flag.Parse()

	switch *mode {
	case "rest":
		fmt.Println("Start REST API server")
		rest.Start(*port)
	case "html":
		fmt.Println("Start HTML server")
		explorer.Start(*port)
	case "all":
		fmt.Println("Start REST API and HTML server\nTo see the CLI command usage, run: go run main.go -mode=help")
		go explorer.Start(*port)
		rest.Start(*port+1)
	case "help":
		usage()
	}}