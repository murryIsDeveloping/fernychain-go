package blockchain

import (
	"fmt"
	"testing"
)

// Test Genisis creation
func TestGenisisBlockchain(t *testing.T) {
	bc := Genisis()
	l := len(bc.blocks)

	if l != 1 {
		t.Errorf("Blockchain should only contain 1 genisis block, contains %v blocks", l)
	}
}

// Test if validChain will validate a valid chain
func TestBlockchainValidation(t *testing.T) {
	bc := Genisis()
	w, _ := GenerateWallet("")
	transactionOne := w.CreateTransaction(bc)
	transactionTwo := w.CreateTransaction(bc)
	transactionOne.AddTransaction(w, 0.0, "FakeAddress")
	transactionTwo.AddTransaction(w, 0.0, "AnotherAddress")

	bc.MineBlock([]Transaction{*transactionOne})
	bc.MineBlock([]Transaction{*transactionTwo})

	if !bc.validate() {
		t.Error("Blockchain should be valid")
	}
}

// Test if validChain will invalidate an invalid chain
func TestBlockchainInvalidation(t *testing.T) {
	bc := Genisis()
	w, _ := GenerateWallet("")
	transactionOne := w.CreateTransaction(bc)
	transactionOne.input.value = 100.0
	transactionTwo := w.CreateTransaction(bc)
	transactionTwo.input.value = 100.0
	transactionThree := w.CreateTransaction(bc)
	transactionThree.input.value = 100.0
	transactionOne.AddTransaction(w, 10.0, "FakeAddress")
	transactionTwo.AddTransaction(w, 20.0, "AnotherAddress")
	transactionThree.AddTransaction(w, 30.0, "HackedAddress")

	bc.MineBlock([]Transaction{*transactionOne})
	bc.MineBlock([]Transaction{*transactionTwo})

	bc.blocks[1].value = []Transaction{*transactionThree}

	if bc.validate() {
		t.Error("Blockchain should be invalid due to injected value")
	}
}

func TestReplaceChain(t *testing.T) {
	bc := Genisis()
	w, _ := GenerateWallet("")
	transactionOne := w.CreateTransaction(bc)
	transactionOne.input.value = 100.0

	transactionOne.AddTransaction(w, 10.0, "randomAddress")

	bc.MineBlock([]Transaction{*transactionOne})

	nc := Genisis()
	nTransactionOne := w.CreateTransaction(nc)
	nTransactionOne.input.value = 100.0
	nTransactionTwo := w.CreateTransaction(nc)
	nTransactionTwo.input.value = 100.0

	nTransactionOne.AddTransaction(w, 10.0, "FakeAddress")
	nTransactionTwo.AddTransaction(w, 10.0, "AnotherAddress")

	nc.MineBlock([]Transaction{*nTransactionOne})
	nc.MineBlock([]Transaction{*nTransactionTwo})

	bc.ReplaceChain(nc)

	if len(bc.blocks) != 3 {
		t.Error("Blockchain should have length of three")
	}

	if fmt.Sprintf("%v", bc.blocks) != fmt.Sprintf("%v", nc.blocks) {
		t.Errorf("Blockchain should replace old blocks with new blocks")
	}
}

func TestDoNotReplaceChain(t *testing.T) {
	bc := Genisis()
	w, _ := GenerateWallet("")
	transactionOne := w.CreateTransaction(bc)
	transactionOne.input.value = 100.0
	transactionTwo := w.CreateTransaction(bc)
	transactionTwo.input.value = 100.0

	transactionOne.AddTransaction(w, 10.0, "randomAddress")
	transactionTwo.AddTransaction(w, 10.0, "AnotherAddress")

	bc.MineBlock([]Transaction{*transactionOne})
	bc.MineBlock([]Transaction{*transactionTwo})

	nc := Genisis()
	nTransactionOne := w.CreateTransaction(nc)
	nTransactionOne.input.value = 100.0

	nTransactionOne.AddTransaction(w, 10.0, "FakeAddress")

	nc.MineBlock([]Transaction{*nTransactionOne})

	bc.ReplaceChain(nc)

	if len(bc.blocks) != 3 {
		t.Error("Blockchain should have length of three")
	}

	if fmt.Sprintf("%v", bc.blocks) == fmt.Sprintf("%v", nc.blocks) {
		t.Errorf("Blockchain blocks should remain the same")
	}
}
