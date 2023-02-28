package douyin

import (
	"fmt"
	"net/url"
	"strings"
)

// TODO: 使用 race api
func GetVideoDetail(videoID string) (*Video, error) {
	page := Browser.MustIncognito().MustPage(fmt.Sprintf("https://www.douyin.com/video/%s", videoID))
	wait := page.MustWaitRequestIdle()
	wait()

	var playAddr string
	srcList := page.MustElements("xg-video-container source")
	for _, src := range srcList {
		link := src.MustAttribute("src")
		if link != nil {
			playAddr = *link
		}
	}

	if len(playAddr) == 0 {
		return nil, fmt.Errorf("获取链接失败")
	}
	url, err := url.Parse(playAddr)
	if err != nil {
		return nil, err
	}

	url.Scheme = "https"
	playAddr = url.String()

	v := Video{
		VideoId:  videoID,
		PlayAddr: playAddr,
		Author: struct {
			SecUid   string
			Nickname string
		}{},
	}

	anchor := page.MustElement(`div[data-e2e="user-info"] > div:nth-child(2)> a`)
	if anchor == nil {
		return nil, fmt.Errorf("获取anchor失败")
	}
	mainPage := anchor.MustAttribute("href")
	if mainPage == nil {
		return nil, fmt.Errorf("获取mainPage失败")
	}
	parts := strings.Split(*mainPage, "/")
	if len(parts) == 0 {
		return nil, fmt.Errorf("mainPage href 格式不对")
	}
	v.Author.SecUid = parts[len(parts)-1]

	nickname := anchor.MustText()
	v.Author.Nickname = nickname

	return &v, nil
}
