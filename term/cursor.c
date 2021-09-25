// +build !windows

#include <fcntl.h>
#include <stdio.h>
#include <stdlib.h>
#include <termios.h>
#include <unistd.h>

int getCursor(int *x, int *y)
{
    struct termios saved, temp;
    const char *dev;
    int fd;

    // get tty
    dev = ttyname(STDIN_FILENO);
    if (!dev) {
        return -1;
    }

    // open tty
    fd = open(dev, O_RDWR | O_NOCTTY);
    if (fd < 0) {
        return -1;
    }

    // save terminal settings
    if (tcgetattr(fd, &saved) < 0) {
        return -1;
    }

    // modify terminal settings
    // - turn off ICANON so input is available immediately
    // - turn off ECHO, otherwise the response to CSI [6n will
    //   be displayed in the terminal
    // - disable receiver
    temp = saved;
    temp.c_lflag &= ~ICANON;
    temp.c_lflag &= ~ECHO;
    //temp.c_cflag &= ~CREAD;
    tcsetattr(fd, TCSANOW, &temp);

    /* This escape sequence queries the current coordinates from the terminal.
     * The terminal responds on stdin */
    printf("\033[6n");
    if (scanf("\033[%d;%dR", y, x) != 2) {
        // restore terminal settings and indicate failure
        tcsetattr(fd, TCSANOW, &saved);
        close(fd);
        return -1;
    }

    // restore terminal settings
    tcsetattr(fd, TCSANOW, &saved);
    close(fd);

    return 0;
}
