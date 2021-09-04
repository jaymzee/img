// Package kitty is for transfering graphics data to the Kitty terminal
// using the Terminal Graphics Protocol described at
// https://sw.kovidgoyal.net/kitty/graphics-protocol/.
package kitty

import (
	"encoding/base64"
	"io"
)

// WriteImage transfers graphics image data to the Kitty terminal.
// It accepts a png image or an RGBA byte array or any other format supported
// by the protocol described in the link above. The data is base64 encoded and
// written to the writer along with the kitty cmd string containing a comma
// separated list of key=value pairs.
//   infile, err := os.Open("cat.png")
//   img := image.Decode(infile)
//   kitty.WriteImage(os.Stdout, "a=T,f=100", img)
func WriteImage(w io.Writer, cmd string, data []byte) error {
	enc := base64.StdEncoding
	encoded := make([]byte, enc.EncodedLen(len(data)))
	enc.Encode(encoded, data)
	return WriteBase64(w, cmd, encoded)
}

// WriteBase64 transfers graphics image data to the Kitty terminal.
// It accepts a png image or an RGBA byte array or any other format supported
// by the terminal graphics protocol. The data must be base64 encoded
// and is broken into chunks <= 4K and written to the writer in terminal
// Application Programming Commands (APC). The first chunk will contain the
// kitty cmd string, a comma separated list of key=value pairs.
func WriteBase64(w io.Writer, cmd string, data []byte) error {
	var chunk []byte
	for len(data) > 0 {
		next := min(4096, len(data))
		chunk, data = data[:next], data[next:]
		if len(data) > 0 {
			cmd += ",m=1"
		} else {
			cmd += ",m=0"
		}
		_, err := w.Write(serializeGfxCmd([]byte(cmd), chunk))
		if err != nil {
			return err
		}
	}
	return nil
}

// serializeGfxCmd returns an APC containing the chunk
func serializeGfxCmd(cmd, payload []byte) []byte {
	img := make([]byte, 0, len(cmd)+len(payload)+6)
	img = append(append(img, '\x1b', '_', 'G'), cmd...)
	if len(payload) > 0 {
		img = append(append(img, ';'), payload...)
	}
	return append(img, '\x1b', '\\')
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
