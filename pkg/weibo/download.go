package weibo

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/nanlei2000/douyin_download/internal/utils"
)

func (w *Weibo) DownLoadShowPics(showLink string, distDir string) error {
	url, err := url.Parse(showLink)
	if err != nil {
		return err
	}
	parts := strings.Split(url.Path, "/")
	id := parts[len(parts)-1]

	pics, err := w.GetShowPics(id)
	if err != nil {
		return err
	}

	distDir, err = filepath.Abs(distDir)
	if err != nil {
		return err
	}

	if _, err := os.Stat(distDir); os.IsNotExist(err) {
		if err := os.MkdirAll(distDir, os.ModePerm); err != nil {
			return err
		}
	}

	for _, p := range pics {
		req, err := http.NewRequest("GET", p, nil)
		if err != nil {
			return err
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

		lastPath, err := utils.GetLastURLPath(p)
		if err != nil {
			return err
		}
		filePath := filepath.Join(distDir, lastPath)

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Printf("解析图像出错 -> [id=%s] [image_url=%s] [err=%s]", id, p, err)
			continue
		}
		_ = resp.Body.Close()
		err = ioutil.WriteFile(filePath, b, os.ModePerm)
		if err != nil {
			log.Printf("解析图像出错 -> [id=%s] [image_url=%s] [err=%s]", id, p, err)
			continue
		}
		time.Sleep(time.Microsecond * 110)
	}

	return nil
}
