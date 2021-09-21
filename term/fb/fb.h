struct fbinfo {
    int xres;
    int yres;
    int pad;
    char *device;
};

struct image {
    int xres;
    int yres;
    int length;
};

int write_image(struct image *img, struct fbinfo *fbinfo, char *imgdata);
