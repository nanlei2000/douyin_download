package douyin

import (
	"fmt"
	"strings"

	"github.com/go-rod/rod"
)

func GetHTMLRod() {
	browser := rod.New().MustConnect()
	defer browser.MustClose()

	page := browser.MustPage("https://www.douyin.com/user/MS4wLjABAAAA6HYh8jBC4Z-32mK0eFvxC1ibpS4RiucDZ9k_RxuHmBQ?vid=7198206406237900084")

	// Start to analyze request events
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
		fmt.Println(len(page.MustElements("li")))
		if strings.Contains(html, "暂时没有更多了") {
			break
		}

	}

	target := page.MustElements(`a[href^="/video/"]`)

	ids := []string{}
	for _, el := range target {
		id := el.MustAttribute("href")
		if id != nil {
			ids = append(ids, *id)
		}
	}

	fmt.Printf("ids: %s ,len: %d\n", ids, len(ids))
}
