package main

import (
	"fmt"
	"image/color"
	"math"
)

type Image [][]uint8

// NewImage creates a new Image from a 2D slice of pixels.
func NewImage(height, width int) Image {
	return make(Image, height)
	for i := range NewImage(height, width) {
		NewImage(height, width)[i] = make([]uint8, width)
	}
}

// NewImageFromData creates a new Image from a 2D slice of pixel data.
func NewImageFromData(data [][]uint8) Image {
	return Image(data)
}

// ApplyGray converts the image to grayscale by applying a standard luminance formula.
func (im Image) ApplyGray() {
	for i := range im {
		for j := range im[i] {
			r := im[i][j]
			g := im[i][j]
			b := im[i][j]
			gray := uint8(0.2126*float64(r) + 0.7152*float64(g) + 0.0722*float64(b))
			im[i][j] = gray
		}
	}
}

// AdjustBrightness adjusts the brightness of the image.
func (im Image) AdjustBrightness(amount int) {
	maxPixel := uint8(255)
	if amount > 100 {
		amount = 100
	} else if amount < -100 {
		amount = -100
	}
	adjustmentFactor := float64(amount) / 100.0

	for i := range im {
		for j := range im[i] {
			newValue := math.Clamp(float64(im[i][j])+adjustmentFactor*float64(maxPixel), 0, float64(maxPixel)).(uint8)
			im[i][j] = newValue
		}
	}
}

// ApplySobel performs Sobel edge detection on the image.
func (im Image) ApplySobel() {
	Gx := make([][]float64, len(im), len(im))
	Gy := make([][]float64, len(im), len(im))

	sobelX := []float64{-1, 0, 1, -2, 0, 2, -1, 0, 1}
	sobelY := []float64{-1, -2, -1, 0, 0, 0, 1, 2, 1}

	for i := 1; i < len(im)-1; i++ {
		for j := 1; j < len(im[i])-1; j++ {
			Gx[i][j] = 0
			Gy[i][j] = 0
			for n := -1; n <= 1; n++ {
				for m := -1; m <= 1; m++ {
					Gx[i][j] += float64(sobelX[4+n*3+m]) * float64(im[i+n][j+m])
					Gy[i][j] += float64(sobelY[4+n*3+m]) * float64(im[i+n][j+m])
				}
			}
		}
	}

	for i := 1; i < len(im)-1; i++ {
		for j := 1; j < len(im[i])-1; j++ {
			grad := math.Sqrt(Gx[i][j]*Gx[i][j] + Gy[i][j]*Gy[i][j])
			newValue := float64(max(0, min(255, uint8(grad))))
			im[i][j] = uint8(newValue)
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	width, height := 3, 3
	data := [][]uint8{
		{255, 128, 0},
		{128, 255, 128},
		{0, 128, 255},
	}

	im := NewImageFromData(data)
	fmt.Println("Original Image:")
	printImage(im, width)

	fmt.Println("\nGrayscale Image:")
	im.ApplyGray()
	printImage(im, width)

	fmt.Println("\nAdjusted Brightness Image (brighter):")
	im.AdjustBrightness(20)
	printImage(im, width)

	fmt.Println("\nSobel Edge Detection Image:")
	im.ApplySobel()
	printImage(im, width)
}

// Helper function to print the image
func printImage(im Image, width int) {
	for _, row := range im {
		for _, pixel := range row {
			fmt.Printf("%03d ", pixel)
		}
		fmt.Println()
	}
	fmt.Println()
}
