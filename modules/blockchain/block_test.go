package blockchain

import (
	"testing"
)

// test the block hash function
func TestBlockHash(t *testing.T) {
	blockOne := &Block{
		timestamp:    1592470203316,
		previousHash: "Genisis",
		value:        "foo",
		difficulty:   4,
		nonce:        1029561,
	}

	blockOne.hash = hashBlock(*blockOne)

	blockTwo := &Block{
		timestamp:    1592470203316,
		previousHash: "Genisis",
		value:        "foo",
		difficulty:   4,
		nonce:        1029561,
	}

	blockTwo.hash = hashBlock(*blockTwo)

	if blockOne.hash != blockTwo.hash {
		t.Errorf("identical blocks hash should match : %v != %v", blockOne.hash, blockTwo.hash)
	}

	blockTwo.value = "bar"
	blockTwo.hash = hashBlock(*blockTwo)

	if blockOne.hash == blockTwo.hash {
		t.Error("different blocks should have different hashes")
	}
}

// Test if block has required nonce
func TestNonce(t *testing.T) {
	// Test has nonce pass
	b := &Block{
		difficulty: 4,
		hash:       "0000ncfhjdekasdnfasjdfasdf",
	}

	n := b.hasNonce()

	if !n {
		t.Error("if block hash has 4 leading `0` and difficulty is `4` `hashNonce` should be `true`")
	}

	// Test has nonce fail
	b.difficulty++
	n = b.hasNonce()

	if n {
		t.Error("if block hash has 4 leading `0` and difficulty is `5` `hashNonce` should be `false`")
	}
}

func TestCalcDifficulty(t *testing.T) {
	pb := &Block{
		timestamp:  1000,
		difficulty: 5,
	}

	cb := &Block{
		timestamp: pb.timestamp + mineRate,
	}

	cb.calcDifficulty(*pb)

	if cb.difficulty != pb.difficulty-1 {
		t.Errorf(`if time is equal to previous block plus mine rate '%v' difficulty should decrease, in order to mine block quicker. 
		- Previous Block difficulty %v 	: Previous Block timestamp : %v
		- Current Block difficulty %v 	: Current Block timestamp : %v`, mineRate, pb.difficulty, pb.timestamp, cb.difficulty, cb.timestamp)
	}

	cb.timestamp = pb.timestamp + mineRate + 1
	cb.calcDifficulty(*pb)

	if cb.difficulty != pb.difficulty-1 {
		t.Errorf(`if time is greater to previous block plus mine rate '%v' difficulty should decrease, in order to mine block quicker. 
		- Previous Block difficulty %v 	: Previous Block timestamp : %v
		- Current Block difficulty %v 	: Current Block timestamp : %v`, mineRate, pb.difficulty, pb.timestamp, cb.difficulty, cb.timestamp)
	}

	cb.timestamp = pb.timestamp + mineRate - 1
	cb.calcDifficulty(*pb)

	if cb.difficulty != pb.difficulty+1 {
		t.Errorf(`if time is less than previous block plus minerate '%v' difficulty should increase, in order to mine block slower. 
		- Previous Block difficulty %v 	: Previous Block timestamp : %v
		- Current Block difficulty %v 	: Current Block timestamp : %v`, mineRate, pb.difficulty, pb.timestamp, cb.difficulty, cb.timestamp)
	}

}
