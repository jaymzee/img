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
    void *p = malloc(sizeof(struct image) + xres * yres * 4);
    if (p == NULL) {
        return NULL;
    }

    struct image *img = p;
    img->pix = p + sizeof(struct image);
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

    //get writable screen memory; 32bit color
    uint32_t *fb = mmap(NULL, fbinfo->xres * fbinfo->yres,
                        PROT_READ | PROT_WRITE, MAP_SHARED, fd, 0);
    if (fb < 0) {
        return -1;
    }

    int stride = fbinfo->xres + fbinfo->pad;

    int32_t *pixels = (void *)img->pix;
    int n = 0;
    for (int i = 0; i < img->yres; i++) {
        for (int j = 0; j < img->xres; j++) {
            fb[(y+i)*stride + j + x] = pixels[n++];
        }
    }

    close(fd);

    return 0;
}
