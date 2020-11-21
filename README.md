# Black & White List

Moved to [bypass](https://github.com/txthinking/bypass)

#### List

- `https://txthinking.github.io/blackwhite/white.list`
- `https://txthinking.github.io/blackwhite/white_cidr.list`
- `https://txthinking.github.io/blackwhite/black.list`
- `https://txthinking.github.io/blackwhite/black_cidr.list`

### Update list

```
$ cd tools
$ ./addWhite.sh china.com
$ ./addWhiteCIDR.sh 1.0.1.0/24
$ ./addBlack.sh google.com
$ ./addBlackCIDR.sh 74.125.0.0/16
$ ./removeWhite.sh china.com
$ ./removeWhiteCIDR.sh 1.0.1.0/24
$ ./removeBlack.sh google.com
$ ./removeBlackCIDR.sh 74.125.0.0/16
```

> Prefer first-level domain

#### Online PAC with `socks5://127.0.0.1:1080`

- `https://txthinking.github.io/blackwhite/white.pac`
- `https://txthinking.github.io/blackwhite/black.pac`

> You can create PAC by yourself with [$ brook pac ...](https://github.com/txthinking/brook)

### Thanks to

- https://github.com/Leask/Flora_Pac
- https://github.com/breakwa11/gfw_whitelist
- https://github.com/n0wa11/gfw_whitelist
- https://github.com/txthinking/google-hosts
