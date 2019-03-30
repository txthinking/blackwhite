//
// https://github.com/txthinking/blackwhite
//
package main

import (
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/codegangsta/cli"
	"github.com/txthinking/blackwhite/pac"
)

var mode string
var domainURL string
var cidrURL string
var proxy string
var listen string

func main() {
	app := cli.NewApp()
	app.Name = "PAC"
	app.Version = "20180918"
	app.Usage = "PAC file generator"
	app.Author = "Cloud"
	app.Email = "cloud@txthinking.com"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "mode, m",
			Usage:       "white/black/global [required]",
			Destination: &mode,
		},
		cli.StringFlag{
			Name:        "domainURL, d",
			Usage:       "domains url, http(s):// or file:// [required when mode is not global]",
			Destination: &domainURL,
		},
		cli.StringFlag{
			Name:        "cidrURL, c",
			Usage:       "CIDR url, http(s):// or file:// [optional]",
			Destination: &cidrURL,
		},
		cli.StringFlag{
			Name:        "proxy, p",
			Usage:       "Proxy, like: SOCKS5 127.0.0.1:1080; SOCKS 127.0.0.1:1080; DIRECT [required]",
			Destination: &proxy,
		},
		cli.StringFlag{
			Name:        "listen, l",
			Usage:       "HTTP server address, like: 127.0.0.1:1980 [optional]",
			Destination: &listen,
		},
	}
	app.Action = func(c *cli.Context) error {
		if mode != "global" && mode != "white" && mode != "black" {
			return errors.New("Invalid mode")
		}
		if mode != "global" && domainURL == "" {
			return errors.New("domainURL required")
		}
		if proxy == "" {
			return errors.New("proxy required")
		}
		return run()
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	r, err := pac.PAC(proxy, mode, domainURL, cidrURL)
	if err != nil {
		return err
	}
	if listen == "" {
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
	return http.ListenAndServe(listen, nil)
}
