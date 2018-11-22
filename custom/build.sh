#!/bin/bash

go run ../pac.go \
    -m white \
    -d file://$(pwd)/../white.list \
    -c file://$(pwd)/../white_cidr.list \
    -p 'SOCKS5 127.0.0.1:1090; SOCKS 127.0.0.1:1090; DIRECT' \
    > white.pac
../node_modules/.bin/uglifyjs --compress --mangle -- white.pac > _
mv _ white.pac

go run ../pac.go \
    -m black \
    -d file://$(pwd)/../black.list \
    -c file://$(pwd)/../black_cidr.list \
    -p 'SOCKS5 127.0.0.1:1090; SOCKS 127.0.0.1:1090; DIRECT' \
    > black.pac
../node_modules/.bin/uglifyjs --compress --mangle -- black.pac > _
mv _ black.pac

go run ../pac.go \
    -m global \
    -p 'SOCKS5 127.0.0.1:1090; SOCKS 127.0.0.1:1090; DIRECT' \
    > global.pac
../node_modules/.bin/uglifyjs --compress --mangle -- global.pac > _
mv _ global.pac

go run ../pac.go \
    -m white \
    -d file://$(pwd)/../white.list \
    -c file://$(pwd)/../white_cidr.list \
    -p 'SOCKS5 [::1]:1090; SOCKS [::1]:1090; DIRECT' \
    > white.pac
../node_modules/.bin/uglifyjs --compress --mangle -- white.pac > _
mv _ white6.pac

go run ../pac.go \
    -m black \
    -d file://$(pwd)/../black.list \
    -c file://$(pwd)/../black_cidr.list \
    -p 'SOCKS5 [::1]:1090; SOCKS [::1]:1090; DIRECT' \
    > black.pac
../node_modules/.bin/uglifyjs --compress --mangle -- black.pac > _
mv _ black6.pac

go run ../pac.go \
    -m global \
    -p 'SOCKS5 [::1]:1090; SOCKS [::1]:1090; DIRECT' \
    > global.pac
../node_modules/.bin/uglifyjs --compress --mangle -- global.pac > _
mv _ global6.pac
