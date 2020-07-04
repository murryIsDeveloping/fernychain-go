package blockchain

import (
	"errors"
)

// Output represents all the transactions made by the input for this block
type Output struct {
	amount  float64
	address string
}

// Input holds all values regarding the wallet that made the transactions
type Input struct {
	address   string
	value     float64
	signature string
}

// Transaction Holds all transaction data made by a wallet
type Transaction struct {
	outputs []Output
	input   Input
}

// CreateTransaction creates a transaction object an empty outputs and sets up the inputs
func (w *Wallet) CreateTransaction(bc *Blockchain) *Transaction {
	input := Input{
		address: w.PublicKey(),
		value:   w.findValue(bc),
	}

	t := &Transaction{
		outputs: []Output{},
		input:   input,
	}

	return t
}

// AddTransaction Adds a transaction to the transaction object
func (trans *Transaction) AddTransaction(w *Wallet, amount float64, address string) error {
	o := Output{
		amount:  amount,
		address: address,
	}

	if amount <= 0.0 {
		return errors.New("Transaction must be greater than zero")
	}

	if trans.input.value-amount < 0.0 {
		return errors.New("Funds not available to complete transaction")
	}

	trans.input.value -= amount

	trans.outputs = append(trans.outputs, o)

	trans.input.signature = w.SignTransaction(trans)

	return nil
}
