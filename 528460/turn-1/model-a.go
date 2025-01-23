package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func rleCompress(input []byte) []byte {
	var result bytes.Buffer
	length := len(input)

	for i := 0; i < length; i++ {
		count := 1
		for i+1 < length && input[i] == input[i+1] {
			count++
			i++
		}
		result.WriteByte(input[i])
		result.WriteByte(byte(count))
	}

	return result.Bytes()
}

func rleDecompress(input []byte) []byte {
	var result bytes.Buffer

	for i := 0; i < len(input); i += 2 {
		b := input[i]
		count := int(input[i+1])
		for j := 0; j < count; j++ {
			result.WriteByte(b)
		}
	}

	return result.Bytes()
}

func compressFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	outputFileName := filename + ".rle"
	out, err := os.Create(outputFileName)
	if err != nil {
		return err
	}
	defer out.Close()

	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, f)
	if err != nil {
		return err
	}

	compressedData := rleCompress(buf.Bytes())
	_, err = out.Write(compressedData)
	return err
}

func decompressFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	outputFileName := filename + ".decompressed"
	out, err := os.Create(outputFileName)
	if err != nil {
		return err
	}
	defer out.Close()

	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, f)
	if err != nil {
		return err
	}

	decompressedData := rleDecompress(buf.Bytes())
	_, err = out.Write(decompressedData)
	return err
}

func main() {
	filename := "example.txt"

	err := compressFile(filename)
	if err != nil {
		fmt.Println("Error compressing:", err)
	}

	err = decompressFile(filename + ".rle")
	if err != nil {
		fmt.Println("Error decompressing:", err)
	}
}
