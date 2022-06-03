package weibo

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/nanlei2000/douyin_download/internal/utils"
)

const MAX_CONCURRENT_NUM = 5

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
		pics = append(pics, fmt.Sprintf("https://wx1.sinaimg.cn/large/%s.jpg", id))
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
		go func(p string) (err error) {
			c <- struct{}{}
			defer func() {
				wg.Done()
				if err != nil {
					lastErr = err
				}
				<-c
			}()

			// 防止频控
			ran := rand.Int31n(100)
			time.Sleep(time.Duration(ran) * time.Millisecond)

			req, err := http.NewRequest("GET", p, nil)
			if err != nil {
				return err
			}
			err = w.setupHeaders(req, false)
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
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Printf("解析图像出错 -> [image_url=%s] [err=%s]", p, err)
				return err
			}
			_ = resp.Body.Close()
			err = ioutil.WriteFile(filePath, b, os.ModePerm)
			if err != nil {
				log.Printf("解析图像出错 -> [image_url=%s] [err=%s]", p, err)
				return err
			}
			log.Printf("写入图片成功, filePath: %s", filePath)

			return nil
		}(p)
	}
	wg.Wait()
	return lastErr
}
