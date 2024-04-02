package trxutils

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/PayRam/go-tron/pkg/models"
	"github.com/btcsuite/btcutil/base58"
	"log"
	"math/big"
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

// HexToBase58 converts a hexadecimal address to a Base58 address.
func HexToBase58(hexAddr string) string {
	decoded, err := hex.DecodeString(hexAddr)
	if err != nil {
		log.Fatal("Failed to decode hex:", err)
	}
	// Encode the payload with Base58Check.
	return ToBase58(decoded)
}

func DecodeTransferData(data string) models.TransferData {
	methodID := data[:8]
	toHex := data[8:72]   // Next 32 bytes after methodID
	valueHex := data[72:] // Next 32 bytes after toHex

	toAddressHex := toHex[len(toHex)-42:]

	// Convert the hex value to a big integer
	valueBigInt := new(big.Int)
	valueBigInt.SetString(valueHex, 16)

	// Convert to Ethereum address format
	toAddress := HexToBase58(toAddressHex)

	return models.TransferData{
		MethodID:  methodID,
		ToAddress: toAddress,
		Value:     *valueBigInt,
	}
}
