package main

import (
	"github.com/PayRam/go-tron/pkg/client"
	"log"
)

func main() {
	// This is the main function
	// This is the entry point of the program
	// This is the first

	tronCleint := client.NewClient("https://api.trongrid.io", "api-key")

	block, _ := tronCleint.GetBlock(32880248, false)

	log.Println(block)
}
