package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"log"
	"testing"
	"time"
)

var secretKey = []byte("N1PCdw3M2B1TfJhoaY2mL736p2vCUc47") // 32-byte key for AES-256

func encryptAES(plaintext []byte) ([]byte, error) {
	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return nil, err
	}

	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)
	return ciphertext, nil
}

func decryptAES(ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

func TestEncryptionPerformance() {
	// Sample data
	data := []byte("This is a sample text for testing encryption performance." * 1000) // 1000 repetitions to increase length

	// Measure encryption time
	startEncrypt := time.Now()
	ciphertext, err := encryptAES(data)
	if err != nil {
		log.Fatal(err)
	}
	elapsedEncrypt := time.Since(startEncrypt)

	// Measure decryption time
	startDecrypt := time.Now()
	plaintext, err := decryptAES(ciphertext)
	if err != nil {
		log.Fatal(err)
	}
	elapsedDecrypt := time.Since(startDecrypt)

	// Output results
	fmt.Printf("Encryption time: %s\n", elapsedEncrypt)
	fmt.Printf("Decryption time: %s\n", elapsedDecrypt)

	// Ensure we successfully decrypted back to original plaintext
	if string(plaintext) != string(data) {
		log.Fatal("Decryption failed: output does not match input.")
	}
}

func main() {
	TestEncryptionPerformance()
}
