package lib

import (
	"crypto/sha256"
	"io"
	"log"
	"os"
)

func GetShaHash(file *os.File) []byte {
	shaHash := sha256.New()
	if _, err := io.Copy(shaHash, file); err != nil {
		log.Fatalf("failed to copy file contents to hash: %v", err)
	}

	SHA := shaHash.Sum(nil)
	return SHA
}
