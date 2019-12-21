## WhiteSocksAndBoots

Connect to the server, it'll give you a flag!
Test server located at listen.runcode.ninja:9005

Example input:

```bash
    $ ./white_socks_with_boots_solve.py listen.runcode.ninja 9005
```

Expected output:

```bash
$ ./white_socks_with_boots_solve.py listen.runcode.ninja 9005
RCN{670_dash_one_is_better_than_350_dash_one}
```

The `SHA1` of the output shall be:
```
fc03b75227628f4fab5f3bfe581efe0443082d3a
```

Proof:

```bash
$ python conn.py listen.runcode.ninja:9005 | python -c "import sys,hashlib; print(hashlib.sha1(sys.stdin.read().strip()).hexdigest())"

fc03b75227628f4fab5f3bfe581efe0443082d3a
```
