package util

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"strconv"
	"strings"

	"github.com/nfnt/resize"
)

// SizeProfile represents the size of the picture in pixels
type SizeProfile struct {
	Height uint
	Width  uint
}

// Square size image dimensions
var Square SizeProfile = SizeProfile{
	Height: 1080,
	Width:  1080,
}

// Vertical size image dimensions
var Vertical SizeProfile = SizeProfile{
	Height: 1080,
	Width:  1350,
}

// Horizontal size image dimensions
var Horizontal SizeProfile = SizeProfile{
	Height: 1080,
	Width:  566,
}

// NewSizeProfile creates a SizeProfile given a dimension string
// such as "500x500"
func NewSizeProfile(str string) (*SizeProfile, error) {
	values := strings.Split(str, "x")

	if values[0] != "" && values[1] != "" {
		width, err := strconv.ParseUint(values[0], 10, 64)

		if err != nil {
			return nil, err
		}

		height, err := strconv.ParseUint(values[1], 10, 64)

		if err != nil {
			return nil, err
		}

		sizeProfile := new(SizeProfile)

		sizeProfile.Height = uint(height)
		sizeProfile.Width = uint(width)

		return sizeProfile, nil
	}

	return nil, fmt.Errorf(`Value %s is not valid as dimensions. Valid values are "500x500"`, str)
}

// Resize image while keeping the aspect ratio of the original
// dimensions
func Resize(source []byte, profile SizeProfile) []byte {
	buff := new(bytes.Buffer)

	img, _, err := image.Decode(bytes.NewReader(source))

	if err != nil {
		log.Fatal(err)
	}

	resizedImage := resize.Resize(profile.Width, profile.Height, img, resize.Lanczos3)

	err = jpeg.Encode(buff, resizedImage, nil)

	if err != nil {
		log.Fatal(err)
	}

	return buff.Bytes()
}
