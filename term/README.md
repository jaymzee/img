# term

## Functions

### func [Isatty](/terminal.go#L6)

`func Isatty() bool`

Isatty determines if stdout is a tty (is it a char mode device?)

## Types

### type [Winsize](/window.go#L4)

`type Winsize struct { ... }`

Winsize is the console window size

## Sub Packages

* [iterm](./iterm): Package iterm is for transering graphics images to the iTerm2 terminal using the protocol described at https://iterm2.com/3.2/documentation-images.html.

* [kitty](./kitty): Package kitty is for transfering graphics data to the Kitty terminal using the Terminal Graphics Protocol described at https://sw.kovidgoyal.net/kitty/graphics-protocol/.
