# A PAC Generator

### The Online PAC

* White List Mode with `socks5://127.0.0.1:1080` `https://www.txthinking.com/pac/white.pac`
* Black List Mode with `socks5://127.0.0.1:1080` `https://www.txthinking.com/pac/black.pac`

### How to update white/black list

* `$ addWhite.sh [domains]`. ep: `$ addWhite.sh china.com`
* `$ removeWhite.sh [domains]`. ep: `$ removeWhite.sh china.com`
* `$ addBlack.sh [domains]`. ep: `$ addBlack.sh google.com`
* `$ removeBlack.sh [domains]`. ep: `$ removeBlack.sh google.com`
* Don't edit white.list/black.list directly
* Prefer first-level domain

### How to build PAC file

```
$ ./build.sh
$ ./compress.sh
```

### Thanks to

* https://github.com/Leask/Flora_Pac
* https://github.com/breakwa11/gfw_whitelist
* https://github.com/n0wa11/gfw_whitelist
* https://github.com/txthinking/google-hosts
