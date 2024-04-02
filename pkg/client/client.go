package client

import (
	"github.com/PayRam/go-tron/internal/rpc"
	"github.com/PayRam/go-tron/pkg/models"
)

// Client represents a client for TRONGrid API.
type Client struct {
	BaseURL string
	APIKey  string
}

// TRONGridClient defines the interface for interacting with TRONGrid.
type TRONGridClient interface {
	GetNowBlock() (models.Block, error)
	GetBlock(idOrNum int64, Detail bool) (models.Block, error)
}

// NewClient creates a new TRONGrid API client with the specified URL and API key.
func NewClient(baseURL, apiKey string) *rpc.RPCClient {
	return rpc.NewRPCClient(baseURL, apiKey)
}
