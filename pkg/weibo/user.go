package weibo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/nanlei2000/douyin_download/pkg/weibo/model/info"
)

func (w *Weibo) setupHeaders(req *http.Request, needCookie bool) error {
	req.Header.Set("Authority", "weibo.com")
	req.Header.Set("Accept", "application/json, text/plain, */*")

	if needCookie {
		cookie := os.Getenv("WB_COOKIE")
		if len(cookie) == 0 {
			return errors.New("env WB_COOKIE is required")
		}
		req.Header.Set("Cookie", cookie)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.64 Safari/537.36 Edg/101.0.1210.53")

	return nil
}

func (w *Weibo) GetUserInfo(uid string) (*info.UserInfo, error) {
	url := fmt.Sprintf("https://weibo.com/ajax/profile/info?uid=%s", uid)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	err = w.setupHeaders(req, true)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var userInfo info.UserInfo
	err = json.Unmarshal(body, &userInfo)
	if err != nil {
		return nil, err
	}

	return &userInfo, nil
}
