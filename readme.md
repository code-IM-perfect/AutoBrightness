# AutoBrightness
Ever got blinded by a sudden blast of light mode? Well this programs aims to eliminate just that by automatically setting your brightness to a lower value when the content on your screen is brighter and higher when the contents are darker. It works for any number of screens.

<!-- #### Note: This project is not complete yet, this will only analyze your screen(s) and determine what your brightness should be -->

## How it works
- Takes all the pixels on the screen
- Converts all of them to greyscale (with weighted rgb values)
- Averages the 'whiteness' of each pixel (range- 0 to 255)
- Does this every 2 seconds.
- Compares it with the previous value
- If the change in whiteness is above a threshold, calculates an appropriate brightness and prints it
- Changes the brightness of the monitor accordingly ~~<span style="color:#ed8796">THIS IS STILL A TODO</span>~~ [Only for Linux and Windows]
- Does all this with **every** connected monitor

## Dependencies
A direct implementation on linux would have simply included editing `/sys/class/backlight/< Name of Display >/brightness` but that would have required sudo access which I felt was dangerous. I think you could create a new user and give it modification access to the file but that felt too complicated for the user. Some programs like `brightnessctl` and `light` [seem to add udev rules](https://wiki.archlinux.org/title/Backlight#Backlight_utilities) to do exactly that, but I don't understand how they do that programatically.

## Buiding
First clone this repo
```
git clone 'https://github.com/code-IM-perfect/AutoBrightness/releases'
```


## Prebuilt Binaries
Prebuilt binaries for Linux, Windows and MacOS can be found in [Releases](https://github.com/code-IM-perfect/AutoBrightness/releases).


## Usage
```

```