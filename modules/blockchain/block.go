package blockchain

import (
	"fmt"
	"strings"

	"github.com/murryIsDeveloping/fernychain-go/modules/hashing"
	"github.com/murryIsDeveloping/fernychain-go/modules/util"
)

// Block represents a block within the blockchain
type Block struct {
	timestamp    int64
	hash         string
	value        []Transaction
	previousHash string
	nonce        int
	difficulty   int
}

func (b *Block) proofOfWork(prevBlock Block) *Block {
	for {
		b.timestamp = util.NowUnixMs()
		b.nonce++

		b.calcDifficulty(prevBlock)

		b.hash = hashBlock(*b)

		if b.hasNonce() {
			break
		}
	}

	return b
}

func validPreviousHash(currentBlock Block, prevBlock Block) bool {
	fmt.Printf("current block %v", currentBlock.value)
	pHash := hashBlock(prevBlock)
	if pHash != currentBlock.previousHash {
		return false
	}

	bHash := hashBlock(currentBlock)
	if bHash != currentBlock.hash {
		return false
	}

	return true
}

func (b *Block) calcDifficulty(prevBlock Block) {
	if prevBlock.timestamp+mineRate > b.timestamp {
		b.difficulty = prevBlock.difficulty + 1
	} else {
		b.difficulty = prevBlock.difficulty - 1
	}
}

func hashBlock(b Block) string {
	return hashing.Hash(fmt.Sprint(b.timestamp, b.value, b.previousHash, b.difficulty, b.nonce))
}

func (b *Block) hasNonce() bool {
	nonceNeeded := strings.Repeat("0", b.difficulty)
	currentNonce := b.hash[:b.difficulty]
	return nonceNeeded == currentNonce
}

// Print prints the value of the block
func (b *Block) Print() {
	fmt.Printf(`Block - hash: %v
	timestamp: %v
	value: %v
	nonce: %v
	difficulty: %v
	previousHash: %v

`, b.hash, b.timestamp, b.value, b.nonce, b.difficulty, b.previousHash)
}
