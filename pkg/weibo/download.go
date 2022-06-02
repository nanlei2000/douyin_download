package weibo

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/nanlei2000/douyin_download/internal/utils"
)

type DownLoadType int16

const (
	Show = iota
	ImageWall
)

type Source struct {
	Type DownLoadType
	Link string
}

func (w *Weibo) DownLoadShowPics(src Source, distDir string) (err error) {
	var imageSet ImageSet

	switch src.Type {
	case Show:
		id, err := utils.GetLastURLPath(src.Link)
		if err != nil {
			return err
		}
		imageSet, err = w.GetShowPics(id)
		if err != nil {
			return err
		}
	case ImageWall:
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

	for _, p := range pics {
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
			continue
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Printf("解析图像出错 -> [src=%v] [image_url=%s] [err=%s]", src, p, err)
			continue
		}
		_ = resp.Body.Close()
		err = ioutil.WriteFile(filePath, b, os.ModePerm)
		if err != nil {
			log.Printf("解析图像出错 -> [src=%v] [image_url=%s] [err=%s]", src, p, err)
			continue
		}
		log.Printf("写入图片成功, filePath: %s", filePath)
		time.Sleep(time.Microsecond * 100)
	}

	return nil
}
