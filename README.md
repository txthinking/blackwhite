# A PAC Generator [![Financial Contributors on Open Collective](https://opencollective.com/blackwhite/all/badge.svg?label=financial+contributors)](https://opencollective.com/blackwhite)

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

## Financial Contributors

Become a financial contributor and help us sustain our community. \[[Contribute](https://opencollective.com/blackwhite/contribute)]

#### Individuals

<a href="https://opencollective.com/blackwhite"><img src="https://opencollective.com/blackwhite/individuals.svg?width=890"></a>

#### Organizations

Support this project with your organization. Your logo will show up here with a link to your website. \[[Contribute](https://opencollective.com/blackwhite/contribute)]

<a href="https://opencollective.com/blackwhite/organization/0/website"><img src="https://opencollective.com/blackwhite/organization/0/avatar.svg"></a>
<a href="https://opencollective.com/blackwhite/organization/1/website"><img src="https://opencollective.com/blackwhite/organization/1/avatar.svg"></a>
<a href="https://opencollective.com/blackwhite/organization/2/website"><img src="https://opencollective.com/blackwhite/organization/2/avatar.svg"></a>
<a href="https://opencollective.com/blackwhite/organization/3/website"><img src="https://opencollective.com/blackwhite/organization/3/avatar.svg"></a>
<a href="https://opencollective.com/blackwhite/organization/4/website"><img src="https://opencollective.com/blackwhite/organization/4/avatar.svg"></a>
<a href="https://opencollective.com/blackwhite/organization/5/website"><img src="https://opencollective.com/blackwhite/organization/5/avatar.svg"></a>
<a href="https://opencollective.com/blackwhite/organization/6/website"><img src="https://opencollective.com/blackwhite/organization/6/avatar.svg"></a>
<a href="https://opencollective.com/blackwhite/organization/7/website"><img src="https://opencollective.com/blackwhite/organization/7/avatar.svg"></a>
<a href="https://opencollective.com/blackwhite/organization/8/website"><img src="https://opencollective.com/blackwhite/organization/8/avatar.svg"></a>
<a href="https://opencollective.com/blackwhite/organization/9/website"><img src="https://opencollective.com/blackwhite/organization/9/avatar.svg"></a>

### Thanks to

- https://github.com/Leask/Flora_Pac
- https://github.com/breakwa11/gfw_whitelist
- https://github.com/n0wa11/gfw_whitelist
- https://github.com/txthinking/google-hosts
