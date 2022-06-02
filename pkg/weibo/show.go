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

func (w *Weibo) GetShowPics(id string) ([]string, error) {
	url := fmt.Sprintf("https://weibo.com/ajax/statuses/show?id=%s", id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []string{}, err
	}
	req.Header.Set("Authority", "weibo.com")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Sec-Ch-Ua", "\" Not A;Brand\";v=\"99\", \"Chromium\";v=\"101\", \"Microsoft Edge\";v=\"101\"")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("Sec-Ch-Ua-Platform", "\"Windows\"")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Traceparent", "00-6448eef8400983a05cac2bd2efc10f5e-6b5eb6f23c3e99d8-00")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.64 Safari/537.36 Edg/101.0.1210.53")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("X-Xsrf-Token", "OjynhPgsETyZ5lYNxY3wwom9")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return []string{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []string{}, err
	}
	var show model.Show

	err = json.Unmarshal(body, &show)
	if err != nil {
		return []string{}, err
	}

	pics := []string{}

	for _, p := range show.PicInfos {
		pics = append(pics, p.Largest.URL)
	}

	return pics, nil
}
