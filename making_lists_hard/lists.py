#!/usr/bin/env python3

import sys

filename = sys.argv[1]
f = open(filename, 'r')

def is_integer(input):
    try:
        int(input)
        return True
    except:
        return False

final = []

for line in f:
    if "," not in line.rstrip():
        if is_integer(line.rstrip()): final.append(int(line.rstrip()))
    else:
        mini_list = []
        for c in line.rstrip().split(","):
            if is_integer(c): mini_list.append(int(c))
        if len(mini_list) == 1: final.append(mini_list[0])
        if len(mini_list) > 1: final.append(mini_list)

print(final, end ="")


f.close()
