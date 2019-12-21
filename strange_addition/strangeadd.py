import sys

filename = sys.argv[1]
fp = open(filename, 'r')
for line in fp:
    sum = 0
    count = 0
    for element in line.split(" "):
        count += 1
        sum += int(element)
    print(count+sum)
fp.close()
