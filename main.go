package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"log"
	"os"

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

const GOOS string = runtime.GOOS

func setBrightness(brightness int16, i int) {
	switch GOOS {
	case "linux":
		backlightDump, err := exec.Command("brightnessctl", "-lc", "backlight").Output()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		noOfDevicesRegex := regexp.MustCompile(" '(.*)' ")
		displays := noOfDevicesRegex.FindAll(backlightDump, -1)
		exec.Command("brightnessctl", "-d", string(displays[i][2:(len(displays[i])-2)]), "s", fmt.Sprintf("%d%%", brightness)).Output()

	case "windows":
		// fmt.Println("oof microshit windows detected")
		exec.Command("powershell", fmt.Sprintf("(Get-WmiObject -Namespace root/WMI -Class WmiMonitorBrightnessMethods).WmiSetBrightness(%d,%d)", i+1, brightness))

	case "darwin":
		fmt.Println("Sorry the changing brightness part is not supported for macOS")

	default:
		fmt.Println("Seriously tf is this OS")
	}
}

func main() {
	n := screenshot.NumActiveDisplays()

	// Flags
	rate := flag.Float64("rate", 0.5, "Set the Rescan Rate (0.5 == rescan every 0.5 sec to refresh)")
	mid := flag.Int("mid", 30, "Set the Normal Brightness (0-100)")
	deviate := flag.Int("deviate", 20, "Set the Maximum Deviation from the normal brightness")
	thresh := flag.Int("thresh", 3, "Set the treshold used while deciding if the change in lightness is large enough (implementaion is complicated, so find an appropriate value by experimenting)")
	// conf := flag.String("conf", "nope", "Use a config file")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n\t%s [Options]\n\n", os.Args[0], os.Args[0])
		fmt.Printf("Options:\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	// fmt.Println(int16(*mid))

	// Configurable Variables
	refreshRate := *rate
	var normalBrightness int16 = int16(*mid)
	var maxDeviation int16 = int16(*deviate)
	var threshold int8 = int8(*thresh)
	prevLightness := make([]int8, n)

	fmt.Printf("\nDetected %d displays connected\n", n)

	for range time.Tick(time.Millisecond * time.Duration(1000*refreshRate)) {
		go func() {
			unchanged := true
			k := screenshot.NumActiveDisplays()
			if n != k {
				unchanged = false
				n = k
			}

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
					if delta > threshold || delta < -threshold {
						brightness := (normalBrightness) + (50-int16(lightness[i]))*(maxDeviation)/50
						fmt.Printf("Changed Brightness for Screen #%d\nLightness: %d%%\nNewBrightness: %d%%\n\n", i, lightness[i], brightness)
						setBrightness(brightness, i)
					}
				}
				prevLightness = lightness
			}
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
