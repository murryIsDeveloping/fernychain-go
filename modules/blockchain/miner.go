package blockchain

// Miner represents the miner of the blockchain
type Miner struct {
	wallet          *Wallet
	currentShard    *Shard
	shards          []*Shard
	transactionPool TransactionPool
}

// FindShard Checks if miner has shard of that ID
func (m *Miner) FindShard(id string) *Shard {
	for i := 0; i < len(m.shards); i++ {
		if m.shards[i].id == id {
			return m.shards[i]
		}
	}

	return nil
}
