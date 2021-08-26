#!/usr/bin/env python3

# displays png images in the kitty terminal
# similar to 'icat'

import sys
from base64 import standard_b64encode


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


def main(argv):
    if len(argv) == 2:
        with open(argv[-1], 'rb') as f:
            data = f.read()
        write_chunked(a='T', f=100, data=data)
        sys.stdout.write('\n')
    elif len(argv) > 2:
        for filename in argv[1:]:
            with open(filename, 'rb') as f:
                data = f.read()
            sys.stdout.write(f'{filename}\n')
            write_chunked(a='T', f=100, data=data)
            sys.stdout.write('\n')
    else:
        print("Usage: pngcat pngfile")
        sys.exit(2)


if __name__ == '__main__':
    main(sys.argv)
