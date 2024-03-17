package main

import (
	"fmt"

	blockchain "github.com/nomadcoders/nomadcoin/blckchain"
)

func main() {
	chain := blockchain.GetBlockchain()
	fmt.Println(chain)
}