// Package iterm is for transering graphics images to the iTerm2 terminal using
// the protocol described at https://iterm2.com/3.2/documentation-images.html.
package iterm

import (
	"encoding/base64"
	"fmt"
	"io"
)

// WriteImage base64 encodes the the graphics image and writes it to the
// writer using the iTerm2 graphics protocol along with the args string,
// a semicolon seperated list of key=value pairs
//   infile, err := os.Open("cat.png")
//   img := image.Decode(infile)
//   iterm.WriteImage(os.Stdout, "inline=1", img)
func WriteImage(w io.Writer, args string, data []byte) error {
	size := len(data)
	enc := base64.StdEncoding
	encoded := make([]byte, enc.EncodedLen(size))
	enc.Encode(encoded, data)

	fmt.Fprintf(w, "\033]1337;File=size=%d", size)
	if len(args) > 0 {
		fmt.Fprintf(w, ";%s:", args)
	} else {
		fmt.Fprint(w, ":")
	}
	_, err := w.Write(encoded)
	if err != nil {
		return err
	}
	_, err = io.WriteString(w, "\x07")
	return err
}
