package main

import (
	"fmt"
)

type ImageSlice [][]uint8

func NewImageSlice(width, height int) ImageSlice {
	return make(ImageSlice, height, height)
	for i := range is {
		is[i] = make([]uint8, width, width)
	}
	return is
}

func (is ImageSlice) GetPixel(x, y int) uint8 {
	return is[y][x]
}

func (is ImageSlice) SetPixel(x, y int, value uint8) {
	is[y][x] = value
}

func (is ImageSlice) Width() int {
	return len(is[0])
}

func (is ImageSlice) Height() int {
	return len(is)
}

func (is ImageSlice) HorizontalFlip() {
	for i := range is {
		for j := 0; j < len(is[i])/2; j++ {
			is[i][j], is[i][len(is[i])-j-1] = is[i][len(is[i])-j-1], is[i][j]
		}
	}
}

func (is ImageSlice) VerticalFlip() {
	for i := 0; i < len(is)/2; i++ {
		is[i], is[len(is)-i-1] = is[len(is)-i-1], is[i]
	}
}

func (is ImageSlice) Grayscale() {
	for i := range is {
		for j := range is[i] {
			r, g, b := is[i][j]>>5&0x7, is[i][j]>>2&0x7, is[i][j]&0x3
			grayscale := uint8(0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b))
			is[i][j] = grayscale<<5 | grayscale<<2 | grayscale
		}
	}
}

func main() {
	// Create a new image slice
	image := NewImageSlice(5, 3)

	// Set some pixel values
	image.SetPixel(1, 1, 0xFF) // White
	image.SetPixel(2, 2, 0x00) // Black
	image.SetPixel(3, 0, 0x80) // Gray

	fmt.Println("Original Image:")
	for _, row := range image {
		fmt.Println(row)
	}

	// Perform some two-dimensional operations
	image.HorizontalFlip()
	image.VerticalFlip()
	image.Grayscale()

	fmt.Println("\nImage after operations:")
	for _, row := range image {
		fmt.Println(row)
	}
}
