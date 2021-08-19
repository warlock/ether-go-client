package main

import (
	"fmt"

	web3 "github.com/umbracle/go-web3"
	"github.com/umbracle/go-web3/jsonrpc"
)

func main() {
	client, err := jsonrpc.NewClient("https://ropsten.infura.io/v3/ID")
	if err != nil {
		panic(err)
	}

	number, err := client.Eth().BlockNumber()
	if err != nil {
		panic(err)
	}

	header, err := client.Eth().GetBlockByNumber(web3.BlockNumber(number), true)
	if err != nil {
		panic(err)
	}

	fmt.Println(header)

	balance, err := client.Eth().GetBalance(web3.HexToAddress("WALLET"), web3.BlockNumber(number))
	if err != nil {
		panic(err)
	}

	fmt.Println(balance)
}
