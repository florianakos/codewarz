#!/usr/bin/env python
import sys

filename = sys.argv[1]
fp = open(filename, 'r')

for line in fp:
    sum = 0
    for n in list(map(float, line.rstrip().split(","))): sum += n
    print(sum)
    
fp.close()
