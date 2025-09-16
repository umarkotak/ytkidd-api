package main

import (
	"fmt"
	"image/jpeg"
	"os"
)

func quickExec() {
	// Open the input JPEG file
	inputFile, err := os.Open("file_bucket/books/berhitung-1pdf/0001.jpeg")
	if err != nil {
		fmt.Printf("Error opening input file: %v\n", err)
		return
	}
	defer inputFile.Close()

	// Decode the JPEG image
	img, err := jpeg.Decode(inputFile)
	if err != nil {
		fmt.Printf("Error decoding JPEG: %v\n", err)
		return
	}

	// Create the output file for the compressed image
	outputFile, err := os.Create("file_bucket/books/berhitung-1pdf/0001-40.jpeg")
	if err != nil {
		fmt.Printf("Error creating output file: %v\n", err)
		return
	}
	defer outputFile.Close()

	// Set the desired compression quality (0-100, higher is better quality/larger file size)
	// For significant compression, a value like 75 or 80 is often used.
	quality := 40

	// Encode the image to the output file with the specified quality
	err = jpeg.Encode(outputFile, img, &jpeg.Options{Quality: quality})
	if err != nil {
		fmt.Printf("Error encoding JPEG: %v\n", err)
		return
	}

	fmt.Println("JPEG image compressed successfully!")
}
