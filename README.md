# go-tron

## Usage

Import in your project
```
require github.com/PayRam/go-tron v0.1.0
```
RPC api calls
```
	tronCleint := client.NewClient("https://api.trongrid.io", "api-key")

	block, _ := tronCleint.GetBlock(32880248, false)

	log.Println(block)
```
