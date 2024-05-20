package lib

import (
	"hash"
	"io"
	"log"
	"os"
)

func GetShaHash(file *os.File, shaHash hash.Hash) []byte {

	if _, err := io.Copy(shaHash, file); err != nil {
		log.Fatalf("failed to copy file contents to hash: %v", err)
	}

	SHA := shaHash.Sum(nil)
	return SHA
}
