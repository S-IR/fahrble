package compress

import (
	"errors"
	"os"
	"path/filepath"
	"sync"
)

type Archive interface {
	AddFile(folderPath, filePath string, fileInfo os.FileInfo) error
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

func ArchiveFolder(folderPath string, config ArchiveConfig) ([]byte, error) {
	info, err := os.Stat(folderPath)
	if err != nil {
		return nil, err
	}
	if !info.IsDir() {
		return nil, errors.New("you must provide a folder")
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

	err = addFilesToArchive(archive, folderPath)
	if err != nil {
		return nil, err
	}

	return archive.WriteToMemory(), nil
}

// Function to add files to the ZIP archive
func addFilesToArchive(archive Archive, basePath string) error {
	var wg sync.WaitGroup

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
			if err := archive.AddFile(basePath, filePath, fileInfo); err != nil {
				panic(err)
			}

		}(filePath, fileInfo)

		return nil
	})
	wg.Wait()
	return err
}
