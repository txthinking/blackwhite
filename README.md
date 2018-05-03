# A PAC Generator

### The Online PAC

* White List Mode with socks5://127.0.0.1:1080
    `https://www.txthinking.com/pac/white.pac`
* Black List Mode with socks5://127.0.0.1:1080
    `https://www.txthinking.com/pac/black.pac`

### How to update white/black list

* **Don't edit white.list/black.list directly**
* Use the `*.sh` to update the white.list/black.list:
* `$ addWhite.sh [domains]`. ep: `$ addWhite.sh china.com`
* `$ removeWhite.sh [domains]`. ep: `$ removeWhite.sh china.com`
* `$ addBlack.sh [domains]`. ep: `$ addBlack.sh google.com`
* `$ removeBlack.sh [domains]`. ep: `$ removeBlack.sh google.com`
* Domain: prefer first-level domain.

### How to build pac file

```
NAME:
   PAC - A smart PAC file

USAGE:
   build [global options] command [command options] [arguments...]

VERSION:
   20180503

AUTHOR:
   Cloud <cloud@txthinking.com>

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --whitelist value  White list file. (default: "./white.list")
   --whitecidr value  White CIDR file (default: "./china_cidr.list")
   --blacklist value  Black list file (default: "./black.list")
   --proxy value      Proxy (default: "SOCKS5 127.0.0.1:1080; SOCKS 127.0.0.1:1080")
   --help, -h         show help
   --version, -v      print the version
```

### Thanks to

* https://github.com/Leask/Flora_Pac
* https://github.com/breakwa11/gfw_whitelist
* https://github.com/n0wa11/gfw_whitelist
* https://github.com/txthinking/google-hosts
