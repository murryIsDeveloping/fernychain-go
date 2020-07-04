package blockchain

import (
	"testing"
)

func TestTransactionCreation(t *testing.T) {
	bc := Genisis()
	w, _ := GenerateWallet("")
	transactionOne := w.CreateTransaction(bc)
	// fake transaction value
	transactionOne.input.value = 100.0
	err := transactionOne.AddTransaction(w, 10.0, "TestAddress")

	if err != nil {
		t.Errorf("Error Adding transaction: %v", err)
	}

	if transactionOne.input.value != 90.0 {
		t.Errorf("Transaction should be removed from input amount")
	}
}
