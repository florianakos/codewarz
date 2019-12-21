## Simple_addition

 Simply add the numbers together on each line and print the sum of the numbers.

Example input:

```bash
$ cat /path/to/somefile.txt
1 1
2 2
3 3
a a
b b
c c
1 a
a 1
-1 2
2 -1
1.1 2.3
9.1
9.a
9.(
9.)
-9(
0.1 0.2
```

Expected output:

```bash
$ ./simple_addition_solve.py /path/to/somefile.txt
2
4
6
1
1
3.4
0.3
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
