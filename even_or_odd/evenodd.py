import sys

filename = sys.argv[1]
fp = open(filename, 'r')
for line in fp:
    if int(line)%2 == 0:
        print(line.rstrip()+" True")
    else:
        print(line.rstrip()+" False")
fp.close()
