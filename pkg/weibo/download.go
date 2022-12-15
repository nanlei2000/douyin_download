package weibo

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/nanlei2000/douyin_download/internal/utils"
)

const (
	MAX_CONCURRENT_NUM = 3
	RETRY_COUNT        = 5
)

type DownLoadType int16

const (
	DownLoadType_Show DownLoadType = iota
	DownLoadType_ImageWall
)

type Source struct {
	Type DownLoadType
	Link string
}

func (w *Weibo) DownLoad(src Source, distDir string) (err error) {
	var imageSet ImageSet

	switch src.Type {
	case DownLoadType_Show:
		id, err := utils.GetLastURLPath(src.Link)
		if err != nil {
			return err
		}
		imageSet, err = w.GetShowPics(id)
		if err != nil {
			return err
		}
	case DownLoadType_ImageWall:
		uid, err := utils.GetLastURLPath(src.Link)
		if err != nil {
			return err
		}
		imageSet, err = w.GetAllImageWallPid(uid)
		if err != nil {
			return err
		}

	default:
		return fmt.Errorf("unsupported src type")
	}

	var pics []string
	for _, id := range imageSet.IdList {
		// e.g https://wx4.sinaimg.cn/large/008s6isIgy1h93m0iw21kj30n618bwn4.jpg
		pics = append(pics, fmt.Sprintf("https://wx4.sinaimg.cn/large/%s.jpg", id))
	}

	distDir, err = filepath.Abs(distDir)
	distDir = filepath.Join(distDir, imageSet.Name)
	if err != nil {
		return err
	}

	if _, err := os.Stat(distDir); os.IsNotExist(err) {
		if err := os.MkdirAll(distDir, os.ModePerm); err != nil {
			return err
		}
	}

	return w.downloadPics(pics, distDir)
}

func (w *Weibo) downloadPics(pics []string, distDir string) error {
	var wg sync.WaitGroup
	c := make(chan struct{}, MAX_CONCURRENT_NUM)
	defer close(c)

	var lastErr error
	for _, p := range pics {
		wg.Add(1)
		p := p
		go func() (err error) {
			run := func() (err error) {
				c <- struct{}{}
				defer func() {
					if pErr := recover(); pErr != nil {
						err = fmt.Errorf("panic: err: %v", pErr)
					}
					if err != nil {
						log.Printf("解析图像出错 -> [image_url=%s] [err=%s]", p, err)
						lastErr = err
					}
					<-c
					wg.Done()
				}()

				// 防止频控
				ran := rand.Int31n(100)
				time.Sleep(time.Duration(ran) * time.Millisecond)

				req, err := http.NewRequest("GET", p, nil)
				if err != nil {
					return err
				}
				err = w.setupHeaders(req, false)

				req.Header.Set("Authority", "wx4.sinaimg.cn")
				req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
				req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
				req.Header.Set("Referer", "https://weibo.com/")
				req.Header.Set("Sec-Ch-Ua", "\"Not?A_Brand\";v=\"8\", \"Chromium\";v=\"108\", \"Microsoft Edge\";v=\"108\"")
				req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
				req.Header.Set("Sec-Ch-Ua-Platform", "\"Windows\"")
				req.Header.Set("Sec-Fetch-Dest", "document")
				req.Header.Set("Sec-Fetch-Mode", "navigate")
				req.Header.Set("Sec-Fetch-Site", "cross-site")
				req.Header.Set("Sec-Fetch-User", "?1")
				req.Header.Set("Upgrade-Insecure-Requests", "1")
				req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36 Edg/108.0.1462.46")

				if err != nil {
					return err
				}

				lastPath, err := utils.GetLastURLPath(p)
				if err != nil {
					return err
				}
				filePath := filepath.Join(distDir, lastPath)

				if _, err := os.Stat(filePath); !os.IsNotExist(err) {
					log.Printf("文件本地已存在, filePath: %s", filePath)
					return nil
				}

				resp, err := http.DefaultClient.Do(req)
				if err != nil {
					return err
				}

				if resp.StatusCode >= 400 {
					return fmt.Errorf("http status code err: %d", resp.StatusCode)
				}

				defer resp.Body.Close()

				f, err := os.Create(filePath)
				if err != nil {
					return err
				}

				_, err = io.Copy(f, resp.Body)
				if err != nil {
					return err
				}

				if err != nil {
					return err
				}
				log.Printf("写入图片成功, filePath: %s", filePath)

				return nil
			}

			for i := 0; i < RETRY_COUNT; i++ {
				err = run()
				if err != nil {
					continue
				}
				return nil
			}

			return err
		}()
	}
	wg.Wait()
	return lastErr
}
