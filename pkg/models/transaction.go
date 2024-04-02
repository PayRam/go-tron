package models

// Transaction represents a transaction in TRONGrid.
type Transaction struct {
	Ret        []Ret              `json:"ret"`
	TxID       string             `json:"txID"`
	RawData    TransactionRawData `json:"raw_data"`
	Signature  []string           `json:"signature"`
	RawDataHex string             `json:"raw_data_hex"`
}
