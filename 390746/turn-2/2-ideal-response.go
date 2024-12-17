package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"time"
)

var (
	key = []byte("N1PCdw3M2B1TfJhoaY2mL736p2vCUc47") // 32-byte AES key
)

func encrypt(plaintext []byte) []byte {
	aesBlock, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	gcm, err := cipher.NewGCM(aesBlock)
	if err != nil {
		panic(err)
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err)
	}

	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)
	return ciphertext
}

func decrypt(ciphertext []byte) []byte {
	aesBlock, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	gcm, err := cipher.NewGCM(aesBlock)
	if err != nil {
		panic(err)
	}

	if len(ciphertext) < gcm.NonceSize() {
		panic("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:gcm.NonceSize()], ciphertext[gcm.NonceSize():]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err)
	}

	return plaintext
}

func main() {
	data := make([]byte, 1024*1024) // 1MB of data
	for i := range data {
		data[i] = byte(i % 256)
	}

	// Encrypt and measure time
	start := time.Now()
	encryptedData := encrypt(data)
	encryptTime := time.Since(start)
	fmt.Println("Encryption time:", encryptTime)

	// Decrypt and measure time
	start = time.Now()
	decrypt(encryptedData)
	decryptTime := time.Since(start)
	fmt.Println("Decryption time:", decryptTime)
}
