## Even Or Odd

Write a program that will read text file from the command line and determine if the numbers are even or odd.

Example input:

```bash
$ cat /path/to/file
0
11
2
33
4
55
6
77
8
99
100
```

Expected output:

```bash
$ ./even_or_odd.ps1 /path/to/file
0 True
11 False
2 True
33 False
4 True
55 False
6 True
77 False
8 True
99 False
100 True
```

The `SHA1` of the output shall be:
```
6544822421219cb2874396fb1d6d4c68c554654f
```

Proof:

```bash
$ python solve.py input.txt | python -c "import sys,hashlib; print(hashlib.sha1(sys.stdin.read().strip()).hexdigest())"

6544822421219cb2874396fb1d6d4c68c554654f
```
