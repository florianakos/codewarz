## Simple_addition

There are several binary numbers on each line. Add up the binary numbers on each line and give the correct sum.

Example input:

```bash
$ cat /path/to/somefile.txt
0b1 0b10 0b11 0b100 0b101
0b1 0b11 0b111 0b1111 0b11111
```

Expected output:

```bash
$ ./simple_addition_solve.py /path/to/somefile.txt
15
57
```

The `SHA1` of the output shall be:
```
cba88f9c5e215b1e28e65800ff2fd34ecb5f850a
```

Proof:

```bash
$ python stupid.py input.txt | python -c "import sys,hashlib; print(hashlib.sha1(sys.stdin.read().strip()).hexdigest())"

cba88f9c5e215b1e28e65800ff2fd34ecb5f850a
```
