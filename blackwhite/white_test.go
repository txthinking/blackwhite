package blackwhite

import (
	"strings"
	"testing"
)

func TestIsWhite(t *testing.T) {
	if !IsWhite("211.94.114.48") {
		t.Fatal("Failed")
	}

	if IsWhite("www.google.com") {
		t.Fatal("Failed")
	}

	if !IsWhite("music.163.com") {
		t.Fatal("Failed")
	}

	if !IsWhite("www.txthinking.com") {
		t.Fatal("Failed")
	}

	if !IsWhite("121.196.205.97") {
		t.Fatal("Failed")
	}

	if !IsWhite("a.cn") {
		t.Fatal("Failed")
	}

	if !IsWhite("txthinking.com") {
		t.Fatal("Failed")
	}
}

func TestGetWhiteAPP(t *testing.T) {
	s := string(strings.TrimSpace(white_app))
	ss := strings.Split(s, "\n")
	t.Log(":" + ss[0] + ":")
	t.Log(":" + ss[len(ss)-1] + ":")
}
