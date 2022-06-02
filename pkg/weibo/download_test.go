package weibo

import "testing"

func TestDownLoadShowPics(t *testing.T) {
	w := Weibo{}
	err := w.DownLoadShowPics("https://weibo.com/2286073303/LvhJiA6Fh", "../../weibo/")

	if err != nil {
		t.Fatalf("err: %s", err)
	}
}
