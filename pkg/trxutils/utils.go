package trxutils

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"github.com/PayRam/go-tron/pkg/models"
	"github.com/btcsuite/btcutil/base58"
	"math/big"
	"strings"
)

// ToBase58 encodes input bytes using ToBase58 encoding.
func ToBase58(input []byte) string {
	checksum := doubleSHA256(input)
	first4Bytes := checksum[:4] // First 4 bytes of the checksum

	fullPayload := append(input, first4Bytes...) // Append the checksum to the input

	encoded := base58.Encode(fullPayload)

	return encoded
}

// doubleSHA256 computes SHA256(SHA256(data)) and returns the resulting hash.
func doubleSHA256(data []byte) []byte {
	firstHash := sha256.Sum256(data)
	secondHash := sha256.Sum256(firstHash[:])
	return secondHash[:]
}

func DecodeTransferData(data string) (*models.TransferData, error) {
	if len(data) < 136 { // 8 chars for MethodID + 64 chars for ToAddress + 64 chars for Value = 136 chars
		return nil, errors.New("data string not long enough to contain method ID, to address, and value")
	}
	methodID := data[:8]
	toHex := data[8:72]   // Next 32 bytes after methodID
	valueHex := data[72:] // Next 32 bytes after toHex

	toAddressHex := toHex[len(toHex)-40:]

	// Convert the hex value to a big integer
	valueBigInt := new(big.Int)
	valueBigInt.SetString(valueHex, 16)

	return &models.TransferData{
		MethodID:  methodID,
		ToAddress: toAddressHex,
		Value:     *valueBigInt,
	}, nil
}

func HexToAddress(hexAddr string) (string, error) {
	// Add "41" prefix to the hex address (TRON addresses start with "41" which is "T" in Base58)
	var prefixedHexAddr string
	if strings.HasPrefix(hexAddr, "41") {
		if len(hexAddr) == 40 { // Excludes "41" prefix
			prefixedHexAddr = "41" + hexAddr
		} else if len(hexAddr) != 42 { // Includes "41" prefix
			return "", errors.New("invalid address length with '41' prefix")
		}
		prefixedHexAddr = hexAddr
	} else {
		if len(hexAddr) != 40 { // Excludes "41" prefix
			return "", errors.New("invalid address length without '41' prefix")
		}
		prefixedHexAddr = "41" + hexAddr
	}

	// Decode the hex string to bytes
	addrBytes, err := hex.DecodeString(prefixedHexAddr)
	if err != nil {
		return "", err
	}

	// Double SHA-256 hash
	hash := sha256.Sum256(addrBytes)
	hash = sha256.Sum256(hash[:])

	// Append first 4 bytes of hash as checksum
	checksummedBytes := append(addrBytes, hash[:4]...)

	// Convert to Base58
	base58Addr := base58.Encode(checksummedBytes)

	return base58Addr, nil
}
