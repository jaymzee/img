package term

// Winsize is the console window size
type Winsize struct {
	Rows uint16
	Cols uint16
	Xres uint16
	Yres uint16
}

// create window size with sensible defaults
func newWinsize() *Winsize {
	return &Winsize{24, 80, 0, 0}
}
