#!/usr/bin/env python
import socket
import sys

s = socket.socket()
s.connect((sys.argv[1], int(sys.argv[2])))
command = s.recv(1024).decode("utf-8").split("\n")[1].replace("!", "").split(" ")
result = None
if command[0] == "add":
    result = int(command[1]) + int(command[3])
    s.sendall(("{:d}".format(result)+"\n\r").encode())
    pass
elif command[0] == "subtract":
    result = int(command[3]) - int(command[1])
    s.sendall(("{:d}".format(result)+"\n\r").encode())
    pass
elif command[0] == "divide":
    result = int(command[1]) / int(command[3])
    s.sendall(("{:.8f}".format(result)+"\n\r").encode())
    pass
elif command[0] == "multiply":
    result = int(command[1]) * int(command[3])
    s.sendall(("{:d}".format(result)+"\n\r").encode())
    pass
print(s.recv(1024).decode("utf-8").rstrip())
s.close()
