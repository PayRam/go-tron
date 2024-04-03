# go-tron

## Usage

Import in your project
```
require github.com/PayRam/go-tron v0.2.0
```
RPC api calls
```
	data, err := trxutils.DecodeTransferData("a9059cbb000000000000000000000041c357d2873ffee9d8771a2dfb8a327ef87622cef1000000000000000000000000000000000000000000000000000000003b9aca00")
	if err != nil {
		return
	}
	
	log.Println(data)
	
	log.Println(trxutils.HexToAddress(data.ToAddress))

	trxClient := trxinit.NewClient("https://nile.trongrid.io", "api-key")

	block, _ := trxClient.GetNowBlock()

	log.Println(block)

	block1, err := trxClient.GetBlockByLimitNext(32880248, 32880349)
	if err != nil {
		return
	}

	log.Println(block1)
```
