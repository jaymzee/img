package kitty

import (
	"encoding/base64"
	"io"
)

func serializeGfxCmd(cmd, payload []byte) []byte {
	img := make([]byte, 0, len(cmd)+len(payload)+6)
	img = append(append(img, '\x1b', '_', 'G'), cmd...)
	if len(payload) > 0 {
		img = append(append(img, ';'), payload...)
	}
	return append(img, '\x1b', '\\')
}

func writeChunked(w io.Writer, cmd string, data []byte) {
	var chunk []byte
	for len(data) > 0 {
		next := min(4096, len(data))
		chunk, data = data[:next], data[next:]
		if len(data) > 0 {
			cmd += ",m=1"
		} else {
			cmd += ",m=0"
		}
		w.Write(serializeGfxCmd([]byte(cmd), chunk))
	}
}

func WriteImage(w io.Writer, head string, data []byte) {
	enc := base64.StdEncoding
	encoded := make([]byte, enc.EncodedLen(len(data)))
	enc.Encode(encoded, data)
	writeChunked(w, head, encoded)
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
