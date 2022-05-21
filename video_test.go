package main

import (
	"testing"
)

func TestVideo_Download(t *testing.T) {
	str := `9.23 zTl:/   https://v.douyin.com/FmnpC1V/ 复制此链接，打开Dou音搜索，直接观看视频！`

	dy := NewDouYin()
	dy.isDebug = true
	video, err := dy.Get(str)
	if err != nil {
		t.Fatal(err)
	}
	p, err := video.Download("./video/")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(p)
}
