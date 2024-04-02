package models

// BlockHeader represents the header part of a TRON block.
type BlockHeader struct {
	RawData          BlockRawData `json:"raw_data"`
	WitnessSignature string       `json:"witness_signature"`
}
