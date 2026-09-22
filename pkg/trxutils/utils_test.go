package trxutils

import (
	"encoding/hex"
	"testing"
)

// A real PayRam merchant hot wallet whose 20-byte account begins with 0x41 —
// the 1-in-256 case — and USDT on Tron mainnet as an ordinary account.
const (
	merchantAccountHex = "413271c913137a2ff70990b1adaab765f776c5a2"
	merchantAddress    = "TFuwMUJPkhNpT5d4GnRgdjetAUDbxKEYkS"
	usdtAccountHex     = "a614f803b6fd780986a42c78ec9c7f77e6ded13c"
	usdtAddress        = "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t"
)

func TestHexToAddress_LengthDecidesThePrefix(t *testing.T) {
	for _, tc := range []struct{ name, in, want string }{
		{"account starting with 41 (the bug)", merchantAccountHex, merchantAddress},
		{"same account, already prefixed", "41" + merchantAccountHex, merchantAddress},
		{"ordinary account", usdtAccountHex, usdtAddress},
		{"ordinary account, already prefixed", "41" + usdtAccountHex, usdtAddress},
	} {
		got, err := HexToAddress(tc.in)
		if err != nil {
			t.Errorf("%s: HexToAddress(%s) error: %v", tc.name, tc.in, err)
			continue
		}
		if got != tc.want {
			t.Errorf("%s: HexToAddress(%s) = %s, want %s", tc.name, tc.in, got, tc.want)
		}
	}
}

func TestHexToAddress_RejectsWrongShapes(t *testing.T) {
	for name, in := range map[string]string{
		"empty":                    "",
		"38 characters":            usdtAccountHex[:38],
		"44 characters":            "41" + usdtAccountHex + "00",
		"42 without the 41 prefix": "42" + usdtAccountHex,
		"not hex":                  "zz" + usdtAccountHex[2:],
	} {
		if got, err := HexToAddress(in); err == nil {
			t.Errorf("%s: HexToAddress(%q) = %s, want an error", name, in, got)
		}
	}
}

// ToBase58 must not write its checksum into the caller's slice.
func TestToBase58_DoesNotMutateInput(t *testing.T) {
	body, _ := hex.DecodeString("41" + usdtAccountHex)
	buf := make([]byte, len(body), 64)
	copy(buf, body)
	spare := buf[len(buf) : len(buf)+4]
	if got := ToBase58(buf); got != usdtAddress {
		t.Fatalf("ToBase58 = %s, want %s", got, usdtAddress)
	}
	for _, b := range spare {
		if b != 0 {
			t.Fatal("ToBase58 wrote past the end of its input slice")
		}
	}
}
