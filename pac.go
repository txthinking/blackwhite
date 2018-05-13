//
// https://github.com/txthinking/pac
//
package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/codegangsta/cli"
	"github.com/txthinking/pac/blackwhite"
)

var mode string
var domainURL string
var cidrURL string
var proxy string
var server string

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
		cli.StringFlag{
			Name:        "server, s",
			Usage:       "HTTP server address, like: 127.0.0.1:1980",
			Destination: &server,
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
	if server == "" {
		if _, err := io.Copy(os.Stdout, r); err != nil {
			return err
		}
		return nil
	}

	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/x-ns-proxy-autoconfig")
		w.Write(b)
	})
	return http.ListenAndServe(server, nil)
}
