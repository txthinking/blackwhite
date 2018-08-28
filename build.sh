#!/bin/bash

go run pac.go \
    -m white \
    -d https://blackwhite.txthinking.com/white.list \
    -c https://blackwhite.txthinking.com/white_cidr.list \
    -p 'SOCKS5 127.0.0.1:1080; SOCKS 127.0.0.1:1080; DIRECT' \
    > ${1}white.pac
./node_modules/.bin/uglifyjs --compress --mangle -- ${1}white.pac > _
mv _ ${1}white.pac

go run pac.go \
    -m black \
    -d https://blackwhite.txthinking.com/black.list \
    -c https://blackwhite.txthinking.com/black_cidr.list \
    -p 'SOCKS5 127.0.0.1:1080; SOCKS 127.0.0.1:1080; DIRECT' \
    > ${1}black.pac
./node_modules/.bin/uglifyjs --compress --mangle -- ${1}black.pac > _
mv _ ${1}black.pac

go run pac.go \
    -m global \
    -p 'SOCKS5 127.0.0.1:1080; SOCKS 127.0.0.1:1080; DIRECT' \
    > ${1}global.pac
./node_modules/.bin/uglifyjs --compress --mangle -- ${1}global.pac > _
mv _ ${1}global.pac

go run pac.go \
    -m white \
    -d https://blackwhite.txthinking.com/white.list \
    -c https://blackwhite.txthinking.com/white_cidr.list \
    -p 'PROXY 127.0.0.1:8080; DIRECT' \
    > ${1}http_white.pac
./node_modules/.bin/uglifyjs --compress --mangle -- ${1}http_white.pac > _
mv _ ${1}http_white.pac

go run pac.go \
    -m black \
    -d https://blackwhite.txthinking.com/black.list \
    -c https://blackwhite.txthinking.com/black_cidr.list \
    -p 'PROXY 127.0.0.1:8080; DIRECT' \
    > ${1}http_black.pac
./node_modules/.bin/uglifyjs --compress --mangle -- ${1}http_black.pac > _
mv _ ${1}http_black.pac

go run pac.go \
    -m global \
    -p 'PROXY 127.0.0.1:8080; DIRECT' \
    > ${1}http_global.pac
./node_modules/.bin/uglifyjs --compress --mangle -- ${1}http_global.pac > _
mv _ ${1}http_global.pac

