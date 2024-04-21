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
- Changes the brightness of the monitor accordingly ~~<span style="color:#ed8796">THIS IS STILL A TODO</span>~~ [Only for Linux and Windows]
- Does all this with **every** connected monitor

## TODO
- Make the transition smooth
- Get rid of the `brightnessctl` dependency on Linux
- Disable it when videos are being played

> **Note:** Windows and MacOS are untested, but they should work, I'll test once I get time.

## Note for Linux users
AutoBrightnessAdjuster requires that `brightnessctl` be installed.

A direct implementation on linux would have simply included editing `/sys/class/backlight/< Name of Display >/brightness` but that would have required sudo access which I felt was dangerous.

Some programs like `brightnessctl` and `light` [seem to add udev rules](https://wiki.archlinux.org/title/Backlight#Backlight_utilities) to get write access to the file, but I don't understand how exactly that works. I decided to go with `brightnessctl` because it is available in the `extra` repo of Arch (plus I already use it lol). This is in my TODO.

## Dependencies
[golang](go.dev) is needed to build, but is not a runtime dependency.

## Buiding
First ensure that you have `go` installed

Clone this repo
```
git clone 'https://github.com/code-IM-perfect/AutoBrightness'
```


## Prebuilt Binaries
Prebuilt binaries for Linux, Windows and MacOS can be found in [Releases](https://github.com/code-IM-perfect/AutoBrightness/releases).


## Usage
```

```





<!-- There is a screen capture liibrary as a go library but go would satisfy it automatically. 
### Linux
AutoBrightnessAdjuster utilises [`brightnessctl`](https://github.com/Hummer12007/brightnessctl) to set the brightness. Look at the [instructions for you distro](https://github.com/Hummer12007/brightnessctl#installation).\
The go dependencies will be satisfied by go itself.

#### Arch
```
sudo pacman -S brightnessctl
```
#### Debian / Ubuntu
```
sudo apt install brightnessctl
```
#### Redhat based distros (Fedora / opensuse)
```
sudo dnf install brightnessctl
```

### Windows
There are only go dependencies which will be handled by go itself.

### MacOS
Brightness control is not supported yet, otherwise there are no other non-go dependencies. -->