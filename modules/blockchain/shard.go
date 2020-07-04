package blockchain

const shardSize = 10

// Shard represents a small chunk of the blockchain
type Shard struct {
	id          string
	lastHash    string
	blocks      []*Block
	prevShardID string
}

// func CreateShardMap() *map[string][]int {
// 	return &make(map[string][]int)
// }

// func RemovePeerFromShardMap(m *map[string][]int, peer string) {
// 	delete(m, peer)
// }

// func AddPeerToShardMap(m *map[string][]int, peer string, shards []int) {
// 	m[peer] = shards
// }
