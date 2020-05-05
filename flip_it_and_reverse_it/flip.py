#!/usr/bin/env python
import sys

filename = sys.argv[1]
f = open(filename, 'r')

counter = 1
for line in f:
    words = line.rstrip().split(" ")
    for w in words:
        if counter % 2 == 1:
            print(w)
        else:
            print(w[::-1])
        counter += 1


f.close()
