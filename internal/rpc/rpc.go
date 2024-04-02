package rpc

import (
	"encoding/json"
	"fmt"
	"github.com/PayRam/go-tron/internal/utils"
	"github.com/PayRam/go-tron/pkg/models"
	"github.com/PayRam/go-tron/pkg/trxclient"
	"strconv"
)

// client implements TRONGridClient interface for TRONGrid API.
type client struct {
	baseURL string
	apiKey  string
}

// Ensure RPCClient implements the trxclient.Client interface.
var _ trxclient.Client = (*client)(nil)

// NewTRXClient creates a new RPC trxclient.
func NewTRXClient(baseURL, apiKey string) trxclient.Client {
	return &client{
		baseURL: baseURL,
		apiKey:  apiKey,
	}
}

func (c *client) GetNowBlock() (models.Block, error) {
	responseBody, err := utils.MakeRequest(c.baseURL, "/wallet/getnowblock", nil)

	if err != nil {
		return models.Block{}, err
	}

	var response models.Block
	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		return models.Block{}, err
	}

	return response, nil
}

func (c *client) GetBlock(idOrNum int64, Detail bool) (models.Block, error) {
	// Construct the request body
	requestBody := map[string]interface{}{
		"id_or_num": strconv.Itoa(int(idOrNum)),
		"detail":    Detail,
	}

	responseBody, err := utils.MakeRequest(c.baseURL, "/wallet/getblock", requestBody)

	if err != nil {
		return models.Block{}, fmt.Errorf("error making request to TRONGrid: %v", err)
	}

	var response models.Block
	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		return models.Block{}, fmt.Errorf("error parsing response from TRONGrid: %v", err)
	}

	return response, nil
}

func (c *client) GetBlockByLimitNext(startNum int64, endNum int64) ([]models.Block, error) {
	// Construct the request body
	requestBody := map[string]interface{}{
		"startNum": startNum,
		"endNum":   endNum,
	}

	responseBody, err := utils.MakeRequest(c.baseURL, "/wallet/getblockbylimitnext", requestBody)

	if err != nil {
		return nil, fmt.Errorf("error making request to TRONGrid: %v", err)
	}

	type Block struct {
		Blocks []models.Block `json:"block"`
	}

	var response Block
	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		return nil, fmt.Errorf("error parsing response from TRONGrid: %v", err)
	}

	return response.Blocks, nil
}
