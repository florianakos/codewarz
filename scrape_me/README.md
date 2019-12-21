## Strange_addition

Get the content of the given web page and print it!
Test server located at http://scrape-me.runcode.ninja/

Example input:

```bash
    $ ./scrape_me_solve.py http://scrape-me.runcode.ninja/
```

Expected output:

```bash
$ ./scrape_me_solve.py http://scrape-me.runcode.ninja/
<html>
<head>
<title>Scrape_me</title>
</head>
<body>
<h1>Scrape me, please</h1>
<p>lorem ipsum yadda yadda...</p>
</body>
</html>
```

The `SHA1` of the output shall be:
```
2758dc7a649c4a1f7e8a139c19128cc32efb5817
```

Proof:

```bash
$ python scrape.py input.txt | python -c "import sys,hashlib; print(hashlib.sha1(sys.stdin.read().strip()).hexdigest())"

2758dc7a649c4a1f7e8a139c19128cc32efb5817
```
