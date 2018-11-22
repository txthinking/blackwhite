#!/bin/bash

go run pac.go \
    -m white \
    -d file://$(pwd)/white.list \
    -c file://$(pwd)/white_cidr.list \
    -p 'SOCKS5 127.0.0.1:1080; SOCKS 127.0.0.1:1080; DIRECT' \
    > white.pac
./node_modules/.bin/uglifyjs --compress --mangle -- white.pac > _
mv _ white.pac

go run pac.go \
    -m black \
    -d file://$(pwd)/black.list \
    -c file://$(pwd)/black_cidr.list \
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

go run pac.go \
    -m white \
    -d file://$(pwd)/white.list \
    -c file://$(pwd)/white_cidr.list \
    -p 'PROXY 127.0.0.1:8080; DIRECT' \
    > http_white.pac
./node_modules/.bin/uglifyjs --compress --mangle -- http_white.pac > _
mv _ http_white.pac

go run pac.go \
    -m black \
    -d file://$(pwd)/black.list \
    -c file://$(pwd)/black_cidr.list \
    -p 'PROXY 127.0.0.1:8080; DIRECT' \
    > http_black.pac
./node_modules/.bin/uglifyjs --compress --mangle -- http_black.pac > _
mv _ http_black.pac

go run pac.go \
    -m global \
    -p 'PROXY 127.0.0.1:8080; DIRECT' \
    > http_global.pac
./node_modules/.bin/uglifyjs --compress --mangle -- http_global.pac > _
mv _ http_global.pac

