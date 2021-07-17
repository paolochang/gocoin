package main

import (
	"github.com/paolochang/gocoin/explorer"
	"github.com/paolochang/gocoin/rest"
)

func main() {
	go explorer.Start(3000)
	rest.Start(4000)
}

