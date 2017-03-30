package blackwhite

import "testing"

func TestIsWhite(t *testing.T) {
	ok := isWhite("211.94.114.48")
	t.Log(ok)

	ok = isWhite("www.google.com")
	t.Log(ok)

	ok = isWhite("music.163.com")
	t.Log(ok)

	ok = isWhite("www.txthinking.com")
	t.Log(ok)
}
