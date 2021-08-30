package term

import "golang.org/x/sys/windows"

func GetWinsize() *Winsize {
	// fetch window size from Windows API
	var csbi windows.ConsoleScreenBufferInfo
	h, err := windows.GetStdHandle(windows.STD_OUTPUT_HANDLE)
	if err != nil {
		panic(err)
	}
	windows.GetConsoleScreenBufferInfo(h, &csbi)
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
	} else {
		return b
	}
}
