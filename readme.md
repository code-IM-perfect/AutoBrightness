# AutoBrightnessAdjuster
Ever got blinded by a sudden blast of light mode? Well this programs aims to eliminate just that by automatically setting your brightness to a lower value when the content on your screen is brighter and higher when the contents are darker. It works for any number of screens.

<!-- #### Note: This project is not complete yet, this will only analyze your screen(s) and determine what your brightness should be -->

## How it works
- Takes all the pixels on the screen
- Converts all of them to greyscale (with weighted rgb values)
- Averages the 'whiteness' of each pixel (range- 0 to 255)
- Repeats this in some interval (every `refreshRate` seconds)
- Compares it with the previous value
- If the change in whiteness is above a threshold, calculates an appropriate brightness and prints it
- Changes the brightness of the monitor accordingly [Only for Linux and Windows]
- Does all this with **every** connected monitor

## TODO
- ~~Add ability for this program to change brightness~~ [Mac is still not implemented tho]
- ~~Add flags to configure parameters~~
- ~~Windows Testing~~
- MacOS Testing
- Make the transition smooth
- Get rid of the `brightnessctl` dependency on Linux
- Disable it when videos are being played
- Add support for changing brightness on macOS (currently you can use this program in combination with apple script)

<!-- > **Note:** Windows and MacOS are untested, but they should work, I'll test once I get time. -->

## Note for Linux users
AutoBrightnessAdjuster requires that `brightnessctl` be installed.

A direct implementation on linux would have simply included editing `/sys/class/backlight/< Name of Display >/brightness` but that would have required sudo access which I felt was dangerous.

Some programs like `brightnessctl` and `light` [seem to add udev rules](https://wiki.archlinux.org/title/Backlight#Backlight_utilities) to get write access to the file, but I don't understand how exactly that works. I decided to go with `brightnessctl` because it is available in the `extra` repo of Arch (plus I already use it lol). This is in my TODO.

## Dependencies
[golang](go.dev) is needed to build, but is not a runtime dependency.

## Prebuilt Binaries
Prebuilt binaries for Linux, Windows and MacOS can be found in [Releases](https://github.com/code-IM-perfect/AutoBrightness/releases). 

Otherwise you can build it yourself from the instructions below.

## Buiding
First ensure that you have `go` installed

Clone this repo and move into the cloned directory
```
git clone 'https://github.com/code-IM-perfect/AutoBrightness'
cd AutoBrightness
```

Now just run
```
go build -o autoBrightness .
```
Or for windows-
```
go build -o autoBrightness.exe .
```
This will download all the required go dependencies (if not already intsalled) and will build the program.




## Usage
After building, it can be used by 
```
./autoBrightness  <<Options>>
```
#### Example usage
```
./autoBrightness --mid 20 --deviate 10
```


### Options
#### `--deviate [int]`
Set the Maximum Deviation from the normal brightness (default 20)

#### `-mid [int]`
Set the Normal Brightness (0-100) (default 30)
#### `-rate [float]`
Set the Rescan Rate (0.5 == rescan every 0.5 sec to refresh) (default 0.5)
#### `-thresh [int]`
Set the treshold used while deciding if the change in lightness is large enough (implementaion is complicated, so find an appropriate value by experimenting) (default 3)