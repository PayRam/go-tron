package trxinit

import (
	"github.com/PayRam/go-tron/internal/rpc"
	"github.com/PayRam/go-tron/pkg/trxclient"
)

// NewClient creates a new TRONGrid API trxclient with the specified URL and API key.
func NewClient(baseURL, apiKey string) trxclient.Client {
	return rpc.NewTRXClient(baseURL, apiKey)
}
