package main

import (
	"os"

	"github.com/S-IR/freible/compress/generateMockups"
	"github.com/s-ir/fahrble/encrypt"
	"github.com/s-ir/fahrble/node/compress"
	"github.com/s-ir/fahrble/node/ledger"
)

func main() {
	const ZIP_FILE_PATH = "assets/zipFile.zip"
	const ENCRYPTED_FILE_PATH = "assets/encryptedZip.bin"
	const DECRYPTED_FILE_PATH = "assets/decryptedZipFile.zip"

	mockupFolderName, err := generateMockups.CreateRandomFolder("assets/TEST_MOCKUP_FOLDER", 1024*100)
	if err != nil {
		panic(err)
	}

	zipBytes, backup, err := compress.ArchiveFolder(mockupFolderName, compress.ArchiveConfig{
		ArchiveType: compress.ZipType,
	})

	ledger.StoreBackup(backup, "assets")

	if err != nil {
		panic(err)
	}

	err = os.WriteFile(ZIP_FILE_PATH, zipBytes, 0777)
	if err != nil {
		panic(err)
	}

	key := []byte("very secretive key brozzer")
	encryptedZipBytes, err := encrypt.EncryptBytes(zipBytes, key)

	if err != nil {
		panic(err)
	}

	err = os.WriteFile(ENCRYPTED_FILE_PATH, encryptedZipBytes, 0777)
	if err != nil {
		panic(err)
	}

	cipherText, err := os.ReadFile(ENCRYPTED_FILE_PATH)
	if err != nil {
		panic(err)
	}

	decryptedBytes, err := encrypt.DecryptBytes(cipherText, key)

	if err != nil {
		panic(err)
	}

	err = os.WriteFile(DECRYPTED_FILE_PATH, decryptedBytes, 0777)
	if err != nil {
		panic(err)
	}

}
