package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

// compressFile compresses the input file using gzip and writes it to output file.
func compressFile(src, dst string) error {
	// Open the source file for reading.
	file, err := os.Open(src)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create the destination file where compressed content will be written.
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	// Create a new gzip writer.
	gzWriter := gzip.NewWriter(out)
	defer gzWriter.Close()

	// Copy the content from source file to the gzip writer.
	_, err = io.Copy(gzWriter, file)
	return err
}

// decompressFile decompresses the input gzip file and writes it to the output file.
func decompressFile(src, dst string) error {
	// Open the compressed file.
	file, err := os.Open(src)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create the gzip reader.
	gzReader, err := gzip.NewReader(file)
	if err != nil {
		return err
	}
	defer gzReader.Close()

	// Create the destination file to store decompressed content.
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	// Copy the content from the gzip reader to the output file.
	_, err = io.Copy(out, gzReader)
	return err
}

func main() {
	srcFile := "example.txt"
	compressedFile := "example.txt.gz"
	decompressedFile := "example_decompressed.txt"

	if err := compressFile(srcFile, compressedFile); err != nil {
		fmt.Println("Compression error:", err)
	} else {
		fmt.Println("Compression successful.", "Compressed to:", compressedFile)
	}

	if err := decompressFile(compressedFile, decompressedFile); err != nil {
		fmt.Println("Decompression error:", err)
	} else {
		fmt.Println("Decompression successful.", "Decompressed to:", decompressedFile)
	}
}
