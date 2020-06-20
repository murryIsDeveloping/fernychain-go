package wallet

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
)

// Wallet represents the users virtual wallet it is responsible for signing their transactions and receiving transaction
type Wallet struct {
	privateKey rsa.PrivateKey
	publicKey  rsa.PublicKey
}

// GenerateWallet generates a wallet for a user
func GenerateWallet() {
	w := Wallet{}
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	publicKey := privateKey.PublicKey

	w.privateKey = *privateKey
	w.publicKey = publicKey

	fmt.Println("Private Key : ", privateKey)
	fmt.Println("Public key ", publicKey)

}
