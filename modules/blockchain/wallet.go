package blockchain

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	b64 "encoding/base64"
	"encoding/pem"
)

// Wallet represents the end users/miners crypto wallet
type Wallet struct {
	key   *rsa.PrivateKey
	value float64
}

// GenerateWallet generates a wallet for a user
func GenerateWallet(pemKey []byte) (*Wallet, error) {
	w := &Wallet{
		value: 0,
	}

	if pemKey != nil {
		key, err := x509.ParsePKCS1PrivateKey(pemKey)

		if err != nil {
			return nil, err
		}

		w.key = key
		return w, nil
	}

	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}

	w.key = key
	return w, nil
}

func privateKeyToBytes(priv *rsa.PrivateKey) []byte {
	privBytes := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(priv),
		},
	)

	return privBytes
}

func publicKeyToBytes(pub *rsa.PublicKey) []byte {
	pubASN1, err := x509.MarshalPKIXPublicKey(pub)
	if err != nil {
		panic(err)
	}

	pubBytes := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pubASN1,
	})

	return pubBytes
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

func (w *Wallet) PublicKey() string {
	bKey := publicKeyToBytes(&w.key.PublicKey)
	return b64.StdEncoding.EncodeToString(bKey)
}

func ToPem(key *rsa.PrivateKey) []byte {
	return x509.MarshalPKCS1PrivateKey(key)
}

func (w *Wallet) findValue(bc *Blockchain) float64 {
	key := w.PublicKey()
	return bc.GetAddressValue(key)
}
