#!/usr/bin/env python
import sys
import re

filename = sys.argv[1]
fp = open(filename, 'r')

def fix_line(input):
    fixed = input.lower()
    locations = []
    final = ""
    for i in [m.start() for m in re.finditer('i ', fixed)]:
        locations.append(i)
    for j in range(0, len(input)):
        if j in locations:
            final += fixed[j].upper()
        else:
            final += fixed[j]
    return final[0].upper()+final[1:]

for l in fp:
    line = l.rstrip()
    if line.count(".") > 1:
        sentences = line.split(".")
        fixed_sentences = []
        for s in sentences:
            if s != '':
                fixed_sentences.append(fix_line(s.lstrip()))
        print(". ".join(fixed_sentences)+".")
    else:
        print(fix_line(line))

fp.close()
