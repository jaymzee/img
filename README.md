# Display Inline Images in the Terminal (iTerm2, mintty, Kitty, linux console)

## imgcat
it works in iTerm2 and mintty
```
pip install imgcat
```

Command-line interface (similar to iTerm2's imgcat):
``` sh
$ imgcat local_image.png
$ imgcat a.png b.png c.png
$ cat from_stdin.gif | imgcat

# height is 10 lines
$ imgcat a.png --height 10
```

Python API:
``` python
>>> from imgcat import imgcat

# from the content of image (e.g. buffer in python3, str in python2)
>>> imgcat(open("./local_image.png"))

# or numpy arrays!
>>> im = skimage.data.chelsea()   # [300, 451, 3] ndarray, dtype=uint8
>>> imgcat(im, height=7)

>>> im = np.asarray(Image.open('treefrog.png'))
>>> imgcat(im)

# matplotlib, PIL.Image, etc.
>>> imgcat(Image.fromarray(im))

>>> import matplotlib.pyplot as plt
>>> fig, ax = plt.subplots(); ax.plot([1, 2, 3, 4, 5])
>>> imgcat(fig)
```

Matplotlib Backend:
```
MPLBACKEND="module://imgcat" python draw_matplotlib.py
```

``` python
>>> import matplotlib
>>> matplotlib.use("module://imgcat")

>>> import matplotlib.pyplot as plt
>>> fig, ax = plt.subplots()
>>> ax.text(0.5, 0.5, "Hello World!");
>>> fig.show()
# an image shall be displayed on your terminal!
```

IPython magic (works both in terminal and notebook)
```
%load_ext imgcat
%imgcat skimage.data.chelsea()
```

## img Go Package
package        | description
---------------|-------------------------
img/plot       | simple plotting
img/term/fb    | framebuffer (linux console)
img/term/iterm | iTerm2 Terminal Graphics
img/term/kitty | Kitty Terminal Graphics

```go
infile, err := os.Open("cat.png")
img := image.Decode(infile)
iterm.WriteImage(os.Stdout, "inline=1", img)
```

## ImageMagick
resize images with:
```
convert frog.png -resize 64x64 frog_sm.png
```

reduce color palette with:
```
convert frog.png -colors 16 frog_16.png
```

reduce depth of color space with:
```
convert frog.png -depth 4 frog_d4.png
```

convert to Netpbm graymap ascii (magic P2)
```
convert frog.png -compress none frog.pgm
```

convert to Netpbm graymap binary (magic P5)
```
convert frog.png frog.pgm
```

identify files with:
```
identify frog.png
identify -verbose frog.png
```

### remarks
- jpeg is interesting because it preserves the comments from the PGM file

