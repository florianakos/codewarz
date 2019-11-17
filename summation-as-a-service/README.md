## Summation As A Service challenge

Write a program, which takes two distinct integers separated by space as input and prints the sum of all the integers between them, including the two given numbers. Note that the numbers can appear in either order.

The `SHA1` of the output shall be:
```
4d5ba6ab19fa9a29c16d5b2fbfb4842a38a842cf
```

Proof:

```bash
$ go run simplecipher.go input.txt | python -c "import sys,hashlib; print(hashlib.sha1(sys.stdin.read().strip()).hexdigest())"

4d5ba6ab19fa9a29c16d5b2fbfb4842a38a842cf
```
