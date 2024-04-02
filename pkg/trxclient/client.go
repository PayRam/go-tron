package trxclient

import (
	"github.com/PayRam/go-tron/pkg/models"
)

// Client defines the interface for interacting with TRONGrid.
type Client interface {
	GetNowBlock() (models.Block, error)
	GetBlock(idOrNum int64, Detail bool) (models.Block, error)
}
