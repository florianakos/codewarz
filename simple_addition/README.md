## Simple_addition

 Simply add the numbers together on each line and print the sum of the numbers.

Example input:

```bash
$ cat /path/to/somefile.txt
1 1
2 2
3 3
4 4
5 5
-1 -2
.5 10
```

Expected output:

```bash
$ ./simple_addition_solve.py /path/to/somefile.txt
2
4
6
8
10
-3
10.5
```

The `SHA1` of the output shall be:
```
cba88f9c5e215b1e28e65800ff2fd34ecb5f850a
```

Proof:

```bash
$ go run bucketter.go input.txt | python -c "import sys,hashlib; print(hashlib.sha1(sys.stdin.read().strip()).hexdigest())"

cba88f9c5e215b1e28e65800ff2fd34ecb5f850a
```
