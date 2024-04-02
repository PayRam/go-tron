package models

// TransactionRawData holds the raw data part of a transaction.
type TransactionRawData struct {
	Contract      []Contract `json:"contract"`
	RefBlockBytes string     `json:"ref_block_bytes"`
	RefBlockHash  string     `json:"ref_block_hash"`
	Expiration    int64      `json:"expiration"`
	Data          string     `json:"data"`
	Timestamp     int64      `json:"timestamp"`
	FeeLimit      int64      `json:"fee_limit"`
}
