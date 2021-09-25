// +build windows plan9

package term

import "fmt"

func TtyName() string {
	return "unknown"
}

func GetCursorCoord() (int, int, error) {
	return 0, 0, fmt.Errorf("GetCursorCoord: not supported on this os")
}
