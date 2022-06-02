package weibo

import "testing"

func TestGetUserInfo(t *testing.T) {
	w := Weibo{}
	info, err := w.GetUserInfo("2286073303")

	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if info.Data.User.ScreenName != "SULENDOOO" {
		t.Fatalf("wrong result: %#v", info)
	}
}
