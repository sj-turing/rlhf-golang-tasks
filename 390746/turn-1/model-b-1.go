package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
)

func encryptAES(key []byte, plaintext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	ciphertext := gcm.Seal(nil, nonce, plaintext, nil)
	return append(nonce, ciphertext...), nil
}

func decryptAES(key []byte, ciphertext []byte) ([]byte, error) {
	if len(ciphertext) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:aes.BlockSize], ciphertext[aes.BlockSize:]
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

func main() {
	key := []byte("N1PCdw3M2B1TfJhoaY2mL736p2vCUc47") // Replace with a 32-byte key
	plaintext := []byte("This is a secret message.")

	ciphertext, err := encryptAES(key, plaintext)
	if err != nil {
		fmt.Println("Error encrypting:", err)
		return
	}

	encryptedBase64 := base64.StdEncoding.EncodeToString(ciphertext)
	fmt.Println("Encrypted data:", encryptedBase64)

	decryptedBytes, err := decryptAES(key, ciphertext)
	if err != nil {
		fmt.Println("Error decrypting:", err)
		return
	}

	fmt.Println("Decrypted data:", string(decryptedBytes))
}
