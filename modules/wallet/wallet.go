package wallet

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
)

// GenerateWallet generates a wallet for a user
func GenerateWallet(pemKey []byte) (*rsa.PrivateKey, error) {
	if pemKey != nil {
		key, err := x509.ParsePKCS1PrivateKey(pemKey)

		if err != nil {
			return nil, err
		}

		return key, nil
	}

	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}

	return key, nil
}

func SignTransaction(wallet *rsa.PrivateKey, transaction string) []byte {
	msgHash := sha256.New()
	_, err := msgHash.Write([]byte(transaction))
	if err != nil {
		panic(err)
	}

	msgHashSum := msgHash.Sum(nil)

	signature, err := rsa.SignPSS(rand.Reader, wallet, crypto.SHA256, msgHashSum, nil)
	if err != nil {
		panic(err)
	}

	return signature
}

func TransactionValid(publicKey *rsa.PublicKey, signature []byte, msgHashSum []byte) bool {
	err := rsa.VerifyPSS(publicKey, crypto.SHA256, msgHashSum, signature, nil)
	if err != nil {
		return false
	}
	return true
}

func ToPem(key *rsa.PrivateKey) []byte {
	return x509.MarshalPKCS1PrivateKey(key)
}
