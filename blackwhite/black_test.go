package blackwhite

import (
	"testing"
)

func TestIsBlack(t *testing.T) {
	if !IsBlack("67.220.91.15") {
		t.Fatal("Failed")
	}
	if !IsBlack("plus.google.com") {
		t.Fatal("Failed")
	}
	if IsBlack("a.cn") {
		t.Fatal("Failed")
	}
	if IsBlack("www.txthinking.com") {
		t.Fatal("Failed")
	}
}
