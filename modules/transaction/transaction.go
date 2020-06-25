package transaction

import "crypto/rsa"

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

func (trans *Transaction) addTransaction(wallet *rsa.PrivateKey, amount float64, address string) {
	o := Output{
		amount:  amount,
		address: address,
	}
	if len(trans.outputs) > 0 {
		trans.outputs = append(trans.outputs, o)
	} else {
		trans.outputs = []Output{o}
	}
}
