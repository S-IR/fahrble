package compress

import (
	"bytes"
	"crypto/sha1"
	"io"
	"os"
	"path/filepath"
	"sync"

	"github.com/S-IR/freible/info"
	"github.com/klauspost/compress/zip"
	"github.com/s-ir/fahrble/node/ledger"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ZipConfig struct {
	DisableStore   bool
	DisableDeflate bool
}
type ZipArchive struct {
	buffer *bytes.Buffer
	writer *zip.Writer

	writerMutex *sync.Mutex
	infos       []*info.File
}

func NewZipArchive(configuration ...*ZipConfig) *ZipArchive {

	var cfg *ZipConfig

	if len(configuration) > 0 {
		cfg = configuration[0]
	} else {
		cfg = &ZipConfig{}
	}

	buffer := new(bytes.Buffer)
	writer := zip.NewWriter(buffer)

	if cfg.DisableDeflate {
		writer.RegisterCompressor(zip.Deflate, nil)
	}
	if cfg.DisableStore {
		writer.RegisterCompressor(zip.Store, nil)
	}
	return &ZipArchive{
		buffer:      buffer,
		writer:      writer,
		writerMutex: &sync.Mutex{},
		infos:       []*info.File{},
	}
}

func (a *ZipArchive) WriteToMemory() []byte {
	if err := a.writer.Close(); err != nil {
		panic(err)
	}

	return a.buffer.Bytes()
}

func (a *ZipArchive) AddFile(folderPath, filePath string, fileInfo os.FileInfo) (*ledger.Ledger, error) {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	a.writerMutex.Lock()
	defer a.writerMutex.Unlock()

	header, err := zip.FileInfoHeader(fileInfo)
	if err != nil {
		return nil, err
	}

	relPath, err := filepath.Rel(folderPath, filePath)
	if err != nil {
		return nil, err
	}
	header.Name = filepath.ToSlash(relPath)
	header.Method = zip.Deflate

	// Create a writer for the file in the ZIP archive
	writer, err := a.writer.CreateHeader(header)
	if err != nil {
		return nil, err
	}

	// Calculate SHA1 hash of the file contents
	shaHash := sha1.New()
	if _, err := io.Copy(shaHash, file); err != nil {
		return nil, err
	}
	SHA := shaHash.Sum(nil)

	// Create a File struct for the backup
	a.infos = append(a.infos, &info.File{
		Sha:          SHA,
		Name:         fileInfo.Name(),
		Size:         uint64(fileInfo.Size()),
		LastModified: timestamppb.New(fileInfo.ModTime()),
	})

	// Reset the file offset to beginning
	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		return nil, err
	}

	// Copy the file's content to the ZIP archive
	_, err = io.Copy(writer, file)
	if err != nil {
		return nil, err
	}
	return ledger.GenerateLedger(*file, fileInfo), nil
}
func (a *ZipArchive) WriteTo(path string) error {
	if err := a.writer.Close(); err != nil {
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
