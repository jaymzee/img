struct fbinfo {
    int xres;
    int yres;
    int pad;
    char *device;
};

#define PIX_DFLTLEN 8

struct image {
    int xres;
    int yres;
    char pix[PIX_DFLTLEN];
};

int write_image(struct image *img, int x, int y, struct fbinfo *fbinfo);
struct image *new_image(int xres, int yres);
