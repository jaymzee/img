"""
read a file and try to display it as an image
in the terminal (iTerm2 or mintty)

supports PNG, JPEG, GIF
         Netpbm PGM, PPM, PBM
         ASCII test (numpy loadtxt)
"""

import sys

if len(sys.argv) < 2:
    print("Usage: imcat files")
    exit(2)

import matplotlib.image as mpimg
import matplotlib.pyplot as plt
import numpy as np
import os
import re
import subprocess


def load_pgm_ascii(filename):
    with open(filename) as f:
        lines = f.readlines()
    lines = [line for line in lines if not line.startswith('#')]
    fmt = lines[0]
    cols, rows = np.array(lines[1].split(' '), dtype=int)
    depth = int(lines[2])
    data = ''.join(lines[3:]).split()
    return np.array(data, dtype=np.byte).reshape(rows, cols), depth


def filetype(filename):
    cmd = f'file "{filename}"'
    results = subprocess.check_output(cmd, shell=False).decode('utf-8')
    return results.split('\n')[0]


def is_image(filetype):
    pattern = 'PNG|JPEG|GIF|Netpbm.+rawbits'
    return re.search(pattern, filetype) is not None


def is_pgm_ascii(filetype):
    pattern = 'Netpbm.+greymap.+ASCII'
    return re.search(pattern, filetype) is not None


def is_ascii(filetype):
    return 'ASCII' in filetype and 'Netpbm' not in filetype


def main(argv):
    for fname in argv[1:]:
        if not os.path.isfile(fname):
            print(f'{fname}: does not exist')
            exit(1)
        ftype = filetype(fname)
        depth = ''
        if is_image(ftype):
            im = mpimg.imread(fname)
        elif is_pgm_ascii(ftype):
            im, depth = load_pgm_ascii(fname)
        elif is_ascii(ftype):
            im = np.loadtxt(fname)
        else:
            print('invalid image')
            exit(1)
        print(ftype)
        print(im.dtype, im.shape, depth)
        plt.imshow(im)
        if im.ndim == 2:
            plt.colorbar()
        plt.show()


main(sys.argv)
