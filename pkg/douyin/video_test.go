package douyin

import (
	"testing"
)

func TestVideo_Download(t *testing.T) {
	dy := NewDouYin()
	dy.isDebug = true
	video, err := dy.Get(Source{
		Type:    SourceType_ShardContent,
		Content: `4.10 tRK:/ 怎么泡男孩子啊，多少水温合适啊%%微胖女生 %%rap版呜呼卡点舞  https://v.douyin.com/F4vTT79/ 复制此链接，打开Dou音搜索，直接观看视频！`,
	})
	if err != nil {
		t.Fatal(err)
	}
	p, err := video.Download("./video/")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(p)
}

func TestVideo_Download_FromID(t *testing.T) {
	dy := NewDouYin()
	dy.isDebug = true
	video, err := dy.Get(Source{
		Type:    SourceType_VideoID,
		Content: `7015482603633511683`,
	})
	if err != nil {
		t.Fatal(err)
	}
	p, err := video.Download("./video/")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(p)
}
