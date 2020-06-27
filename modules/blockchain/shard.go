package blockchain

// Shard represents a small chunk of the blockchain
type Shard struct {
	lastHash string
	blocks   []*Block
	blockID  int
}
