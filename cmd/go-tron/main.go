package main

import (
	"github.com/PayRam/go-tron/pkg/trxclient/trxinit"
	"log"
)

func main() {
	// This is the main function
	// This is the entry point of the program
	// This is the first

	trxClient := trxinit.NewClient("https://api.trongrid.io", "api-key")

	block, _ := trxClient.GetBlock(32880248, false)

	log.Println(block)
}
