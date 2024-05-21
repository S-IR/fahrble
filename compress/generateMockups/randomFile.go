package generateMockups

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/S-IR/freible/lib"
)

// CreateRandomFile generates a file with random content of the specified IN KB
func CreateRandomFile(path string, size uint64) (fileName string, err error) {
	// Convert size from kilobytes to bytes
	sizeInBytes := size * 1024

	extensions := []string{".txt", ".png", ".svg", ".json", ".xml"}
	extension := extensions[rand.Intn(len(extensions))]
	fileName = fmt.Sprintf("%s/%s%s", path, lib.RandomString(10), extension)

	file, err := os.Create(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	const bufferSize = 1024
	buffer := make([]byte, bufferSize)
	for written := uint64(0); written < sizeInBytes; {
		remaining := sizeInBytes - written
		if remaining < bufferSize {
			for i := 0; i < int(remaining); i++ {
				buffer[i] = byte(rand.Intn(256))
			}
			_, err = file.Write(buffer[:remaining])
		} else {
			for i := 0; i < bufferSize; i++ {
				buffer[i] = byte(rand.Intn(256))
			}
			_, err = file.Write(buffer)
		}
		if err != nil {
			return "", err
		}
		written += uint64(len(buffer))
	}

	return fileName, err
}
