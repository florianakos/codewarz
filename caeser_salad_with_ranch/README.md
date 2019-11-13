## Caeser_salad_with_ranch challenge

Use caesar encryption to decrypt the payload!

Sample input:

```
This jt vjg iluvw wirxirgi.
This jt vjg vhfrqg.
And uijt ku wkh xlmvh.
This tfoufodf jcu d jia rtwj cuxjy aoha apwctl anjuuh rovz jzf agf!
```


The `SHA1` of the output shall be:
```
badb570520cce1d410d901e53d3eb5079b02a15d
```

Proof:

```bash
$ go run main.go input.txt | python -c "import sys,hashlib; print(hashlib.sha1(sys.stdin.read().strip()).hexdigest())"


```
