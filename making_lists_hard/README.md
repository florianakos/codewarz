## Making Lists Hard

Take the appropriate input file and process it into a list. If the line has a single element, add it as a single element to a new list. If the line contains multiple elements then write those elements of the line into their own list and then add that list to the final list. Finally, someone seems to have added characters and words to the files, be sure to ignore those entries, also we do not care about empty lists...

Example input:

```bash
1a
1
a
2
3
asdasd,a,33
4
5,b,6,7,c,8
9
10
d
11
e
12,13,f,14,15,h,16,17
18
19
 , , ,
20
dsaasdas
i
```

Expected output:

```bash
[1, 2, 3, 33, 4, [5, 6, 7, 8], 9, 10, 11, [12, 13, 14, 15, 16, 17], 18, 19, 20]
```

The `SHA1` of the output shall be:
```
