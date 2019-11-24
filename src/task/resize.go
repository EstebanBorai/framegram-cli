package task

import (
	"io/ioutil"
	"log"

	"github.com/estebanborai/framegram/src/util"
)

// ResizeImage resizes an image
func ResizeImage(srcPath, outputPath string, sizeProfile util.SizeProfile) {
	data, err := ioutil.ReadFile(srcPath)

	if err != nil {
		log.Fatal(err)
	}

	resizedBytes := util.Resize(data, sizeProfile)

	err = ioutil.WriteFile(outputPath, resizedBytes, 0777)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("File is ready! %s\n", outputPath)
}
