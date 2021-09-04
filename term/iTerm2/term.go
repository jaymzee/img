package iTerm2

import (
	"encoding/base64"
	"fmt"
	"io"
)

func WriteImage(w io.Writer, args string, data []byte) {
	size := len(data)
	enc := base64.StdEncoding
	encoded := make([]byte, enc.EncodedLen(size))
	enc.Encode(encoded, data)

	fmt.Fprintf(w, "\033]1337;File=size=%d", size)
	if len(args) > 0 {
		fmt.Fprintf(w, ";%s", args)
	}
	io.WriteString(w, ":")
	w.Write(encoded)
	io.WriteString(w, "\x07")
}
