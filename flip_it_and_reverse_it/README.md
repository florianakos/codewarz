## Network_Awkwardness

Given a file with several words (some lines contain multiple words and some lines only contain one word). Print the words, one per line and reverse every other word following the first word.

Eample input:

```
$ cat /path/to/somefile.txt
this
is
something
something
you
are
used
to
and
nothing
else
```

Expected output:

```
$ ./flip_it_and_reverse_it_solve.py /path/to/somefile.txt
this
si
something
gnihtemos
you
era
used
ot
and
gnihton
else
```

The `SHA1` of the output shall be:
```
 050b6e4c4fb22d39bc92c267627267e10b4e4ec2
```

Proof:

```bash
$ go run pcapper.go input.txt | python -c "import sys,hashlib; print(hashlib.sha1(sys.stdin.read().strip()).hexdigest())"

2643d669feb38df4f151bc0f779defffd8d445fd
```
