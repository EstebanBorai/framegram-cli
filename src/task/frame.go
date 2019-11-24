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
func FrameImage(srcPath, outputPath string, sizeProfile util.SizeProfile) {
	imageBytes, err := ioutil.ReadFile(srcPath)

	if err != nil {
		log.Fatal(err)
	}

	resizedBytes, resizedDimensions := util.Resize(imageBytes, util.SizeProfile{
		Height: 0,
		Width:  sizeProfile.Width,
	})

	imagePoint := image.Point{0, -int(resizedDimensions.Height / 4)}

	buff := bytes.NewBuffer(resizedBytes)

	img, _, err := image.Decode(buff)

	if err != nil {
		log.Fatal(err)
	}

	bounds := image.Rect(0, 0, int(sizeProfile.Width), int(sizeProfile.Height))
	mask := image.NewRGBA(bounds)

	draw.Draw(mask, bounds, image.NewUniform(color.White), image.ZP, draw.Src)
	draw.Draw(mask, bounds, img, imagePoint, draw.Src)

	toImage, _ := os.Create(outputPath)
	defer toImage.Close()

	jpeg.Encode(toImage, mask, &jpeg.Options{
		Quality: 100,
	})
}
