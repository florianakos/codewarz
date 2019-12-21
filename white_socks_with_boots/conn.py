#!/usr/bin/env python3
import socket
import sys
with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
    s.connect((sys.argv[1], int(sys.argv[2])))
    data = s.recv(1024)
print(data.decode("utf-8").rstrip())
