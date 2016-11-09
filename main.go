//
// https://github.com/txthinking/pac
//
package main

import (
	"log"
	"net/http"
	"os"

	"github.com/codegangsta/cli"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

var listen string
var whiteList string
var blackList string
var customizeMap string
var cycle int

func main() {
	app := cli.NewApp()
	app.Name = "PAC"
	app.Usage = "A smart PAC file"
	app.Author = "Cloud"
	app.Email = "cloud@txthinking.com"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "listen",
			Value:       ":1901",
			Usage:       "Listen address.",
			Destination: &listen,
		},
		cli.StringFlag{
			Name:        "white",
			Value:       "https://raw.githubusercontent.com/txthinking/pac/master/white.list",
			Usage:       "White list file path or http link.",
			Destination: &whiteList,
		},
		cli.StringFlag{
			Name:        "black",
			Value:       "https://raw.githubusercontent.com/txthinking/pac/master/black.list",
			Usage:       "Black list file path or http link.",
			Destination: &blackList,
		},
		cli.StringFlag{
			Name:        "customize",
			Value:       "https://raw.githubusercontent.com/txthinking/pac/master/customize.map",
			Usage:       "Customized map file path or http link.",
			Destination: &customizeMap,
		},
		cli.IntFlag{
			Name:        "cycle",
			Value:       0,
			Usage:       "Cycle time(s) for updating white list and/or black list and/or customized map from the source.",
			Destination: &cycle,
		},
	}
	app.Action = func(c *cli.Context) error {
		update()
		return run()
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}

func run() error {
	r := mux.NewRouter()
	r.Methods("GET").Path("/{mode}/{proxy}").HandlerFunc(pac)
	r.Methods("GET").Path("/{mode}/{proxy}/pac.pac").HandlerFunc(pac)

	n := negroni.New()
	n.Use(negroni.NewRecovery())
	n.Use(negroni.NewLogger())
	n.UseHandler(r)

	return http.ListenAndServe(listen, n)
}
