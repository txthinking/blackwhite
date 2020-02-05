# A PAC Generator

### Online

#### PAC

- White List Mode with `socks5://127.0.0.1:1080` `https://blackwhite.txthinking.com/white.pac`
- Black List Mode with `socks5://127.0.0.1:1080` `https://blackwhite.txthinking.com/black.pac`
- White List Mode with `http://127.0.0.1:8080` `https://blackwhite.txthinking.com/http_white.pac`
- Black List Mode with `http://127.0.0.1:8080` `https://blackwhite.txthinking.com/http_black.pac`

#### List

- `https://blackwhite.txthinking.com/white.list`
- `https://blackwhite.txthinking.com/white_cidr.list`
- `https://blackwhite.txthinking.com/black.list`
- `https://blackwhite.txthinking.com/black_cidr.list`

### How to update list

- `$ addWhite.sh china.com`
- `$ addWhiteCIDR.sh 1.0.1.0/24`
- `$ addBlack.sh google.com`
- `$ addBlackCIDR.sh 74.125.0.0/16`
- `$ removeWhite.sh china.com`
- `$ removeWhiteCIDR.sh 1.0.1.0/24`
- `$ removeBlack.sh google.com`
- `$ removeBlackCIDR.sh 74.125.0.0/16`

> Don't edit list directly<br/>
> Prefer first-level domain

### How to build PAC file

Require `go`, `nodejs` installed, and run `$ go get`, `$ npm install`

```
# build white.pac
$ go run pac.go \
    -m white \
    -d https://blackwhite.txthinking.com/white.list \
    -c https://blackwhite.txthinking.com/white_cidr.list \
    -p 'SOCKS5 127.0.0.1:1080; SOCKS 127.0.0.1:1080; DIRECT' \
    > white.pac

# build more and compress
$ ./build.sh
```

### How to run PAC server

```
$ go run pac.go \
    -m white \
    -d https://blackwhite.txthinking.com/white.list \
    -c https://blackwhite.txthinking.com/white_cidr.list \
    -p 'SOCKS5 127.0.0.1:1080; SOCKS 127.0.0.1:1080; DIRECT' \
    -l ':1980'

$ curl http://127.0.0.1:1980/proxy.pac
```

### Thanks to

- https://github.com/Leask/Flora_Pac
- https://github.com/breakwa11/gfw_whitelist
- https://github.com/n0wa11/gfw_whitelist
- https://github.com/txthinking/google-hosts
