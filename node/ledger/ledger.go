package ledger

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/s-ir/fahrble/lib"
	"google.golang.org/protobuf/proto"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func GenerateLedger(file os.File, info os.FileInfo) *Ledger {
	fileLedgerInfo := Ledger{
		Sha:          []byte(lib.GetShaHash(&file)),
		Name:         filepath.Base(filepath.Base(file.Name())),
		LastModified: timestamppb.New(info.ModTime()),
		Size:         uint64(info.Size() / 1024), // size in KB
	}
	return &fileLedgerInfo
}
func StoreBackup(backup *BackupSchema, folderPath string) error {

	fileName := fmt.Sprintf("%s-%s", backup.Name, backup.LastModified)

	fullPath := fmt.Sprintf("%s/%s", folderPath, fileName)
	file, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer file.Close()

	backupBytes, err := proto.Marshal(backup)
	if err != nil {
		return err
	}

	_, err = file.Write(backupBytes)
	if err != nil {
		return err
	}
	return nil

}
