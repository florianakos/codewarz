#!/usr/bin/env python
import sys

filename = sys.argv[1]
fp = open(filename, 'r')

for line in fp:
    elements = line.rstrip().split(" ")
    sum = 0
    for e in elements:
        sum = sum + int(e, 16)
    print(sum)

fp.close()
