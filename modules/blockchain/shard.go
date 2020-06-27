package blockchain

// Shard represents a small chunk of the blockchain
type Shard struct {
	id          string
	lastHash    string
	blocks      []*Block
	prevShardID string
}
