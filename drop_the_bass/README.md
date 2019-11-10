## Drop the bass challenge

The task is to read an input file line by line and try to decipher the english word which is hidden behind each line. Encoding may have been applied multiple times...

The `SHA1` of the output shall be:
```
9acad67e0f742d92e4c6a0fa60eb6a9319ee5ab2
```

Proof:

```bash
$ go run main.go input.txt | python -c "import sys,hashlib; print(hashlib.sha1(sys.stdin.read().strip()).hexdigest())"

9acad67e0f742d92e4c6a0fa60eb6a9319ee5ab2
```
