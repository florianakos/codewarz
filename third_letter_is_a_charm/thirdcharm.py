#!/usr/bin/env python3
import sys

def fix_word(w):
    wr = list(w)
    temp = wr[0]
    wr[0] = wr[1]
    wr[1] = wr[2]
    wr[2] = temp
    return ''.join(wr)

filename = sys.argv[1]
fp = open(filename, 'r')
for line in fp:
    words = line.split(" ")
    out = ""
    for w in words:
        if len(w) < 3:
            out += w + " "
        else:
            out += fix_word(w) + " "
    print(out.rstrip())
fp.close()
