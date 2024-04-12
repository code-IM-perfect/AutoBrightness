package main

import (
	"fmt"
	"image"
	"image/color"
	"log"

	// _ "image/png"
	// "os"

	"time"

	"github.com/kbinani/screenshot"
)

func getPercentLightness(img image.Image) int8 {
	imgSize := img.Bounds().Size()

	var sum uint
	for x := 0; x < imgSize.X; x++ {
		for y := 0; y < imgSize.Y; y++ {
			sum += uint(color.GrayModel.Convert(img.At(x, y)).(color.Gray).Y)
		}
	}
	sum = sum / uint(imgSize.X*imgSize.Y)

	var percentage int8 = int8(100 * int(sum) / 255)

	return percentage
}

func main() {

	n := screenshot.NumActiveDisplays()
	var normalBrightness int = 40
	var maxDeviation int = 20
	var threshold int8 = 3
	prevLightness := make([]int8, n)

	fmt.Printf("Detected %d displays connected\n", n)
	for i := 0; i < n; i++ {
		img, err := screenshot.CaptureDisplay(i)
		if err != nil {
			log.Fatalln(err)
		}
		prevLightness[i] = getPercentLightness(img)
		fmt.Printf("Display #%d\nLightness: %d%%\n\n", i, prevLightness[i])
	}

	for range time.Tick(time.Second * 2) {
		go func() {
			unchanged := true
			k := screenshot.NumActiveDisplays()
			if n != k {
				unchanged = false
				n = k
			}

			// var lightness []uint8
			lightness := make([]int8, n)

			copy(lightness, prevLightness)

			for i := 0; i < n; i++ {
				img, err := screenshot.CaptureDisplay(i)
				if err != nil {
					log.Fatalln(err)
				}
				lig := getPercentLightness(img)
				if unchanged && (prevLightness[i] != lig) {
					unchanged = false
				}
				lightness[i] = lig
			}

			if !(unchanged) {
				for i := 0; i < n; i++ {
					delta := (lightness[i]) - (prevLightness[i])
					// var delta int8 = 3
					// fmt.Printf("There was a change: delta =    %d %d\n", prevLightness[i], lightness[i])
					if delta > threshold || delta < -threshold {
						// var brightness int8 = normalBrightness + ((int8(lightness[i]) - 50) * maxDeviation / 50)
						var brightness int = (normalBrightness) + (int(lightness[i])-50)*(maxDeviation)/50
						fmt.Printf("Changed Brightness for Screen #%d\nLightness: %d%%\nNewBrightness: %d%%\n\n", i, lightness[i], brightness)
					}
				}
				prevLightness = lightness
			}
			// else {
			// 	fmt.Printf("No changes were found, %d\n", lightness[0])
			// }

		}()
	}

}

// inputFile, err := os.Open("test_img/test_img_copy.png")
// if err != nil {
// 	log.Fatalln(err)
// }
// defer inputFile.Close()

// img, _, err := image.Decode(inputFile)
// if err != nil {
// 	log.Fatalln(err)
// }

// var percentage uint = getPercentLightness(img)

// fmt.Printf("Brightness: %d%%\n", percentage)
