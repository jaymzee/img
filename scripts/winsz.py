#!/usr/bin/env python3
# similar to stty size but gives screen dimensions in pixels too

import array, fcntl, sys, termios

buf = array.array('H', [0, 0, 0, 0])
fcntl.ioctl(sys.stdout, termios.TIOCGWINSZ, buf)
print(*buf)
