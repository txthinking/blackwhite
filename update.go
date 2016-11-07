//
// https://github.com/txthinking/pac
//
package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var wl []string
var bl []string
var cm map[string]string

func fetchData(where string) ([]byte, error) {
	if strings.HasPrefix(where, "http://") || strings.HasPrefix(where, "https://") {
		res, err := http.Get(where)
		if err != nil {
			return nil, err
		}
		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		err = res.Body.Close()
		return data, err
	}

	f, err := os.Open(where)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	err = f.Close()
	return data, err
}
func update() {
	makeList := func(data []byte) []string {
		return strings.Split(strings.TrimSpace(bytes.NewBuffer(data).String()), "\n")
	}
	makeMap := func(data []byte) map[string]string {
		cm := make(map[string]string)
		ss := strings.Split(strings.TrimSpace(bytes.NewBuffer(data).String()), "\n")
		for _, s := range ss {
			tmp := strings.SplitN(s, ":", 2)
			if len(tmp) < 2 {
				log.Println("Invalid rule in customized map:", s)
				continue
			}
			cm[tmp[0]] = tmp[1]
		}
		return cm
	}
	makeUpdate := func() {
		if whiteList != "" {
			if data, err := fetchData(whiteList); err != nil {
				log.Println(err)
			} else {
				wl = makeList(data)
			}
		}
		if blackList != "" {
			if data, err := fetchData(blackList); err != nil {
				log.Println(err)
			} else {
				bl = makeList(data)
			}
		}
		if customizeMap != "" {
			if data, err := fetchData(customizeMap); err != nil {
				log.Println(err)
			} else {
				cm = makeMap(data)
			}
		}
	}

	if cycle == 0 {
		makeUpdate()
		return
	}
	go func() {
		for {
			makeUpdate()
			time.Sleep(time.Duration(int64(cycle)) * time.Second)
		}
	}()
}
