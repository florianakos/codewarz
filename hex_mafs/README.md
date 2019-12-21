## Simple_addition

There are several HEX encoded numbers on each line. Add up the numbers on each line and give the correct sum.

Example input:

```bash
$ cat /path/to/somefile.txt
0x539 0x539
0x7a69 0x7a69
```

Expected output:

```bash
$ ./simple_addition_solve.py /path/to/somefile.txt
2674
62674
```

The `SHA1` of the output shall be:
```
7d337536fb614a8bec4353b44ff682c0c8ee6b81
```

Proof:

```bash
$ python hexmafs.py input.txt | python -c "import sys,hashlib; print(hashlib.sha1(sys.stdin.read().strip()).hexdigest())"

7d337536fb614a8bec4353b44ff682c0c8ee6b81
```
