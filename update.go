//
// https://github.com/txthinking/pac
//
package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/txthinking/ant"
)

var wl map[string]map[string][]string
var bl map[string]map[string][]string
var cm map[string]string
var wc []map[string]int64
var uplock *sync.RWMutex = &sync.RWMutex{}

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
	makeList := func(data []byte) map[string]map[string][]string {
		w := make(map[string]map[string][]string)
		w["_"] = map[string][]string{"_": []string{}}
		bb := bytes.Split(bytes.TrimSpace(data), []byte("\n"))
		for _, d := range bb {
			if net.ParseIP(string(d)) != nil {
				w["_"]["_"] = append(w["_"]["_"], string(d))
				continue
			}

			rd := reverseAsCopy(d)
			i := bytes.IndexByte(rd, '.')
			if i == 0 || i == len(rd)-1 {
				// invalid
				continue
			}
			if i == -1 {
				// cn/local/...
				w["_"]["_"] = append(w["_"]["_"], string(d))
				continue
			}

			suffix := string(rd[:i])
			index := string(rd[i+1 : i+2])
			_, ok := w[suffix]
			if !ok {
				w[suffix] = make(map[string][]string)
			}
			_, ok = w[suffix][index]
			if !ok {
				w[suffix][index] = make([]string, 0)
			}
			w[suffix][index] = append(w[suffix][index], string(d))
		}
		return w
	}
	makeMap := func(data []byte) map[string]string {
		cm := make(map[string]string)
		ss := strings.Split(strings.TrimSpace(bytes.NewBuffer(data).String()), "\n")
		for _, s := range ss {
			tmp := strings.SplitN(s, ":", 2)
			if len(tmp) < 2 {
				continue
			}
			cm[tmp[0]] = tmp[1]
		}
		return cm
	}
	makeWhiteCIDR := func(data []byte) []map[string]int64 {
		wc := make([]map[string]int64, 0)
		ss := strings.Split(strings.TrimSpace(bytes.NewBuffer(data).String()), "\n")
		for _, s := range ss {
			c, err := ant.CIDR(s)
			if err != nil {
				continue
			}
			first, err := ant.IP2Decimal(c.First)
			if err != nil {
				continue
			}
			last, err := ant.IP2Decimal(c.Last)
			if err != nil {
				continue
			}
			m := make(map[string]int64)
			m["first"] = first
			m["last"] = last
			wc = append(wc, m)
		}
		return wc
	}
	makeUpdate := func() {
		uplock.Lock()
		defer uplock.Unlock()
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
		if whiteCIDR != "" {
			if data, err := fetchData(whiteCIDR); err != nil {
				log.Println(err)
			} else {
				wc = makeWhiteCIDR(data)
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

func reverseAsCopy(s []byte) []byte {
	a := make([]byte, len(s))
	copy(a, s)
	i := 0
	j := len(a) - 1
	for i < j {
		x := a[i]
		a[i] = a[j]
		a[j] = x
		i++
		j--
	}
	return a
}
