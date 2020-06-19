package blockchain

import (
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
	bc.MineBlock("test1")
	bc.MineBlock("test2")

	if !bc.validChain() {
		t.Error("Blockchain should be valid")
	}
}

// Test if validChain will invalidate an invalid chain
func TestBlockchainInvalidation(t *testing.T) {
	bc := Genisis()
	bc.MineBlock("test1")
	bc.MineBlock("test2")
	bc.blocks[1].value = "INSERTED"

	if bc.validChain() {
		t.Error("Blockchain should be invalid due to injected value")
	}
}

func TestReplaceChain(t *testing.T) {
	bc := Genisis()
	bc.MineBlock("block1")

	nc := Genisis()
	nc.MineBlock("newBlock1")
	nc.MineBlock("newBlock2")

	bc.ReplaceChain(nc)

	if len(bc.blocks) != 3 {
		t.Error("Blockchain should have length of three")
	}

	if bc.blocks[1].value != "newBlock1" || bc.blocks[2].value != "newBlock2" {
		t.Errorf("Blockchain should replace old blocks with new blocks")
	}
}

func TestDoNotReplaceChain(t *testing.T) {
	bc := Genisis()
	bc.MineBlock("block1")
	bc.MineBlock("block2")

	nc := Genisis()
	nc.MineBlock("newBlock1")

	bc.ReplaceChain(nc)

	if len(bc.blocks) != 3 {
		t.Error("Blockchain should have length of three")
	}

	if bc.blocks[1].value != "block1" || bc.blocks[2].value != "block2" {
		t.Errorf("Blockchain blocks should remain the same")
	}
}
