#!/usr/bin/env python
import sys

filename = sys.argv[1]
fp = open(filename, 'r')

for line in fp:
    nums = line.rstrip().split(" ")
    if "." in line.rstrip():
        try:
          print(float(nums[0]) + float(nums[1]))
        except:
          continue
    else:
        try:
            print(int(nums[0]) + int(nums[1]))
        except:
            continue

fp.close()
