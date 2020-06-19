package blockchain

import (
	"sync"

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
	if rc.validChain() && len(rc.blocks) > len(bc.blocks) {
		bc.blocks = rc.blocks
	}
}

func monitorWorker(wg *sync.WaitGroup, ch chan bool) {
	wg.Wait()
	close(ch)
}

func batch(upperIndex int, lowerIndex int, bc *Blockchain) bool {
	wg := &sync.WaitGroup{}
	c := make(chan bool)

	for i := upperIndex; i > lowerIndex; i-- {
		wg.Add(1)
		go validPreviousHash(*bc.blocks[i], *bc.blocks[i-1], c, wg)
	}

	go monitorWorker(wg, c)

	for res := range c {
		if !res {
			return false
		}
	}

	return true
}

func calcLower(upperIndex int) int {
	if upperIndex-batchSize >= 0 {
		return upperIndex - batchSize
	}

	return 0
}

func (bc *Blockchain) validChain() bool {
	upperIndex := len(bc.blocks) - 1
	lowerIndex := calcLower(upperIndex)
	valid := true

	for {
		valid = batch(upperIndex, lowerIndex, bc)
		upperIndex = lowerIndex
		lowerIndex = calcLower(upperIndex)
		if upperIndex <= 0 || valid == false {
			break
		}
	}

	return valid
}

// PrintBlocks prints formatted blocks to console
func (bc *Blockchain) PrintBlocks() {
	for _, b := range bc.blocks {
		b.Print()
	}
}
