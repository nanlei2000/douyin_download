package douyin

import (
	"net/url"
	"testing"
)

func TestGetAllVideoIDList(t *testing.T) {
	dy := NewDouYin()
	dy.isDebug = true

	list, err := dy.GetAllVideoIDList("MS4wLjABAAAAF0HqDm-8U9TiT_9AfqqPGiNbrP0c93AdB3_oRG7Em_Q")

	if err != nil {
		t.Fatalf("err: %s", err)
	}

	if len(list) != 73 {
		t.Fatalf("get all id failed, expect: 73, actual: %d", len(list))
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
