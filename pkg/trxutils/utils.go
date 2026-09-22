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

// ToBase58 base58check-encodes input exactly as given: it adds no prefix, so
// for a Tron address the caller must pass all 21 bytes (0x41 + account).
func ToBase58(input []byte) string {
	checksum := doubleSHA256(input)
	first4Bytes := checksum[:4] // First 4 bytes of the checksum

	// Copy before appending: append(input, …) writes the checksum into the
	// caller's backing array whenever input has spare capacity.
	fullPayload := make([]byte, 0, len(input)+len(first4Bytes))
	fullPayload = append(fullPayload, input...)
	fullPayload = append(fullPayload, first4Bytes...)

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

// HexToAddress encodes a Tron address (the "T…" form) from hex.
//
// The LENGTH decides whether Tron's 0x41 prefix is already there: 40 hex
// characters are a bare 20-byte account and always get the prefix; 42 must
// already start with "41". Deciding by content instead — "it starts with 41,
// so the prefix is present" — is wrong for about 1 account in 256, because
// 0x41 is also an ordinary first byte for a random account. Until this fix, a
// missing `else` sent exactly those accounts down the already-prefixed path,
// and they came back as a 33-character string that is not an address (a
// PayRam merchant's hot wallet read 0 TRX because of it, 2026-09-19).
func HexToAddress(hexAddr string) (string, error) {
	var prefixedHexAddr string
	switch len(hexAddr) {
	case 40: // bare account
		prefixedHexAddr = "41" + hexAddr
	case 42: // already prefixed
		if !strings.HasPrefix(hexAddr, "41") {
			return "", errors.New("invalid address: 42 hex characters must start with the '41' prefix")
		}
		prefixedHexAddr = hexAddr
	default:
		return "", errors.New("invalid address length: want 40 hex characters (bare account) or 42 (with the '41' prefix)")
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
