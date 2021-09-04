# kitty

Package kitty is for transfering graphics data to the Kitty terminal
using the Terminal Graphics Protocol described at
[https://sw.kovidgoyal.net/kitty/graphics-protocol/](https://sw.kovidgoyal.net/kitty/graphics-protocol/).

## Functions

### func [WriteBase64](/term.go#L32)

`func WriteBase64(w io.Writer, cmd string, data []byte) error`

WriteBase64 transfers graphics image data to the Kitty terminal.
It accepts a png image or an RGBA byte array or any other format supported
by the terminal graphics protocol. The data must be base64 encoded
and is broken into chunks <= 4K and written to the writer in terminal
Application Programming Commands (APC). The first chunk will contain the
kitty cmd string, a comma separated list of key=value pairs.

### func [WriteImage](/term.go#L19)

`func WriteImage(w io.Writer, cmd string, data []byte) error`

WriteImage transfers graphics image data to the Kitty terminal.
It accepts a png image or an RGBA byte array or any other format supported
by the protocol described in the link above. The data is base64 encoded and
written to the writer along with the kitty cmd string containing a comma
separated list of key=value pairs.

```go
infile, err := os.Open("cat.png")
img := image.Decode(infile)
kitty.WriteImage(os.Stdout, "a=T,f=100", img)
```
