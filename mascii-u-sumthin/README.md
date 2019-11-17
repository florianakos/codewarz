## Mascii-U-Sumthin challenge

Our ordinal generator is on the fritz and is spewing out all sorts of numbers, sadface.jpg. There appears to be 4 different types??? I thought numbers were numbers, I don't get it. Can you print out the ascii that corresponds to the input because I don't know how to number.

Sample input:
```bash
$ cat /path/to/somefile.txt
0x54 0x68 105 0b1110011 0b100000 0x69 0x73 0x20 97 32 0x6c 0b1101001 0x6e 0b1100101 0b101110 0b1010 0b1010100 0150 0x69 0b1110011 0b100000 105 0x73 0x20 0b1100001 0b1101100 0b1110011 0157 040 0x61 0x20 0x6c 0b1101001 0x6e 0b1100101 056 012 0x54 104 101 114 0145 040 0b1100111 0157 0145 0b1110011 0b100000 0b1110100 0x68 0145 040 0x6e 0145 105 0147 0x68 0142 0b1101111 114 104 0157 0b1101111 0b1100100 33 10
```

The `SHA1` of the output shall be:
```
806a1f499e2b1b89948bae2d5a90da59617739ba
```

Proof:

```bash
$ go run mascii.go input.txt | python -c "import sys,hashlib; print(hashlib.sha1(sys.stdin.read().strip()).hexdigest())"

806a1f499e2b1b89948bae2d5a90da59617739ba
```
