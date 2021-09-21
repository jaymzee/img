#include <fcntl.h>
#include <stdlib.h>
#include <stdint.h>
#include <sys/mman.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <unistd.h>
#include "fb.h"

struct image *
new_image(int xres, int yres)
{
    int pixlenbytes = xres * yres * 4; // RGBA format
    struct image *img = malloc(sizeof(struct image) + pixlenbytes - PIX_DFLTLEN);
    if (img == NULL) {
        return NULL;
    }
    img->xres = xres;
    img->yres = yres;
    return img;
}

// write image to framebuffer
// return -1 on failure, 0 on success
int write_image(struct image *img, int x, int y, struct fbinfo *fbinfo)
{
    int fd = open(fbinfo->device, O_RDWR);
    if (fd < 0) {
        return -1;
    }

    long pagesize = sysconf(_SC_PAGE_SIZE);
    int desired_offset = (y * (fbinfo->xres + fbinfo->pad) + x) << 2;
    int pages = desired_offset / pagesize;
    off_t offset = pages * pagesize;
    int delta = (desired_offset - offset) >> 2;
    int size = (fbinfo->xres + fbinfo->pad) * fbinfo->yres + delta;

    //get writable screen memory; 32bit color
    uint32_t *fb = mmap(NULL, size, PROT_READ | PROT_WRITE, MAP_SHARED, fd, offset);
    if (fb < 0) {
        close(fd);
        return -1;
    }

    int stride = fbinfo->xres + fbinfo->pad;

    int32_t *pixels = (void *)img->pix;
    for (int i = 0; i < img->yres; i++) {
        for (int j = 0; j < img->xres; j++) {
            fb[i*stride + j + x + delta] = *pixels++;
        }
    }

    munmap(fb, size);
    close(fd);
    return 0;
}
