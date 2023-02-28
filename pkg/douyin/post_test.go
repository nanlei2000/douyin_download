package douyin

import (
	"net/url"
	"testing"
)

func TestGetAllVideoIDList(t *testing.T) {
	defer Browser.Close()

	list, err := GetAllVideoIDList("https://www.douyin.com/user/MS4wLjABAAAApISbVdbu_OzxCWeq75MzzWcqtd8sSa8hNUvizQpvIBg")

	if err != nil {
		t.Fatalf("err: %s", err)
	}

	if len(list) == 0 {
		t.Fatalf("get all id failed")
	}
	t.Logf("list: %s", list)
}

func TestParseURL(t *testing.T) {
	url, err := url.Parse("https://www.douyin.com/user/MS4wLjABAAAAZimxk0o3KWTEJNNrzwSF3HBjCy4TkS6mpPyHNxEYC2A?relation=1")
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	t.Logf("%#v", url)
}
