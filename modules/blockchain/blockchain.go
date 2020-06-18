package blockchain

import (
	"fmt"
	"strings"

	"github.com/murryIsDeveloping/fernychain-go/modules/hashing"
	"github.com/murryIsDeveloping/fernychain-go/modules/util"
)

const mineRate = 1000
const startingDifficulty = 3

// Blockchain represents the chain of blocks
type Blockchain struct {
	blocks []*Block
}

// Genisis creates an instance of the blockchain containing a Genisis Block
func Genisis() *Blockchain {
	bc := &Blockchain{
		blocks: []*Block{},
	}

	g := &Block{
		timestamp:  util.NowUnixMs(),
		difficulty: startingDifficulty,
		hash:       hashing.Hash("Genisis"),
	}

	bc.blocks = append(bc.blocks, g)
	return bc
}

// MineBlock mines a block by generating a new block and adding it to the blockchain
func (bc *Blockchain) MineBlock(value string) *Block {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	nonce := 0
	b := &Block{
		timestamp:    util.NowUnixMs(),
		previousHash: prevBlock.hash,
		value:        value,
		difficulty:   prevBlock.difficulty,
		nonce:        nonce,
	}

	nb := b.proofOfWork(*prevBlock)
	bc.blocks = append(bc.blocks, nb)
	nb.print()
	return nb
}

// PrintBlocks prints formatted blocks to console
func (bc *Blockchain) PrintBlocks() {
	for _, b := range bc.blocks {
		b.print()
	}
}

// Block represents a block within the blockchain
type Block struct {
	timestamp    int64
	hash         string
	value        string
	previousHash string
	nonce        int
	difficulty   int
}

func (b *Block) proofOfWork(prevBlock Block) *Block {
	for {
		b.timestamp = util.NowUnixMs()
		b.nonce++

		if prevBlock.timestamp+mineRate >= b.timestamp {
			b.difficulty = prevBlock.difficulty + 1
		} else {
			b.difficulty = prevBlock.difficulty - 1
		}

		b.hashBlock()

		if b.hasNonce() {
			fmt.Printf("Increase %v \n %v >= %v \n", prevBlock.timestamp+mineRate >= b.timestamp, prevBlock.timestamp+mineRate, b.timestamp)
			break
		}
	}

	fmt.Printf("hash %v block %v", b.hash, b)
	return b
}

func (b *Block) hashBlock() *Block {
	h := hashing.Hash(fmt.Sprint(b.timestamp, b.value, b.previousHash, b.difficulty, b.nonce))
	b.hash = h
	return b
}

func (b *Block) hasNonce() bool {
	nonceNeeded := strings.Repeat("0", b.difficulty)
	currentNonce := b.hash[:b.difficulty]
	return nonceNeeded == currentNonce
}

func (b *Block) print() {
	fmt.Printf(`Block - hash: %v
	timestamp: %v
	value: %v
	nonce: %v
	difficulty: %v
	previousHash: %v

`, b.hash, b.timestamp, b.value, b.nonce, b.difficulty, b.previousHash)
}
