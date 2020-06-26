package blockchain

// TransactionPool Current pool of unconfimed transaction
type TransactionPool struct {
	transactions []*Transaction
}

// CreateTransactionPool Creates an empty Transaction Pool
func CreateTransactionPool() *TransactionPool {
	return &TransactionPool{
		transactions: []*Transaction{},
	}
}

// PushTransaction pushes a transaction to the transaction pool
func (tp *TransactionPool) PushTransaction(t *Transaction) {
	tp.transactions = append(tp.transactions, t)
}
