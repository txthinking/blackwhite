package pac

import (
	"io/ioutil"
	"log"
	"testing"
)

func TestTest(t *testing.T) {
	r, _ := PACFromString("SOCKS5 127.0.0.1:1080; SOCKS 127.0.0.1:1080; DIRECT", "white", "", "")
	b, _ := ioutil.ReadAll(r)
	log.Println(string(b))
}
