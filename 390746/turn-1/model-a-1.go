package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"io/ioutil"
	"log"
)

func generateKey() ([]byte, error) {
	key := make([]byte, 32) // AES-256 requires a 32-byte key
	_, err := rand.Read(key)
	if err != nil {
		return nil, err
	}
	return key, nil
}

func encrypt(plaintext []byte, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := rand.Read(iv); err != nil {
		return "", err
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)
	return hex.EncodeToString(ciphertext), nil
}

func main() {
	key, err := generateKey()
	if err != nil {
		log.Fatal(err)
	}

	// Example of plaintext data
	plaintext := []byte("Sensitive user data")

	encrypted, err := encrypt(plaintext, key)
	if err != nil {
		log.Fatal(err)
	}

	// Save 'encrypted' data to the database
	log.Printf("Encrypted Data: %s", encrypted)
	// Save 'key' securely (e.g., environment variable, secure vault, etc.)
}
