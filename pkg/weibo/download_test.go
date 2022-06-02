package weibo

import "testing"

func TestDownLoadShowPics(t *testing.T) {
	w := Weibo{}
	err := w.DownLoadShowPics(Source{
		Type: Show,
		Link: "https://weibo.com/2286073303/LvhJiA6Fh",
	}, "../../weibo/")

	if err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestDownLoadImageWallPics(t *testing.T) {
	t.SkipNow()
	w := Weibo{}
	err := w.DownLoadShowPics(Source{
		Type: ImageWall,
		Link: "https://weibo.com/u/2286073303?tabtype=album",
	}, "../../weibo/")

	if err != nil {
		t.Fatalf("err: %s", err)
	}
}
