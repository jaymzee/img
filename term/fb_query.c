#include <fcntl.h>
#include <linux/fb.h>
#include <stdlib.h>
#include <sys/ioctl.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <unistd.h>

// return 0 on success
// return -1 on failure
int query_fb(const char *device, struct fb_var_screeninfo *fbinfo) {
    int fd = open(device, O_RDWR);
    if (fd < 0) {
        return -1;
    }

    if (ioctl(fd, FBIOGET_VSCREENINFO, fbinfo) < 0) {
        close(fd);
        return -1;
    }

    close(fd);
    return 0;
}
