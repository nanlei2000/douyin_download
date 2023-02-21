package douyin

import (
	"fmt"
	"strings"

	"github.com/go-rod/rod"
)

func (d *DouYin) GetAllVideoIDList(userLink string) ([]string, error) {
	browser := rod.New().MustConnect()
	defer browser.MustClose()

	page := browser.MustPage(userLink)
	wait := page.MustWaitRequestIdle()
	wait()

	if page.MustElement(".dy-account-close") != nil {
		page.MustElement(".dy-account-close").MustClick()
	}

	for {
		wait := page.MustWaitRequestIdle()
		page.MustElements("li").Last().ScrollIntoView()
		wait()
		html, _ := page.HTML()
		if strings.Contains(html, "暂时没有更多了") {
			break
		}

	}

	target := page.MustElements(`a[href^="/video/"]`)

	ids := []string{}
	for _, el := range target {
		href := el.MustAttribute("href")
		if href == nil {
			continue
		}
		parts := strings.Split(*href, "/")
		ids = append(ids, parts[len(parts)-1])
	}

	fmt.Printf("主页：%s，拉取到的视频数：%d\n", userLink, len(ids))

	return ids, nil
}
