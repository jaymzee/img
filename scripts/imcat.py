#!/usr/bin/env python3

# display a numpy array (text format) in the kitty terminal

import sys
import itertools
from base64 import standard_b64encode

flatten = itertools.chain.from_iterable

# kitty stuff

def serialize_gfxcmd(**cmd):
    payload = cmd.pop('payload', None)
    head = ','.join('{}={}'.format(k, v) for k, v in cmd.items())
    resp = [b'\033_G', head.encode('ascii')]
    if payload:
        resp.extend([b';', payload])
    resp.append(b'\033\\')
    return b''.join(resp)


def write_chunked(**cmd):
    data = standard_b64encode(cmd.pop('data'))
    while data:
        chunk, data = data[:4096], data[4096:]
        m = 1 if data else 0
        sys.stdout.buffer.write(serialize_gfxcmd(payload=chunk, m=m, **cmd))
        sys.stdout.flush()
        cmd.clear()


# could use numpy to load the data but it's better to not have any dependencies
# and the file format is very simple

# return the maximum value in array
def maximum(arr, init):
    m = init
    for r in arr:
        m = max(m, max(r))
    return m


# return a list of lines from the txt file
def load(filename):
    with open(filename, 'r') as f:
        lines = f.readlines()
    return lines


# parse the list of lines into a 2D array
def parse(lines):
    arr = [[float(v) for v in ln.strip().split(' ')] for ln in lines if not ln.startswith('#')]
    width, height = len(arr[0]), len(arr)
    return width, height, arr


# convert the 2D array into rgb data bytes
def to_rgb(width, height, array):
    def rgb(v):
        v = int(255 * v / max_luminance)
        return (v, v, v)
    max_luminance = maximum(array, 0)
    array = ((rgb(v) for v in r) for r in array)
    data = bytes(flatten(flatten(array)))
    return {'s':width, 'v':height, 'data':data}


def main(argv):
    if len(argv) < 2:
        img = to_rgb(*parse(sys.stdin.readlines()))
        write_chunked(a='T', f=24, **img)
        sys.stdout.write('\n')
    else:
        for filename in argv[1:]:
            img = to_rgb(*parse(load(filename)))
            sys.stdout.write(f"{filename}: {img['s']}x{img['v']}\n")
            write_chunked(a='T', f=24, **img)
            sys.stdout.write("\n")


if __name__ == '__main__':
    main(sys.argv)
