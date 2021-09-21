struct fbinfo {
    int xres;
    int yres;
    int pad;
    char *device;
};

struct image {
    int xres;
    int yres;
    char *pix;
};

int write_image(struct image *img, int x, int y, struct fbinfo *fbinfo);
struct image *
new_image(int xres, int yres);
