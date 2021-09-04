// +build plan9

package term

// Getwinsize returns the console window size. If the os doesn't have an api
// for this a default window size of 24 80 0 0 is returned
func GetWinsize() *Winsize {
	return newWinsize()
}
