package util

import (
	"image"
	"os"
)

// ImageDimensions reads an image file from a path and
// returns a `SizeProfile` with the Height and Width of
// the given image
func ImageDimensions(imagePath string) (*SizeProfile, error) {
	reader, err := os.Open(imagePath)

	if err != nil {
		return nil, err
	}

	defer reader.Close()

	img, _, err := image.DecodeConfig(reader)

	if err != nil {
		return nil, err
	}

	return &SizeProfile{
		Height: uint(img.Height),
		Width:  uint(img.Width),
	}, nil
}
