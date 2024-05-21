package generateMockups_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/S-IR/freible/compress/generateMockups"
	"github.com/S-IR/freible/lib"
)

func TestCreateRandomFile(t *testing.T) {

	const FILE_SIZE = uint64(1024)
	tempFolderName := fmt.Sprintf("./temp-%s", lib.RandomString(6))
	os.Mkdir(tempFolderName, 0777)
	defer os.RemoveAll(tempFolderName)

	fileName, err := generateMockups.CreateRandomFile(tempFolderName, FILE_SIZE)
	if err != nil {
		t.Error(err)
	}
	file, err := os.Open(fmt.Sprintf("%s/%s", tempFolderName, fileName))
	if err != nil {
		t.Error(err)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		t.Error(err)
	}
	if fileInfo.Size() != 1024*int64(FILE_SIZE) {
		t.Fatalf("generated random file is not of the specified size, the file size is %d kb while the requested file size was %d kb \n", fileInfo.Size()*1024, FILE_SIZE)
	}

}

func TestCreateRandomFolder(t *testing.T) {
	//100 mb size
	const FOLDER_SIZE = uint64(1024 * 100)

	folderName, err := generateMockups.CreateRandomFolder(".", FOLDER_SIZE)

	if err != nil {
		t.Error(err)
	}
	folderSize, err := dirSize(folderName)
	if err != nil {
		t.Error(err)
	}

	if folderSize != 1024*int64(FOLDER_SIZE) {
		t.Fatalf("generated random folder is not of the specified size, the folder size is %d while the requested file size was %d kb \n", folderSize, FOLDER_SIZE)
	}
	err = os.RemoveAll(folderName)

	if err != nil {
		t.Error(err)
	}
}
func dirSize(path string) (int64, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	return size, err
}
