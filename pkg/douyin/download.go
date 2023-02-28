package douyin

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// Video 抖音视频
type Video struct {
	VideoId  string
	PlayAddr string
	Author   struct {
		SecUid   string
		Nickname string
	}
}

// https://www.douyin.com/aweme/v1/play/?video_id=v0d00fg10000cfukpcbc77u4ta5n9mk0&line=0&file_id=b1e06183b7cb4268a3b10e9cffa0d68d&sign=2404e9d800a44f5f8878cc758704c36d&is_play_url=1&source=PackSourceEnum_AWEME_DETAIL&aid=6383

// Download 下载视频
func DownloadVideo(v Video, distDir string) (path string, err error) {
	defer func() {
		if pErr := recover(); pErr != nil {
			log.Printf("出现panic, filename: %s, err: %s", distDir, err)
			err = fmt.Errorf("%s", pErr)
		}
	}()
	distDir, err = filepath.Abs(distDir)
	if err != nil {
		log.Printf("获取绝对地址失败, filename: %s, err: %s", distDir, err)
		return "", err
	}
	folderName := fmt.Sprintf("%s_%s", v.Author.Nickname, v.Author.SecUid)
	distDir = filepath.Join(distDir, folderName, v.VideoId+".mp4")
	dir := filepath.Dir(distDir)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return "", err
		}
	}

	if _, err := os.Stat(distDir); !os.IsNotExist(err) {
		log.Printf("视频本地已存在， filename: %s", distDir)
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
		log.Printf("创建文件失败, filename: %s, err: %s", distDir, err)
		return "", err
	}
	defer f1.Close()
	_, err = io.Copy(f1, resp.Body)
	if err != nil {
		log.Printf("创建文件失败, filename: %s, err: %s", distDir, err)
		return "", err
	}

	log.Printf("写入文件成功, filename:%s", distDir)

	return distDir, nil
}
