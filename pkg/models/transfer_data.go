package models

import "math/big"

type TransferData struct {
	MethodID  string  `json:"method"`
	ToAddress string  `json:"to_address"`
	Value     big.Int `json:"value"`
}
