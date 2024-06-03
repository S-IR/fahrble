package compress

import (
	"errors"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"

	"github.com/s-ir/fahrble/node/ledger"
)

type Archive interface {
	AddFile(folderPath, filePath string, fileInfo os.FileInfo) (*ledger.Ledger, error)
	WriteTo(outputPath string) error
	WriteToMemory() []byte
}

// ArchiveType represents the type of archive format.
type ArchiveType int

const (
	ZipType ArchiveType = iota
	TarGzipType
)

type ArchiveConfig struct {
	ArchiveType      ArchiveType
	CompressionLevel uint
}

func ArchiveFolder(folderPath string, config ArchiveConfig) ([]byte, *ledger.BackupSchema, error) {
	info, err := os.Stat(folderPath)
	if err != nil {
		return nil, nil, err
	}
	if !info.IsDir() {
		return nil, nil, errors.New("you must provide a folder")
	}

	var archive Archive
	switch config.ArchiveType {
	case TarGzipType:
		archive = NewGzipTarArchive()
	case ZipType:
		archive = NewZipArchive()
	default:
		archive = NewZipArchive()
	}

	schema, err := addFilesToArchive(archive, folderPath)
	if err != nil {
		return nil, nil, err
	}

	return archive.WriteToMemory(), schema, nil
}

// Function to add files to the ZIP archive
func addFilesToArchive(archive Archive, basePath string) (*ledger.BackupSchema, error) {
	var wg sync.WaitGroup

	backupSchema := ledger.BackupSchema{}
	ledgers := backupSchema.Ledgers
	// Walk the folder and add files to the archive
	err := filepath.Walk(basePath, func(filePath string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if fileInfo.IsDir() {
			return nil // Skip directories
		}

		wg.Add(1)
		go func(filePath string, fileInfo os.FileInfo) {
			defer wg.Done()
			ledger, err := archive.AddFile(basePath, filePath, fileInfo)
			if err != nil {
				panic(err)
			}
			ledgers = append(ledgers, ledger)

		}(filePath, fileInfo)

		return nil
	})
	wg.Wait()
	sort.Slice(backupSchema.Ledgers, func(i, j int) bool {
		return strings.Compare(backupSchema.Ledgers[i].Name, backupSchema.Ledgers[j].Name) < 0
	})
	return &backupSchema, err
}
