## ASCII Stuffs challenge

Given a bunch of hexadecimal digits, return a string that is made from their ASCII characters.

The `SHA1` of the output shall be:
```
21c720c90dfe4faf21bed751e475147c46150662
```

Proof:

```bash
$ go run main.go input.txt | python -c "import sys,hashlib; print(hashlib.sha1(sys.stdin.read().strip()).hexdigest())"

21c720c90dfe4faf21bed751e475147c46150662
```
