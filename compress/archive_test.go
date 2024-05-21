package compress_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/klauspost/compress/zip"

	"github.com/S-IR/freible/compress"
)

func TestZip(t *testing.T) {
	const DIR_PATH = "../mockup_files"
	const OUTPUT_PATH = "OUTPUT.zip"
	os.Remove(OUTPUT_PATH)

	err := compress.ArchiveFolder(DIR_PATH, OUTPUT_PATH, compress.ArchiveConfig{
		ArchiveType: compress.ZipType,
	})
	if err != nil {
		t.Error(err)
	}

	dirFiles := make(map[string]bool)
	err = filepath.Walk(DIR_PATH, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			relPath, err := filepath.Rel(DIR_PATH, path)
			if err != nil {
				return err
			}
			relPath = filepath.ToSlash(relPath)
			dirFiles[relPath] = true
		}
		return nil
	})

	if err != nil {
		t.Error(err)
	}

	zipFile, err := zip.OpenReader(OUTPUT_PATH)
	if err != nil {
		t.Error(err)
	}
	defer zipFile.Close()

	zipFiles := make(map[string]bool)
	// Iterate through each file in the zip and populate the map
	for _, file := range zipFile.File {
		zipFiles[file.Name] = true
	}

	for file := range dirFiles {
		dirFilePath := filepath.ToSlash(file)
		if _, ok := zipFiles[dirFilePath]; !ok {
			t.Fatalf("file %s is missing from the zip \n", dirFilePath)
			return
		}
	}

}

func TestGzipTar(t *testing.T) {
	err := compress.ArchiveFolder("../mockup_files", "test.tar.gz", compress.ArchiveConfig{
		ArchiveType: compress.TarGzipType,
	})
	if err != nil {
		panic(err)
	}
	os.RemoveAll("test.tar.gz")

}
func BenchmarkZip(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := compress.ArchiveFolder("../mockup_files", "test.zip", compress.ArchiveConfig{
			ArchiveType: compress.ZipType,
		})
		if err != nil {
			panic(err)
		}
	}
	os.RemoveAll("test.zip")
}
