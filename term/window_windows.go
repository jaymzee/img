package term

type Winsize struct {
	Rows uint16
	Cols uint16
	Xres uint16
	Yres uint16
}

func GetWinsize() (*Winsize, error) {
	return &Winsize{24, 80, 0, 0}, nil
}
