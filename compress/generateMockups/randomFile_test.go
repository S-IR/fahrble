package generateMockups_test

import (
	"fmt"
	"os"
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
	if fileInfo.Size() == 1024*int64(FILE_SIZE) {
		t.Fatalf("generated random file is not of the specified size \n")
	}

}

func TestCreateRandomFolder(t *testing.T) {
	const FOLDER_SIZE = uint64(1024 * 10)

	folderName, err := generateMockups.CreateRandomFolder(".", FOLDER_SIZE)
	defer os.RemoveAll(folderName)

	if err != nil {
		t.Error(err)
	}

	folder, err := os.Open(folderName)
	if err != nil {
		t.Error(err)
	}
	defer folder.Close()

	folderInfo, err := folder.Stat()
	if err != nil {
		t.Error(err)
	}
	if folderInfo.Size() == 1024*int64(FOLDER_SIZE) {
		t.Fatalf("generated random folder is not of the specified size \n")
	}
}
