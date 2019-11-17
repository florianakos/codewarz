## Find_Random_Stuffs challenge

Write a program which that receives two numbers with a space between them and will find all such numbers which are divisible by 7 but are not a multiple of 5, between those two numbers (both included). The numbers obtained should be printed in a comma-separated sequence on a single line.

Sample Input:
```bash
$ cat /path/to/somefile.txt
10 100
50 60
```

The `SHA1` of the output shall be:
```
ffebec9a679f6b31167c99efac9a066da9fe2d4c
```

Proof:

```bash
$ go run main.go input.txt | python -c "import sys,hashlib; print(hashlib.sha1(sys.stdin.read().strip()).hexdigest())"

ffebec9a679f6b31167c99efac9a066da9fe2d4c
```
