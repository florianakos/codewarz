#!/usr/bin/env python3
import sys
import requests
print(requests.get(sys.argv[1]).text)
