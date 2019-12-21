## Simple_addition

Can you decode the message?

Example input:

```bash
$ cat samplefile
67 87 78 123 87 104 48 95 100 48 51 36 110 39 116 95 76 105 107 51 95 48 114 68 105 78 64 114 89 95 119 48 114 68 36 125
```

Expected output:

```bash
$ ./solve.py /path/to/samplefile
CWN{Wh0_d03$nt_Lik3_0rDiN@rY_w0rD$}
```

The `SHA1` of the output shall be:
```
2fce8ce323e232e1a2f4f1c647cf00d33b9b255a
```

Proof:

```bash
$ python solve.py input.txt | python -c "import sys,hashlib; print(hashlib.sha1(sys.stdin.read().strip()).hexdigest())"

2fce8ce323e232e1a2f4f1c647cf00d33b9b255a
```
