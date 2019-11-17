## Hooked on Encoding challenge

The task is to read an input file line by line and decipher the poem hidden between the lines

The `SHA1` of the output shall be:
```
7dd8e58b19dc8eb7a81f053780d2dd1ac7d4d7d1
```

Proof:

```bash
$ go run hooked.go input.txt | python -c "import sys,hashlib; print(hashlib.sha1(sys.stdin.read().strip()).hexdigest())"

7dd8e58b19dc8eb7a81f053780d2dd1ac7d4d7d1
```
