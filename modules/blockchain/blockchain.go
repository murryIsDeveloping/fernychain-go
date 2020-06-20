package blockchain

import (
	"github.com/murryIsDeveloping/fernychain-go/modules/util"
)

const mineRate = 1000
const startingDifficulty = 3
const batchSize = 256

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
	}

	g.hash = hashBlock(*g)

	bc.blocks = append(bc.blocks, g)
	return bc
}

// GetBlock gets a read only value of the block
func (bc *Blockchain) GetBlock(index int) Block {
	return *bc.blocks[index]
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
	nb.Print()
	return nb
}

// ReplaceChain will replace current chains blocks with input blockchain blocks
// if all blocks are valid and length of the new chain is greater than the current chain
func (bc *Blockchain) ReplaceChain(rc *Blockchain) {
	if rc.validate() && len(rc.blocks) > len(bc.blocks) {
		bc.blocks = rc.blocks
	}
}

func validationWorker(currentBlock Block, prevBlock Block, i int, c chan bool, done chan bool) {
	for {
		switch {
		case <-c:
			value := validPreviousHash(currentBlock, prevBlock)

			if !value {
				done <- false
				return
			}

			if i == 1 {
				done <- true
				return
			}

			c <- value
		case <-done:
			return
		}
	}

}

func (bc *Blockchain) validate() bool {
	c := make(chan bool, batchSize)
	done := make(chan bool)

	for i := len(bc.blocks) - 1; i > 0; i-- {
		go validationWorker(*bc.blocks[i], *bc.blocks[i-1], i, c, done)
	}

	for i := 0; i < batchSize; i++ {
		c <- true
	}

	return <-done
}

// PrintBlocks prints formatted blocks to console
func (bc *Blockchain) PrintBlocks() {
	for _, b := range bc.blocks {
		b.Print()
	}
}
