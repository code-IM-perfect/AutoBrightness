package main

import (
	"fmt"
	"image"
	"image/color"
	"log"

	// _ "image/png"
	// "os"
	"os/exec"
	"regexp"
	"runtime"

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

// func setBrightnessAll(brightness []int16) {
// 	switch runtime.GOOS {
// 	case "linux":
// 		// fmt.Println("yoo deez linux")

// 		noOfDevicesRegex := regexp.MustCompile(" '(.*)' ")
// 		backlightDump, err := exec.Command("brightnessctl", "-lc", "backlight").Output()
// 		if err != nil {
// 			fmt.Println(err.Error())
// 			return
// 		}
// 		displays := noOfDevicesRegex.FindAll(backlightDump, -1)

// 		// fmt.Printf("the devices: %q\n", displays)

// 		for i := 0; i < len(displays); i++ {
// 			// out, err := exec.Command("brightnessctl", "-d", strings.ReplaceAll(string(displays[i]), "'", " "), "s", fmt.Sprintf("%d%%", brightness[i])).Output()
// 			out, err := exec.Command("brightnessctl", "-d", "intel_backlight", "s", fmt.Sprintf("%d%%", brightness[i])).Output()
// 			fmt.Println(out, err)
// 			fmt.Printf("%s %s %s %s %s\n", "brightnessctl", "-d", strings.ReplaceAll(string(displays[i]), " ", ""), "s", fmt.Sprintf("%d%%", brightness[i]))
// 		}

// 	case "windows":
// 		fmt.Printf("microshit windows detected")

// 	default:
// 		fmt.Println("tf is this OS")
// 	}
// }

func setBrightness(brightness int16, i int) {
	switch runtime.GOOS {
	case "linux":
		// fmt.Println("yoo deez linux")

		noOfDevicesRegex := regexp.MustCompile(" '(.*)' ")
		backlightDump, err := exec.Command("brightnessctl", "-lc", "backlight").Output()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		displays := noOfDevicesRegex.FindAll(backlightDump, -1)

		// fmt.Printf("the devices: %q\n", displays)

		exec.Command("brightnessctl", "-d", string(displays[i][2:(len(displays[i])-2)]), "s", fmt.Sprintf("%d%%", brightness)).Output()
		// out, err := exec.Command("brightnessctl", "-d", "intel_backlight", "s", fmt.Sprintf("%d%%", brightness)).Output()
		// fmt.Println(out, err)
		// fmt.Printf("%s %s %s %s %s\n", "brightnessctl", "-d", strings.ReplaceAll(string(displays[i]), " ", ""), "s", fmt.Sprintf("%d%%", brightness))
		// fmt.Printf("%s %s %s %s %s\n", "brightnessctl", "-d", string(displays[i][2:(len(displays[i])-2)]), "s", fmt.Sprintf("%d%%\n", brightness))

		// for i := 0; i < len(displays); i++ {
		// 	// out, err := exec.Command("brightnessctl", "-d", strings.ReplaceAll(string(displays[i]), "'", " "), "s", fmt.Sprintf("%d%%", brightness[i])).Output()
		// 	out, err := exec.Command("brightnessctl", "-d", "intel_backlight", "s", fmt.Sprintf("%d%%", brightness)).Output()
		// 	fmt.Println(out, err)
		// 	fmt.Printf("%s %s %s %s %s\n", "brightnessctl", "-d", strings.ReplaceAll(string(displays[i]), " ", ""), "s", fmt.Sprintf("%d%%", brightness[i]))
		// }

	case "windows":
		fmt.Println("oof microshit windows detected")
		exec.Command("powershell", fmt.Sprintf("(Get-WmiObject -Namespace root/WMI -Class WmiMonitorBrightnessMethods).WmiSetBrightness(%d,%d)", i+1, brightness))

	case "darwin":
		fmt.Println("Sorry the changing brightness part is not supported for macOS")

	default:
		fmt.Println("Seriously tf is this OS")
	}
}

func main() {

	// setBrightness(30, 0)

	n := screenshot.NumActiveDisplays()
	refreshRate := 0.5
	var maxDeviation int16 = 20
	var threshold int8 = 3
	prevLightness := make([]int8, n)

	fmt.Println()

	fmt.Printf("Detected %d displays connected\n", n)
	// for i := 0; i < n; i++ {
	// 	img, err := screenshot.CaptureDisplay(i)
	// 	if err != nil {
	// 		log.Fatalln(err)
	// 	}
	// 	prevLightness[i] = getPercentLightness(img)
	// 	fmt.Printf("Display #%d\nLightness: %d%%\n\n", i, prevLightness[i])
	// }

	for range time.Tick(time.Millisecond * time.Duration(1000*refreshRate)) {
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
						brightness := (normalBrightness) + (50-int16(lightness[i]))*(maxDeviation)/50
						fmt.Printf("Changed Brightness for Screen #%d\nLightness: %d%%\nNewBrightness: %d%%\n\n", i, lightness[i], brightness)
						setBrightness(brightness, i)
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
