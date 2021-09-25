package term

import "golang.org/x/sys/windows"

// GetWinsize returns the console window size.
func GetWinsize() *Winsize {
	// fetch window size from Windows API
	var csbi windows.ConsoleScreenBufferInfo
	h, err := windows.GetStdHandle(windows.STD_OUTPUT_HANDLE)
	if err != nil {
		panic(err)
	}
	err = windows.GetConsoleScreenBufferInfo(h, &csbi)
	if err != nil {
		// handle invalid in mintty
		//panic(err)
		return &Winsize{80, 24, 0, 0}
	}
	cols := csbi.Window.Right - csbi.Window.Left + 1
	rows := csbi.Window.Bottom - csbi.Window.Top + 1

	// return a window size that is no smaller than the default
	w := newWinsize()
	w.Cols = uint16(max(cols, int16(w.Cols)))
	w.Rows = uint16(max(rows, int16(w.Rows)))
	return w
}

func max(a int16, b int16) int16 {
	if a > b {
		return a
	}
	return b
}
