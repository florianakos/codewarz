## Forwards_is_backwards challenge

Write a program that receives a single word as input and checks to see if the word is a palindrome (i.e. words that look the same written backwards).

The `SHA1` of the output shall be:
```
2129dd88ed274d9bfd1e3facf9fb252d91103196
```

Proof:

```bash
$ go run main.go input.txt | python -c "import sys,hashlib; print(hashlib.sha1(sys.stdin.read().strip()).hexdigest())"

2129dd88ed274d9bfd1e3facf9fb252d91103196
```
