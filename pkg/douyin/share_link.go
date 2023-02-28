package douyin

import (
	"errors"
	"io"
	"net/http"
	"regexp"
)

var (
	urlReg           = regexp.MustCompile(`http[s]?://(?:[a-zA-Z]|[0-9]|[$-_@.&+]|[!*\(\),]|(?:%[0-9a-fA-F][0-9a-fA-F]))+`)
	DefaultUserAgent = `Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1`
	digitReg         = regexp.MustCompile(`\d+`)
)

func getVideoID(urlStr string) (string, error) {
	header := http.Header{}
	header.Add("User-Agent", DefaultUserAgent)
	header.Add("Upgrade-Insecure-Requests", "1")

	req, err := http.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return "", err
	}
	req.Header = header
	resp, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	result := digitReg.FindString(string(body))
	if result == "" {
		return "", errors.New("解析参数失败 ->" + string(body))
	}

	return result, nil
}

func GetVideoIDBySharedLink(sharedLink string) (string, error) {
	urlStr := urlReg.FindString(sharedLink)
	if urlStr == "" {
		return "", errors.New("获取视频链接失败")
	}

	return getVideoID(urlStr)
}
