package compress

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"crypto/sha1"
	"io"
	"os"
	"path/filepath"

	"github.com/S-IR/freible/info"
	"github.com/s-ir/fahrble/node/ledger"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type TarGzipArchive struct {
	buffer     *bytes.Buffer
	gzipWriter *gzip.Writer

	tarWriter *tar.Writer
	infos     []*info.File
}

func NewGzipTarArchive() *TarGzipArchive {
	buffer := new(bytes.Buffer)
	gzipWriter := gzip.NewWriter(buffer)
	tarWriter := tar.NewWriter(gzipWriter)

	return &TarGzipArchive{
		buffer:     buffer,
		gzipWriter: gzipWriter,
		tarWriter:  tarWriter,
		infos:      []*info.File{},
	}
}
func (a *TarGzipArchive) AddFile(folderPath, filePath string, fileInfo os.FileInfo) (*ledger.Ledger, error) {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	relPath, err := filepath.Rel(folderPath, filePath)
	relPath = filepath.ToSlash(relPath)

	if relPath == "." {
		relPath = fileInfo.Name()
	}

	if err != nil {
		return nil, err
	}

	header, err := tar.FileInfoHeader(fileInfo, relPath)
	if err != nil {
		return nil, err
	}

	header.Name = relPath

	if err := a.tarWriter.WriteHeader(header); err != nil {
		return nil, err
	}
	// Calculate SHA1 hash of the file contents
	shaHash := sha1.New()
	if _, err := io.Copy(shaHash, file); err != nil {
		return nil, err
	}
	SHA := shaHash.Sum(nil)

	a.infos = append(a.infos, &info.File{
		Sha:          SHA,
		Name:         fileInfo.Name(),
		Size:         uint64(fileInfo.Size()),
		LastModified: timestamppb.New(fileInfo.ModTime()),
	})
	if fileInfo.IsDir() {
		return nil, nil
	}
	// Reset the file offset to beginning
	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(a.tarWriter, file)
	if err != nil {
		return nil, err
	}

	return ledger.GenerateLedger(*file, fileInfo), nil
}
func (a *TarGzipArchive) WriteToMemory() []byte {
	if err := a.tarWriter.Close(); err != nil {
		panic(err)
	}

	return a.buffer.Bytes()
}

func (a *TarGzipArchive) WriteTo(path string) error {

	if err := a.tarWriter.Close(); err != nil {
		return err
	}

	if err := a.gzipWriter.Close(); err != nil {
		return err
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}

	_, err = file.Write(a.buffer.Bytes())
	if err != nil {
		return err
	}
	return nil
}
