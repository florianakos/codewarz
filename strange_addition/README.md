## Strange_addition

Simply add the numbers together on each line, also adding how many numbers where on each line. 

Example input:

```bash
$ cat /path/to/somefile.txt
1 1
2 2 2
3 3 3 3
4 4 4 4 4
```

Expected output:

```bash
$ ./strange_addition_solve.py /path/to/somefile.txt
4
9
16
25
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
