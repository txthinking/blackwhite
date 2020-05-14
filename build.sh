#!/bin/bash
brook pac -p 'SOCKS5 127.0.0.1:1080; SOCKS 127.0.0.1:1080; DIRECT' -m white -d ./white.list -c ./white_cidr.list -f ./white.pac
brook pac -p 'SOCKS5 127.0.0.1:1080; SOCKS 127.0.0.1:1080; DIRECT' -m black -d ./black.list -c ./black_cidr.list -f ./black.pac
