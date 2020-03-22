#!/usr/bin/env python
import sys

filename = sys.argv[1]
fp = open(filename, 'r')

for line in fp:
    parts = line.rstrip().split(" ")
    if len(parts) != 2:
        continue
    index = int(parts[0])
    numbers = list(map(int, parts[1].split(",")))
    # if index is out of bounds of the given list!
    if index > len(numbers) or index < 0 or len(numbers) < 1:
        continue
    # if chosen number is already the max of the whole list!
    if numbers[index-1] >= max(numbers):
        continue
    # in case the right side is not empty
    if len(numbers[index:]) != 0:
        if numbers[index-1] < max(numbers[index:]):
            for e in numbers[index:]:
                if e <= numbers[index-1]:
                    continue
                else:
                    print(e)
                    break
        else:
            for e in numbers[:index-1]:
                if e <= numbers[index-1]:
                    continue
                else:
                    print(e)
                    break
    else:
        for e in numbers[:index-1]:
            if e <= numbers[index-1]:
                continue
            else:
                print(e)
                break

fp.close()
