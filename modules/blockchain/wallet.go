package blockchain

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	b64 "encoding/base64"
	"encoding/pem"
	"fmt"
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

func privateKeyToB64(priv *rsa.PrivateKey) string {
	privBytes := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(priv),
		},
	)

	return b64.StdEncoding.EncodeToString(privBytes)
}

func b64ToPrivateKey(address string) (*rsa.PrivateKey, error) {
	bytes, err := b64.StdEncoding.DecodeString(address)
	if err != nil {
		return nil, err
	}

	return x509.ParsePKCS1PrivateKey(bytes)
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

func b64ToPublicKey(address string) (*rsa.PublicKey, error) {
	bytes, err := b64.StdEncoding.DecodeString(address)
	if err != nil {
		return nil, err
	}
	return x509.ParsePKCS1PublicKey(bytes)
}

// SignTransaction signs the transaction so it is clear transaction was signed by original sender
func (w *Wallet) SignTransaction(transaction *Transaction) string {
	msgHash := sha256.New()
	_, err := msgHash.Write([]byte(fmt.Sprintf("%v", transaction.outputs)))
	if err != nil {
		panic(err)
	}

	msgHashSum := msgHash.Sum(nil)

	signature, err := rsa.SignPSS(rand.Reader, w.key, crypto.SHA256, msgHashSum, nil)
	if err != nil {
		panic(err)
	}

	return b64.StdEncoding.EncodeToString(signature)
}

// TransactionValid Allows others to check if the signature is valid
func TransactionValid(publicKey string, signature []byte, msgHashSum []byte) bool {
	// decode public key from bytes
	key, err := b64ToPublicKey(publicKey)
	if err != nil {
		return false
	}

	err = rsa.VerifyPSS(key, crypto.SHA256, msgHashSum, signature, nil)
	if err != nil {
		return false
	}
	return true
}

// PublicKey Returns the public key of the wallet as a string
func (w *Wallet) PublicKey() string {
	bKey := publicKeyToBytes(&w.key.PublicKey)
	return b64.StdEncoding.EncodeToString(bKey)
}

// ToPem converts the private key to a pem
func ToPem(key *rsa.PrivateKey) []byte {
	return x509.MarshalPKCS1PrivateKey(key)
}

func (w *Wallet) findValue(bc *Blockchain) float64 {
	key := w.PublicKey()
	return bc.GetAddressValue(key)
}
