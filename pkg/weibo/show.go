package weibo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/nanlei2000/douyin_download/pkg/weibo/model"
)

type Weibo struct {
}

type ImageSet struct {
	Name   string
	IdList []string
}

func (w *Weibo) GetShowPics(id string) (ImageSet, error) {
	url := fmt.Sprintf("https://weibo.com/ajax/statuses/show?id=%s", id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ImageSet{}, err
	}

	err = w.setupHeaders(req, false)
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
	var show model.Show

	err = json.Unmarshal(body, &show)
	if err != nil {
		return ImageSet{}, err
	}

	imageSet := ImageSet{
		Name:   show.User.ScreenName,
		IdList: show.PicIDS,
	}

	return imageSet, nil
}
