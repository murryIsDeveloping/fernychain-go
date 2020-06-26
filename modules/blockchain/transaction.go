package blockchain

import (
	"errors"
)

type Output struct {
	amount  float64
	address string
}

type Input struct {
	address   string
	value     float64
	signature string
}

type Transaction struct {
	outputs []Output
	input   Input
}

func (w *Wallet) CreateTransaction() *Transaction {
	input := Input{
		address: w.PublicKey(),
		value:   w.value,
	}

	t := &Transaction{
		outputs: []Output{},
		input:   input,
	}

	return t
}

func (trans *Transaction) addTransaction(amount float64, address string) (*Transaction, error) {
	o := Output{
		amount:  amount,
		address: address,
	}

	if trans.input.value-amount < 0.0 {
		return trans, errors.New("Funds not available to complete transaction")
	}

	trans.input.value -= amount

	if len(trans.outputs) > 0 {
		trans.outputs = append(trans.outputs, o)
	} else {
		trans.outputs = []Output{o}
	}

	return trans, nil
}
