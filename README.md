## A smart PAC for kiss freedom and more.

### How to use

#### PAC: `https://pac.txthinking.com/:MODE/:PROXY`

* `MODE` is one of `white`/`black`/`all`
* `PROXY` can be this: (replace 127.0.0.1:8888 with your own proxy server address)

    * `DIRECT` for no proxy
    * `PROXY 127.0.0.1:8888` for http proxy
    * `SOCKS5 127.0.0.1:8888; SOCKS 127.0.0.1:8888` for socks5 proxy

#### Example:

* White list mode: `https://pac.txthinking.com/white/SOCKS5%20127.0.0.1:1080;%20SOCKS%20127.0.0.1:1080`
* Black list mode: `https://pac.txthinking.com/black/SOCKS5%20127.0.0.1:1080;%20SOCKS%20127.0.0.1:1080`
* All mode: `https://pac.txthinking.com/all/PROXY%20127.0.0.1:8118`

## Developer

### Install

```
$ go get github.com/txthinking/pac
```

### PAC USAGE

```
[USAGE]:
   $ PAC [OPTIONS]

[OPTIONS]
    --listen value       Listen address. (default: ":1901")
    --white value        White list file path or http link. (default: "https://raw.githubusercontent.com/txthinking/pac/master/white.list")
    --black value        Black list file path or http link. (default: "https://raw.githubusercontent.com/txthinking/pac/master/black.list")
    --customize value    Customized map file path or http link. (default: "https://raw.githubusercontent.com/txthinking/pac/master/customize.map")
    --cycle value        Cycle time(s) for updating white list and/or black list and/or customized map from the source. (default: 0)
    --help, -h           show help
    --version, -v        print the version
```

### How to update white/black list
* **Don't edit white.list/black.list directly**
* Use the `*.sh` to update the white.list/black.list:
* $ addWhite.sh [First-Level Domains]. ep: $addWhite.sh gov.cn
* $ removeWhite.sh [First-Level Domains]. ep: $ removeWhite.sh gov.cn
* $ addBlack.sh [First-Level Domains]. ep: $ addBlack.sh google.com
* $ removeBlack.sh [First-Level Domains]. ep: $ removeBlack.sh google.com

