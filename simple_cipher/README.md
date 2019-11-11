## Simple Cipher

The task is to read an input file and decrypt the message hidden in the sentences... according to the hint given in the challenge description.

The `SHA1` of the output shall be:
```
36749e5e0a717d99a6836dbb25ec6ad42ef77848
```

Proof:

```bash
$ go run simplecipher.go input.txt | python -c "import sys,hashlib; print(hashlib.sha1(sys.stdin.read().strip()).hexdigest())"

36749e5e0a717d99a6836dbb25ec6ad42ef77848
```
