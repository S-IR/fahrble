package encrypt

import (
	"bytes"
	"crypto/rand"
	"testing"
)

func TestEncryptDecryptBytes(t *testing.T) {
	// Original plaintext
	plaintext := []byte("This is a test plaintext.")

	// Generate a random 32-byte key for AES-256
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		t.Fatalf("Failed to generate key: %v", err)
	}

	// Encrypt the plaintext
	ciphertext, err := EncryptBytes(plaintext, key)
	if err != nil {
		t.Fatalf("Encryption failed: %v", err)
	}

	// Decrypt the ciphertext
	decryptedText, err := DecryptBytes(ciphertext, key)
	if err != nil {
		t.Fatalf("Decryption failed: %v", err)
	}

	// Check if the decrypted text matches the original plaintext
	if !bytes.Equal(plaintext, decryptedText) {
		t.Fatalf("Decrypted text does not match original plaintext.\nOriginal: %s\nDecrypted: %s", plaintext, decryptedText)
	}
}
