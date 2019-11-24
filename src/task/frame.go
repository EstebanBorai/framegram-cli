package task

import (
	"bytes"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"io/ioutil"
	"log"
	"os"

	"github.com/estebanborai/framegram/src/util"
)

// FrameImage add frames to a given image either Horizonal/Veritcal
func FrameImage(srcPath, outputPath string) {
	var squareSize int
	var resizeProfile util.SizeProfile

	imageBytes, err := ioutil.ReadFile(srcPath)
	originalFileDimensions, _ := util.ImageDimensions(srcPath)

	if err != nil {
		log.Fatal(err)
	}

	switch {
	case originalFileDimensions.Height > originalFileDimensions.Width:
		resizeProfile.Height = originalFileDimensions.Height
		resizeProfile.Width = 0
		squareSize = int(originalFileDimensions.Height)
		break
	case originalFileDimensions.Width > originalFileDimensions.Height:
		resizeProfile.Height = 0
		resizeProfile.Width = originalFileDimensions.Width
		squareSize = int(originalFileDimensions.Width)
		break
	case originalFileDimensions.Height == originalFileDimensions.Width:
		resizeProfile.Height = originalFileDimensions.Height
		resizeProfile.Width = originalFileDimensions.Width
		squareSize = int(originalFileDimensions.Height)
		break
	}

	resizedBytes, resizedDimensions := util.Resize(imageBytes, resizeProfile)

	imagePoint := calcImagePoint(&resizedDimensions, &resizeProfile)

	buff := bytes.NewBuffer(resizedBytes)

	img, _, err := image.Decode(buff)

	if err != nil {
		log.Fatal(err)
	}

	bounds := image.Rect(0, 0, squareSize, squareSize)
	mask := image.NewRGBA(bounds)

	draw.Draw(mask, bounds, image.NewUniform(color.White), image.ZP, draw.Src)
	draw.Draw(mask, bounds, img, *imagePoint, draw.Src)

	toImage, _ := os.Create(outputPath)

	defer toImage.Close()

	jpeg.Encode(toImage, mask, &jpeg.Options{
		Quality: 100,
	})
}

// calcImagePoint returns the first point of the target image
// into the mask given the actual dimensions of the
// input image after resize
func calcImagePoint(imageDimensions, maskDimensions *util.SizeProfile) *image.Point {
	height, width := imageDimensions.Height, imageDimensions.Width

	if width > height {
		return &image.Point{
			X: 0,
			Y: -int(height / 4),
		}
	}

	if height > width {
		return &image.Point{
			X: 0,
			Y: int(maskDimensions.Width / 4),
		}
	}

	return &image.Point{
		X: 0,
		Y: 0,
	}
}
