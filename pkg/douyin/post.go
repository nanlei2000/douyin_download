package douyin

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/nanlei2000/douyin_download/pkg/model"
)

func (d *DouYin) GetAllVideoIDList(secUid string) ([]string, error) {
	init := true
	cursor := 0
	count := 35
	idList := []string{}

	for {
		if cursor == 0 && !init {
			break
		}

		//  'https://www.iesdouyin.com/aweme/v1/web/aweme/post/?sec_user_id=%s&count=35&max_cursor=0&aid=1128&version_name=23.5.0&device_platform=android&os_version=2333'
		apiURL := fmt.Sprintf("https://www.iesdouyin.com/aweme/v1/web/aweme/post/?sec_user_id=%s&count=%d&max_cursor=%d&aid=1128&version_name=23.5.0&device_platform=android&os_version=2333", secUid, count, cursor)
		req, err := http.NewRequest(http.MethodGet, apiURL, nil)
		req.Header = SetupHeaders()

		if err != nil {
			log.Printf("fail to get post, secUid: %s, err: %s", secUid, err)
			return []string{}, err
		}

		resp, err := http.DefaultClient.Do(req)

		if err != nil {
			log.Printf("fail to get post, secUid: %s, err: %s", secUid, err)
			return []string{}, err
		}
		defer resp.Body.Close()

		d.printf("resp status: %s", resp.Status)

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Printf("fail to read body, secUid: %s, err: %s", secUid, err)
			return []string{}, err
		}

		var userPost model.UserPost
		err = json.Unmarshal(body, &userPost)

		if err != nil {
			log.Printf("fail to unmarshal json, secUid: %s, err: %s, body: %s", secUid, err, body)
			return []string{}, err
		}
		if userPost.StatusCode != 0 {
			return []string{}, fmt.Errorf("resp err, resp: %s", body)
		}

		init = false
		cursor = int(userPost.MaxCursor)
		for _, item := range userPost.AwemeList {
			idList = append(idList, item.AwemeID)
		}
	}

	return idList, nil
}
