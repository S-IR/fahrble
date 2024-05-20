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
func (a *TarGzipArchive) AddFile(folderPath, filePath string, fileInfo os.FileInfo) error {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	relPath, err := filepath.Rel(folderPath, filePath)
	relPath = filepath.ToSlash(relPath)

	if relPath == "." {
		relPath = fileInfo.Name()
	}

	if err != nil {
		return err
	}

	header, err := tar.FileInfoHeader(fileInfo, relPath)
	if err != nil {
		return err
	}

	header.Name = relPath

	if err := a.tarWriter.WriteHeader(header); err != nil {
		return err
	}
	// Calculate SHA1 hash of the file contents
	shaHash := sha1.New()
	if _, err := io.Copy(shaHash, file); err != nil {
		return err
	}
	SHA := shaHash.Sum(nil)

	a.infos = append(a.infos, &info.File{
		Sha:          SHA,
		Name:         fileInfo.Name(),
		Size:         uint64(fileInfo.Size()),
		LastModified: timestamppb.New(fileInfo.ModTime()),
	})
	if fileInfo.IsDir() {
		return nil
	}
	// Reset the file offset to beginning
	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		return err
	}
	_, err = io.Copy(a.tarWriter, file)
	if err != nil {
		return err
	}

	return nil
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
