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

		// there's a 15% chance that a nested folder will be created inside
		if rand.Float32() < 0.15 {

			//random folder size is either the remaining size needed or between 10 and 100 mb
			randFolderSize := min(remaining, uint64(rand.Intn(9)+1)*10*1024)
			_, err = CreateRandomFolder(folderName, randFolderSize)
			if err != nil {
				return "", err
			}
			written += randFolderSize
			continue
		}

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
