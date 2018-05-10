//
// https://github.com/txthinking/pac
//
package main

import (
	"io"
	"log"
	"os"

	"github.com/codegangsta/cli"
	"github.com/txthinking/pac/blackwhite"
)

var mode string
var domainURL string
var cidrURL string
var proxy string

func main() {
	app := cli.NewApp()
	app.Name = "PAC"
	app.Version = "20180510"
	app.Usage = "PAC file generator"
	app.Author = "Cloud"
	app.Email = "cloud@txthinking.com"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "mode, m",
			Value:       "white",
			Usage:       "white/black/global",
			Destination: &mode,
		},
		cli.StringFlag{
			Name:        "domainURL, d",
			Value:       "https://www.txthinking.com/pac/white.list",
			Usage:       "domains url, http(s):// or file://",
			Destination: &domainURL,
		},
		cli.StringFlag{
			Name:        "cidrURL, c",
			Value:       "https://www.txthinking.com/pac/white_cidr.list",
			Usage:       "CIDR url, http(s):// or file://",
			Destination: &cidrURL,
		},
		cli.StringFlag{
			Name:        "proxy, p",
			Value:       "SOCKS5 127.0.0.1:1080; SOCKS 127.0.0.1:1080; DIRECT",
			Usage:       "Proxy",
			Destination: &proxy,
		},
	}
	app.Action = func(c *cli.Context) error {
		return run()
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	r, err := blackwhite.PAC(proxy, mode, domainURL, cidrURL)
	if err != nil {
		return err
	}
	if _, err := io.Copy(os.Stdout, r); err != nil {
		return err
	}
	return nil
}
