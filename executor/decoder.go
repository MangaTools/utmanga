package executor

import (
	"fmt"
	"image"
	"io"
	"os"
	"path/filepath"

	_ "image/jpeg"
	_ "image/png"
)

func readImageFromReader(input io.Reader) (image.Image, error) {
	parsedImage, _, err := image.Decode(input)
	if err != nil {
		return nil, fmt.Errorf("decode image: %w", err)
	}

	return parsedImage, nil
}

func readImageFile(imagePath string) (image.Image, error) {
	imageFile, err := os.Open(imagePath)
	if err != nil {
		return nil, fmt.Errorf("open image file \"%s\": %w", filepath.Base(imagePath), err)
	}
	defer imageFile.Close()

	img, err := readImageFromReader(imageFile)
	if err != nil {
		return nil, fmt.Errorf("decode image file %s: %w", imagePath, err)
	}

	return img, nil
}
