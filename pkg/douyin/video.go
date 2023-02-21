package douyin

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// VideoType 视频类型
type VideoType int

const (
	//VideoPlayType 视频类
	VideoPlayType VideoType = 0
	//ImagePlayType 图文类
	ImagePlayType VideoType = 1
)

// Video 抖音视频
type Video struct {
	VideoId         string   `json:"video_id"`
	PlayId          string   `json:"play_id"`
	PlayAddr        string   `json:"play_addr"`
	VideoRawAddr    string   `json:"video_raw_addr"`
	PlayRawAddr     string   `json:"play_raw_addr"`
	OriginCoverList []string `json:"origin_cover_list"`
	Desc            string   `json:"desc"`
	RawLink         string   `json:"raw_link"`
	Author          struct {
		Id           string `json:"id"`
		ShortId      string `json:"short_id"`
		Nickname     string `json:"nickname"`
		AvatarLarger string `json:"avatar_larger"`
		Signature    string `json:"signature"`
	} `json:"author"`
	Images    []ImageItem `json:"images"`
	VideoType VideoType   `json:"video_type"`
}

type ImageItem struct {
	ImageUrl string `json:"image_url"`
	ImageId  string `json:"image_id"`
}

func (v *Video) GetFilename() string {
	if ext := filepath.Ext(v.PlayId); ext != "" {
		return v.VideoId + ext
	}
	return v.VideoId + ".mp4"
}

// Download 下载视频、图文到文件到指定目录，返回视频地址（图文为背景音乐视频地址）
func (v *Video) Download(distDir string) (path string, err error) {
	defer func() {
		if pErr := recover(); pErr != nil {
			log.Printf("出现panic: [filename=%s] [errmsg=%s]", distDir, err)
			err = fmt.Errorf("%s", pErr)
		}
	}()
	distDir, err = filepath.Abs(distDir)
	if err != nil {
		log.Printf("获取报错地址失败 [filename=%s] [error=%+v]", distDir, err)
		return "", err
	}
	folderName := fmt.Sprintf("%s_%s", v.Author.Nickname, v.Author.ShortId)
	distDir = filepath.Join(distDir, folderName, v.GetFilename())
	dir := filepath.Dir(distDir)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return "", err
		}
	}
	//如果是图片类，则将图片下载到指定目录
	if v.VideoType == ImagePlayType {
		return "", fmt.Errorf("暂时不支持图文下载")
	}

	if _, err := os.Stat(distDir); !os.IsNotExist(err) {
		log.Printf("视频本地已存在，[filename=%s]", distDir)
		return distDir, nil
	}

	req, err := http.NewRequest(http.MethodGet, v.PlayAddr, nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("User-Agent", DefaultUserAgent)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	f1, err := os.Create(distDir)
	if err != nil {
		log.Printf("创建文件失败 [filename=%s] [errmsg=%+v]", distDir, err)
		return "", err
	}
	defer f1.Close()
	_, err = io.Copy(f1, resp.Body)
	if err != nil {
		log.Printf("创建文件失败 [filename=%s] [errmsg=%+v]", distDir, err)
		return "", err
	}

	log.Printf("写入文件成功： [filename=%s]", distDir)

	return distDir, nil
}

func (v *Video) String() string {
	b, err := json.Marshal(v)
	if err != nil {
		log.Printf("编码失败 -> %s", err)
	} else {
		return string(b)
	}
	return fmt.Sprintf("%+v", *v)
}
