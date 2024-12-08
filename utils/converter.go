package utils

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
)

func ConvertImageToPNG(img image.Image) ([]byte, error) {
	var pngBytes []byte
	pngBuffer := new(bytes.Buffer)

	if err := png.Encode(pngBuffer, img); err != nil {
		return nil, fmt.Errorf("error encoding PNG: %v", err)
	}
	pngBytes = pngBuffer.Bytes()

	return pngBytes, nil
}

func ConvertImageToJPEG(img image.Image, quality int) ([]byte, error) {
	var jpegBytes []byte
	jpegBuffer := new(bytes.Buffer)

	// Create JPEG encoder with specified quality (1-100)
	// Lower quality means more compression
	err := jpeg.Encode(jpegBuffer, img, &jpeg.Options{Quality: quality})
	if err != nil {
		return nil, fmt.Errorf("error encoding JPEG: %v", err)
	}
	jpegBytes = jpegBuffer.Bytes()

	return jpegBytes, nil
}
