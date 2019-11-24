package task

import (
	"bytes"
	"image"
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

	// TODO: Get image size

	resizedBytes := util.Resize(imageBytes, util.SizeProfile{
		Height: 0,
		Width:  sizeProfile.Width,
	})

	buff := bytes.NewBuffer(resizedBytes)

	pict := image.NewRGBA(image.Rect(0, 0, int(sizeProfile.Width), int(sizeProfile.Height)))

	img, _, err := image.Decode(buff)

	if err != nil {
		log.Fatal(err)
	}

	draw.Draw(pict, pict.Bounds(), img, image.Point{0, -1 * start}, draw.Src)

	toImage, _ := os.Create(outputPath)
	defer toImage.Close()

	jpeg.Encode(toImage, pict, nil)
}
