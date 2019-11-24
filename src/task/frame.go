package task

import (
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/estebanborai/framegram/src/util"
)

// FrameImage add frames to a given image either Horizonal/Veritcal
func FrameImage(outputPath string, sizeProfile util.SizeProfile) {
	lowRight := image.Point{
		int(sizeProfile.Width),
		int(sizeProfile.Height),
	}

	upLeft := image.Point{
		0,
		0,
	}

	img := image.NewRGBA(image.Rectangle{
		upLeft,
		lowRight,
	})

	for x := 0; x < int(sizeProfile.Width); x++ {
		for y := 0; y < int(sizeProfile.Height); y++ {
			switch {
			case x < int(sizeProfile.Width/2) && y < int(sizeProfile.Height/2):
				img.Set(x, y, color.White)
			case x >= int(sizeProfile.Width/2) && y >= int(sizeProfile.Height/2):
				img.Set(x, y, color.Black)
			}
		}
	}

	f, _ := os.Create(outputPath)

	png.Encode(f, img)
}
