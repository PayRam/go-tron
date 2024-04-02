package models

// Contract represents the contract part of a transaction's raw data.
type Contract struct {
	Parameter Parameter `json:"parameter"`
	Type      string    `json:"type"`
}

type Ret struct {
	ContractRet string `json:"contractRet"`
}

type Parameter struct {
	Value   Value  `json:"value"`
	TypeURL string `json:"type_url"`
}

type Value struct {
	Data            string `json:"data,omitempty"`
	OwnerAddress    string `json:"owner_address,omitempty"`
	ContractAddress string `json:"contract_address,omitempty"`
	CallValue       int64  `json:"call_value,omitempty"`
	Amount          int64  `json:"amount,omitempty"`     // For transfer amount in TransferContract
	AssetName       string `json:"asset_name,omitempty"` // For token transfers
	ToAddress       string `json:"to_address,omitempty"` // For recipient address in transfers
}
