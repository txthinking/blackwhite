# A PAC Generator

### Online

#### PAC

* White List Mode with `socks5://127.0.0.1:1080` `https://www.txthinking.com/pac/white.pac`
* Black List Mode with `socks5://127.0.0.1:1080` `https://www.txthinking.com/pac/black.pac`

#### List

* `https://www.txthinking.com/pac/white.list`
* `https://www.txthinking.com/pac/white_cidr.list`
* `https://www.txthinking.com/pac/white_app.list`
* `https://www.txthinking.com/pac/black.list`
* `https://www.txthinking.com/pac/black_cidr.list`
* `https://www.txthinking.com/pac/black_app.list`

### How to update list

* `$ addWhite.sh          china.com`
* `$ addWhiteCIDR.sh      1.0.1.0/24`
* `$ addWhiteApp.sh       com.tencent.mm`
* `$ addBlack.sh          google.com`
* `$ addBlackCIDR.sh      74.125.0.0/16`
* `$ addBlackApp.sh       com.android.chrome`
* `$ removeWhite.sh       china.com`
* `$ removeWhiteCIDR.sh   1.0.1.0/24`
* `$ removeWhiteApp.sh    com.tencent.mm`
* `$ removeBlack.sh       google.com`
* `$ removeBlackCIDR.sh   74.125.0.0/16`
* `$ removeBlackApp.sh    com.android.chrome`

> Don't edit white.list/black.list directly<br/>
> Prefer first-level domain

### How to build PAC file

Require `go`, `nodejs` installed, and `npm install`

```
$ ./build.sh
$ ./compress.sh
```

### How to run PAC server

```
$ go run pac.go \
    -m white \
    -d https://www.txthinking.com/pac/white.list \
    -c https://www.txthinking.com/pac/white_cidr.list \
    -p 'SOCKS5 127.0.0.1:1080; SOCKS 127.0.0.1:1080; DIRECT' \
    -l ':1980'

$ curl http://127.0.0.1:1980/proxy.pac
```

### Thanks to

* https://github.com/Leask/Flora_Pac
* https://github.com/breakwa11/gfw_whitelist
* https://github.com/n0wa11/gfw_whitelist
* https://github.com/txthinking/google-hosts
