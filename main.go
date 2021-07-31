package main

import (
	"fmt"
	"os"
)

func usage() {
	fmt.Printf("Welcome to Go Coin\n\n")
	fmt.Printf("Please us the following commands:\n\n")
	fmt.Printf("explorer:		Start the HTML Explorer\n")
	fmt.Printf("rest:				Start the REST API (recommended)\n\n")
	os.Exit(0)
}

func main() {
	// go explorer.Start(3000)
	// rest.Start(4000)
	if len(os.Args) < 2 {
		usage()
	}

	switch os.Args[1] {
	case "explorer":
		fmt.Println("Start Explorer")
	case "rest":
		fmt.Println("Start REST API")
	default:
		usage()
	}
}

