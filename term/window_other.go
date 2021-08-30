// +build plan9

package term

func GetWinsize() *Winsize {
	return newWinsize()
}
