import sys

filename = sys.argv[1]
fp = open(filename, 'r')
output = ""
for line in fp:
    elements = line.rstrip().split(" ")
    for e in elements:
        output += str(chr(int(e)))
print(output)
fp.close()
