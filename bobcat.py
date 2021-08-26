#!/usr/bin/env python3

# display an img file (from the bob ray tracer) in the kitty terminal

import sys
from base64 import standard_b64encode


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


# load a bob img file
def load_img(filename):
    # make sure the filename has an .img extension
    if not filename.lower().endswith('.img'):
        filename += '.img'

    data = []
    with open(filename, 'rb') as f:
        # read img header
        xres = int.from_bytes(f.read(2), byteorder='big')
        yres = int.from_bytes(f.read(2), byteorder='big')
        f.read(6)

        # read image
        data = []
        for y in range(yres):
            total = xres
            while total > 0:
                runlen = f.read(1)
                if runlen != b'':
                    runlen = runlen[0]
                    blu, grn, red = f.read(3)
                else: # at eof, so make up data
                    runlen = xres
                    blu, grn, red = 0, 255, 0
                total -= runlen
                for i in range(runlen):
                    data.extend((red, grn, blu))
    return {'s': xres, 'v': yres, 'data': bytes(data)}


def main(argv):
    if len(argv) == 2:
        write_chunked(a='T', f=24, **load_img(argv[1]))
        sys.stdout.write('\n')
    elif len(argv) > 2:
        for filename in argv[1:]:
            sys.stdout.write(f'{filename}\n')
            write_chunked(a='T', f=24, **load_img(filename))
            sys.stdout.write('\n')
    else:
        print('Usage: bobcat <imgfile>')


if __name__ == '__main__':
    main(sys.argv)
