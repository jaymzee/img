package plot

import (
	"encoding/base64"
	"fmt"
	"os"
)

func WriteImage(data []byte) {
	size := len(data)
	head := fmt.Sprintf("\033]1337;File=size=%v;inline=1:", size)
	enc := base64.StdEncoding
	encoded := make([]byte, enc.EncodedLen(size))
	enc.Encode(encoded, data)

	os.Stdout.WriteString(head)
	os.Stdout.Write(encoded)
	os.Stdout.WriteString("\x07")
}
