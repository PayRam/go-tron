package models

// BlockRawData holds the raw data part of the block header.
type BlockRawData struct {
	Timestamp        int64  `json:"timestamp"`
	TxTrieRoot       string `json:"txTrieRoot"`
	ParentHash       string `json:"parentHash"`
	Number           int64  `json:"number"`
	WitnessID        int64  `json:"witness_id"`
	WitnessAddress   string `json:"witness_address"`
	Version          int32  `json:"version"`
	AccountStateRoot string `json:"accountStateRoot"`
}
