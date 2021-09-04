# iterm

Package iterm is for transering graphics images to the iTerm2 terminal using
the protocol described at [https://iterm2.com/3.2/documentation-images.html](https://iterm2.com/3.2/documentation-images.html).

## Functions

### func [WriteImage](/term.go#L17)

`func WriteImage(w io.Writer, args string, data []byte) error`

WriteImage base64 encodes the the graphics image and writes it to the
writer using the iTerm2 graphics protocol along with the args string,
a semicolon seperated list of key=value pairs

```go
infile, err := os.Open("cat.png")
img := image.Decode(infile)
iterm.WriteImage(os.Stdout, "inline=1", img)
```
