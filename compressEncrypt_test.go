package main

import (
	"bytes"
	"crypto/rand"
	"os"
	"testing"

	"github.com/s-ir/fahrble/compress"
	"github.com/s-ir/fahrble/compress/generateMockups"
	"github.com/s-ir/fahrble/encrypt"
)

func TestCompressAndEncrypt(t *testing.T) {

	folderName, err := generateMockups.CreateRandomFolder("./TEST_MOCKUP_FOLDER", 100*1024)
	if err != nil {
		t.Fatalf("Compression failed: %v", err)
	}

	defer os.RemoveAll(folderName)

	// Compress the folder
	zipBytes, err := compress.ArchiveFolder(folderName, compress.ArchiveConfig{
		ArchiveType: compress.ZipType,
	})
	if err != nil {
		t.Fatalf("Compression failed: %v", err)
	}

	// Generate a random 32-byte key for AES-256
	key := make([]byte, 32)
	_, err = rand.Read(key)
	if err != nil {
		t.Fatalf("Failed to generate key: %v", err)
	}

	// Encrypt the compressed bytes
	encryptedBytes, err := encrypt.EncryptBytes(zipBytes, key)
	if err != nil {
		t.Fatalf("Encryption failed: %v", err)
	}

	// Decrypt the encrypted bytes
	decryptedBytes, err := encrypt.DecryptBytes(encryptedBytes, key)
	if err != nil {
		t.Fatalf("Decryption failed: %v", err)
	}

	// Check if the decrypted bytes match the original compressed bytes
	if !bytes.Equal(zipBytes, decryptedBytes) {
		t.Fatalf("Decrypted bytes do not match original compressed bytes.\nOriginal: %x\nDecrypted: %x", zipBytes, decryptedBytes)
	}
}
