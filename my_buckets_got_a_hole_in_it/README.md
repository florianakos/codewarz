## My_buckets_got_a_hole_in_it

Parse a file and count instances for each group of '10s (0-9, 10-19, 20-29,...) up to the max group from the input. You should account for every group of 10 available in range from 0 to the max group of 10's from the input.

Example input:

```bash
$ cat /path/to/somefile.txt
12
87
55
32
99
85
```

Expected output:

```bash
$ ./My_buckets_got_a_hole_in_it_solve.py /path/to/somefile.txt
0101010021
```

The `SHA1` of the output shall be:
```
354df78fc77f241bb89003434d1c27b977208a26
```

Proof:

```bash
$ go run bucketter.go input.txt | python -c "import sys,hashlib; print(hashlib.sha1(sys.stdin.read().strip()).hexdigest())"

354df78fc77f241bb89003434d1c27b977208a26
```
