package main

import "github.com/S-IR/freible/compress"

func main() {
	compress.ArchiveFolder("./compression-folder-mockup", "test.zip", compress.ArchiveConfig{})
}
