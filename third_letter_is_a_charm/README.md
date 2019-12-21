## Third_letter_is_a_charm

Simply add the numbers together on each line, also adding how many numbers where on each line.

Example input:

```bash
$ cat /path/to/somefile.txt
eElctricity! nCostitution! iPhladelphia!
sFih! nPoy! pHi, pHi pHo, pHi pHo oannymous? rDan uyo! uYo vgae mhi eth seay eons.
```

Expected output:

```bash
$ ./third_letter_is_a_charm_solve.py /path/to/somefile.txt
Electricity! Constitution! Philadelphia!
Fish! Pony! Hip, Hip Hop, Hip Hop anonymous? Darn you! You gave him the easy ones.
```

The `SHA1` of the output shall be:
```
40ee04887d370995c55eda0c91175738178311e4
```

Proof:

```bash
$ python solve.py input.txt | python -c "import sys,hashlib; print(hashlib.sha1(sys.stdin.read().strip()).hexdigest())"

40ee04887d370995c55eda0c91175738178311e4
```
