package blockchain

import (
	"fmt"
	"testing"
)

func testWalletGeneration(t *testing.T) {
	// Generate new wallet
	w, err := GenerateWallet("")

	if err != nil {
		t.Errorf("Error generating wallet : %v", err)
	}

	// convert key to private address
	address := privateKeyToB64(w.key)

	if len(address) != 256 {
		t.Errorf("Error generating private wallet hash : %v", address)
		t.Errorf("address length %v should be 256 characters", address)
	}

	// regenerate wallet from private address
	w2, err := GenerateWallet(address)

	if walletToString(w) != walletToString(w2) {
		t.Error("Error regenerating wallet from private address")
	}
}

func walletToString(w *Wallet) string {
	return fmt.Sprintf("%v", w.key)
}
