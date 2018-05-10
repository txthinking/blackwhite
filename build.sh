#!/bin/bash

go run pac.go \
    -m white \
    -d https://www.txthinking.com/pac/white.list \
    -c https://www.txthinking.com/pac/white_cidr.list \
    -p 'SOCKS5 127.0.0.1:1080; SOCKS 127.0.0.1:1080; DIRECT' \
    > white.pac

go run pac.go \
    -m black \
    -d https://www.txthinking.com/pac/black.list \
    -c "" \
    -p 'SOCKS5 127.0.0.1:1080; SOCKS 127.0.0.1:1080; DIRECT' \
    > black.pac

go run pac.go \
    -m global \
    -d "" \
    -c "" \
    -p 'SOCKS5 127.0.0.1:1080; SOCKS 127.0.0.1:1080; DIRECT' \
    > global.pac
