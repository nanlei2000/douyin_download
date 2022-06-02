package weibo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	imagewall "github.com/nanlei2000/douyin_download/pkg/weibo/model/image_wall"
)

func (w *Weibo) GetAllImageWallPid(uid string) (ImageSet, error) {
	idList := []string{}
	sinceId := "0"

	for {
		url := fmt.Sprintf("https://weibo.com/ajax/profile/getImageWall?uid=%s&sinceid=%s", uid, sinceId)

		fmt.Printf("url: %s\n", url)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return ImageSet{}, err
		}

		err = w.setupHeaders(req, true)
		if err != nil {
			return ImageSet{}, err
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return ImageSet{}, err
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return ImageSet{}, err
		}

		var imageWall imagewall.ImageWall
		err = json.Unmarshal(body, &imageWall)
		if err != nil {
			return ImageSet{}, err
		}

		for _, item := range imageWall.Data.List {
			idList = append(idList, item.PID)
		}

		switch s := imageWall.Data.SinceID.(type) {
		case float64:
			imageWall.Data.SinceID = fmt.Sprintf("%v", s)
		}

		if imageWall.Data.SinceID == "0" {
			break
		}

		sinceId = imageWall.Data.SinceID.(string)
		time.Sleep(100 * time.Millisecond)
	}

	userInfo, err := w.GetUserInfo(uid)
	if err != nil {
		return ImageSet{}, err
	}

	imageSet := ImageSet{
		Name:   userInfo.Data.User.ScreenName,
		IdList: idList,
	}

	return imageSet, nil
}
