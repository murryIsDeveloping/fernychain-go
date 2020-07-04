package blockchain

import (
	"testing"
)

func TestWalletGeneration(t *testing.T) {
	// Generate new wallet
	w, err := GenerateWallet("")
	if err != nil {
		t.Errorf("Error generating wallet : %v", err)
	}

	// convert key to private address
	address := privateKeyToB64(w.key)

	// regenerate wallet from private address
	w2, err := GenerateWallet(address)

	if err != nil {
		t.Errorf("Error generating wallet : %v", err)
	}

	if w.PublicKey() != w2.PublicKey() {
		t.Errorf("Error regenerating wallet from private address: %v != %v", w.PublicKey(), w2.PublicKey())
	}
}
