package douyin

import "fmt"

func GetPlayAddr(videoID string) (string, error) {
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

	return fmt.Sprintf("https:%s", playAddr), nil
}
