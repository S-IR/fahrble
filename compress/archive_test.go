package compress_test

import (
	"testing"

	"github.com/S-IR/freible/compress"
)

func TestZip(t *testing.T) {
	err := compress.ArchiveFolder("../compression-folder-mockup", "test.zip", compress.ArchiveConfig{
		ArchiveType: compress.ZipType,
	})
	if err != nil {
		panic(err)
	}
}

func TestGzipTar(t *testing.T) {
	err := compress.ArchiveFolder("../compression-folder-mockup", "test.tar.gz", compress.ArchiveConfig{
		ArchiveType: compress.TarGzipType,
	})
	if err != nil {
		panic(err)
	}
}
func BenchmarkZip(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := compress.ArchiveFolder("../compression-folder-mockup", "test.zip", compress.ArchiveConfig{
			ArchiveType: compress.ZipType,
		})
		if err != nil {
			panic(err)
		}
	}
}
