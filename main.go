package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/png"
	"log"
	"os"

	// "time"

	"github.com/kbinani/screenshot"
)

func getAvgGray(img image.Image) uint {
	imgSize := img.Bounds().Size()

	var sum uint
	for x := 0; x < imgSize.X; x++ {
		for y := 0; y < imgSize.Y; y++ {
			sum += uint(color.GrayModel.Convert(img.At(x, y)).(color.Gray).Y)
		}
	}
	sum = sum / uint(imgSize.X*imgSize.Y)

	return sum
}

func main() {
	inputFile, err := os.Open("test_img/test_img_copy.png")
	if err != nil {
		log.Fatalln(err)
	}
	defer inputFile.Close()

	img, _, err := image.Decode(inputFile)
	if err != nil {
		log.Fatalln(err)
	}

	var percentage uint = 100 * getAvgGray(img) / 255

	fmt.Printf("Brightness: %d%%\n", percentage)

	n := screenshot.NumActiveDisplays()

	for i := 0; i < n; i++ {
		bounds := screenshot.GetDisplayBounds(i)

		img, err := screenshot.CaptureRect(bounds)
		if err != nil {
			panic(err)
		}

		fmt.Printf("\nDisplay #%d\nLightness: %d%%\n\n", i, 100*getAvgGray(img)/255)
	}
}
