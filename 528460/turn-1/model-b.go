package main

import (
	"fmt"
	"io"
	"os"
)

func rleCompress(input []byte, output io.Writer) error {
	var count uint16
	var prev byte
	if len(input) == 0 {
		return nil
	}
	prev = input[0]
	for i := 1; i < len(input); i++ {
		if input[i] == prev {
			count++
		} else {
			if err := writeRLEBlock(output, prev, count); err != nil {
				return err
			}
			prev = input[i]
			count = 1
		}
	}
	return writeRLEBlock(output, prev, count)
}

func writeRLEBlock(output io.Writer, b byte, count uint16) error {
	if _, err := output.Write([]byte{b}); err != nil {
		return err
	}
	var bCount [2]byte
	bCount[0] = byte(count >> 8)
	bCount[1] = byte(count & 0xFF)
	if _, err := output.Write(bCount[:]); err != nil {
		return err
	}
	return nil
}
func rleDecompress(input io.Reader, output io.Writer) error {
	var b [1]byte
	var count uint16
	for {
		if _, err := input.Read(b[:]); err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		prev := b[0]
		if _, err := input.Read(b[:]); err != nil {
			return err
		}
		count = uint16(b[0])<<8 | uint16(b[1])
		for i := 0; i < int(count); i++ {
			if _, err := output.Write([]byte{prev}); err != nil {
				return err
			}
		}
	}
	return nil
}

func main() {
	inputFilePath := "input.txt"
	outputFilePath := "input.txt.rle"
	decompressedFilePath := "decompressed.txt"
	// Compress file
	inputFile, err := os.Open(inputFilePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer inputFile.Close()
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer outputFile.Close()
	if err := rleCompress(inputFile, outputFile); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Compression successful!")
	// Decompress file
	rleFile, err := os.Open(outputFilePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rleFile.Close()
	decompressedFile, err := os.Create(decompressedFilePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer decompressedFile.Close()
	if err := rleDecompress(rleFile, decompressedFile); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Decompression successful!")
}
