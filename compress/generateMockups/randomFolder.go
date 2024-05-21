package generateMockups

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/S-IR/freible/lib"
)

// Generates a folder of  SIZE (in kb) size with random mockup files of random sizes between 1 and 10 mb
// Name is always temp-{} , {} being a random string of len 6
func CreateRandomFolder(path string, size uint64) (folderName string, err error) {
	folderName = fmt.Sprintf("%s/temp-%s", path, lib.RandomString(6))
	err = os.MkdirAll(folderName, 0777)
	if err != nil {
		return "", err
	}

	for written := uint64(0); written < size; {
		remaining := size - written

		//this is between 1 and 10 mb

		randFileSize := uint64((rand.Intn(9) + 1) * 1024)

		if remaining < randFileSize {
			CreateRandomFile(folderName, remaining)
			written += remaining
		} else {
			CreateRandomFile(folderName, randFileSize)
			written += randFileSize
		}
	}
	return folderName, nil
}
