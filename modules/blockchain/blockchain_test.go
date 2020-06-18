package blockchain

import "testing"

// test the block hash function
func TestBlockHash(t *testing.T) {
	blockOne := &Block{
		timestamp:    1592470203316,
		previousHash: "Genisis",
		value:        "foo",
		difficulty:   4,
		nonce:        1029561,
	}

	blockOne.hashBlock()

	blockTwo := &Block{
		timestamp:    1592470203316,
		previousHash: "Genisis",
		value:        "foo",
		difficulty:   4,
		nonce:        1029561,
	}

	blockTwo.hashBlock()

	if blockOne.hash != blockTwo.hash {
		t.Errorf("identical blocks hash should match : %v != %v", blockOne.hash, blockTwo.hash)
	}

	blockTwo.value = "bar"
	blockTwo.hashBlock()

	if blockOne.hash == blockTwo.hash {
		t.Error("different blocks should have different hashes")
	}
}
