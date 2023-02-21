package douyin

import (
	"os"
	"testing"
)

func TestVideo_Download(t *testing.T) {
	dy := NewDouYin()
	dy.isDebug = false
	video, err := dy.Get(Source{
		Type:    SourceType_ShardContent,
		Content: `8.97 tRK:/ 这条裙子是有些洋娃娃的感觉在身上的！%%穿搭 %%甜妹 %%腿 %%少女感  https://v.douyin.com/Y8dCdFC/ 复制此链接，打开Dou音搜索，直接观看视频！`,
	})
	if err != nil {
		t.Fatal(err)
	}
	p, err := video.Download("../../video/")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(p); os.IsNotExist(err) {
		t.Fatalf("download file failed, path: %s", p)
	}
	t.Log(p)
}

func TestImage_Download(t *testing.T) {
	dy := NewDouYin()
	dy.isDebug = true
	video, err := dy.Get(Source{
		Type:    SourceType_ShardContent,
		Content: `0.79 cNj:/ %%这座城市   https://v.douyin.com/FTdTfDw/ 复制此链接，打开Dou音搜索，直接观看视频！`,
	})
	if err != nil {
		t.Fatal(err)
	}
	p, err := video.Download("../../video/")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(p); os.IsNotExist(err) {
		t.Fatalf("download file failed, path: %s", p)
	}
	t.Log(p)
}

func TestVideo_Download_FromID(t *testing.T) {
	dy := NewDouYin()
	dy.isDebug = true
	video, err := dy.Get(Source{
		Type:    SourceType_VideoID,
		Content: `7015873671797411111`,
	})
	if err != nil {
		t.Fatal(err)
	}
	p, err := video.Download("../../video/")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(p); os.IsNotExist(err) {
		t.Fatalf("download file failed, path: %s", p)
	}
	t.Log(p)
}
