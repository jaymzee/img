#include <fcntl.h>
#include <stdlib.h>
#include <stdint.h>
#include <sys/mman.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <unistd.h>
#include "fb.h"

#include <stdio.h>

/*
void draw_colors(uint32_t *fb, uint32_t xres, uint32_t yres, uint32_t pad)
{
    // draw some color bars
    uint32_t colors[7] = {0xff0000, 0xff8000, 0xffff00, 0x00ff00, 0x00ffff, 0x0000ff, 0x8000ff};
    for (int i = 0; i < 64; i++) {
        for (int j = 0; j < 7; j++) {
            int offset = i*(xres+pad) + 10*j;
            for (int k = 0; k < 10; k++) {
                fb[offset+k] = colors[j];
            }
        }
    }
}
*/

int write_image(struct image *img, struct fbinfo *fbinfo, char *imgdata)
{
    int fd = open(fbinfo->device, O_RDWR);
    if (fd < 0) {
        return 0;
    }

    //get writable screen memory; 32bit color
    uint32_t *fb = mmap(NULL, fbinfo->xres * fbinfo->yres,
                        PROT_READ | PROT_WRITE, MAP_SHARED, fd, 0);

    int stride = fbinfo->xres + fbinfo->pad;

    printf("\n%d\n", img->yres * img->xres * 4);
    printf("%d\n", img->length);
    int32_t *pixels = (void *)imgdata;
    int n = 0;
    for (int i = 0; i < img->yres; i++) {
        for (int j = 0; j < img->xres; j++) {
            if (n*4 < img->length) {
                fb[i*stride + j] = pixels[n];
            }
            n++;
        }
    }

    close(fd);

    return 1;
}
