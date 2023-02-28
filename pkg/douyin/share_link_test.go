package douyin

import "testing"

func TestGetRedirectUrl(t *testing.T) {
	res, err := getVideoID("https://v.douyin.com/SNwacoU/")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(res)
}

func TestGetVideoIDBySharedLink(t *testing.T) {
	res, err := GetVideoIDBySharedLink(`4.10 tRK:/ 怎么泡男孩子啊，多少水温合适啊%%微胖女生 %%rap版呜呼卡点舞 https://v.douyin.com/F4vTT79/ 复制此链接，打开Dou音搜索，直接观看视频！`)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(res)
}
