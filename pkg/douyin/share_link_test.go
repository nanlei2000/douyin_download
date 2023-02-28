package douyin

import "testing"

func TestGetRedirectUrl(t *testing.T) {
	res, err := getVideoID("https://v.douyin.com/SNwacoU/")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(res)
}
