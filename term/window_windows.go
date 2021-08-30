package term

func GetWinsize() (*Winsize, error) {
	return &Winsize{24, 80, 0, 0}, nil
}
