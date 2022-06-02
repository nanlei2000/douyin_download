package weibo

import "testing"

func TestGetShowPics(t *testing.T) {
	w := Weibo{}
	pics, err := w.GetShowPics("LvhJiA6Fh")

	if err != nil {
		t.Fatalf("err: %s", err)
	}
	t.Logf("pics: %s", pics)
}
