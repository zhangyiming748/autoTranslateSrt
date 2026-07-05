package trans

import "testing"

func TestTranslate(t *testing.T) {
	dst := Translate("hello")
	t.Log(dst)
}
