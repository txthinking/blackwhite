#!/bin/bash

go run pac.go \
    -m white \
    -d https://www.txthinking.com/pac/white.list \
    -c https://www.txthinking.com/pac/white_cidr.list \
    -p 'SOCKS5 127.0.0.1:1080; SOCKS 127.0.0.1:1080; DIRECT' \
    > white.pac
./node_modules/.bin/uglifyjs --compress --mangle -- white.pac > _
mv _ white.pac

go run pac.go \
    -m black \
    -d https://www.txthinking.com/pac/black.list \
    -c https://www.txthinking.com/pac/black_cidr.list \
    -p 'SOCKS5 127.0.0.1:1080; SOCKS 127.0.0.1:1080; DIRECT' \
    > black.pac
./node_modules/.bin/uglifyjs --compress --mangle -- black.pac > _
mv _ black.pac

go run pac.go \
    -m global \
    -p 'SOCKS5 127.0.0.1:1080; SOCKS 127.0.0.1:1080; DIRECT' \
    > global.pac
./node_modules/.bin/uglifyjs --compress --mangle -- global.pac > _
mv _ global.pac

