package douyin

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"

	iteminfo "github.com/nanlei2000/douyin_download/pkg/model/item_info"
)

var (
	urlReg           = regexp.MustCompile(`http[s]?://(?:[a-zA-Z]|[0-9]|[$-_@.&+]|[!*\(\),]|(?:%[0-9a-fA-F][0-9a-fA-F]))+`)
	digitReg         = regexp.MustCompile(`\d+`)
	DefaultUserAgent = `Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1`
	relRrlStr        = `https://www.iesdouyin.com/web/api/v2/aweme/iteminfo/?item_ids=`
)

type SourceType uint

const (
	SourceType_ShardContent = iota
	SourceType_VideoID
)

type Source struct {
	Type    SourceType
	Content string
}

type DouYin struct {
	pattern *regexp.Regexp
	isDebug bool
	log     *log.Logger
}

func NewDouYin() *DouYin {
	return &DouYin{pattern: urlReg, isDebug: true, log: log.Default()}
}

func (d *DouYin) IsDebug(debug bool) {
	d.isDebug = debug
}

func (d *DouYin) GetRedirectUrl(urlStr string) (string, error) {
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
	d.printf("请求原始内容： %s", string(body))
	result := digitReg.FindString(string(body))
	if result == "" {
		return "", errors.New("解析参数失败 ->" + string(body))
	}
	return relRrlStr + result, nil
}

func (d *DouYin) GetVideoInfo(urlStr string) (string, error) {
	req, err := http.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("User-Agent", DefaultUserAgent)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	return string(body), nil
}

func (d *DouYin) Get(src Source) (v Video, err error) {
	defer func() {
		if pErr := recover(); pErr != nil {
			log.Printf("[DouYin.Get]panic, src: %v,err: %v", src, pErr)
			err = fmt.Errorf("%s", pErr)
		}
	}()
	var rawUrlStr string
	var shardContent string
	var urlStr string

	switch src.Type {
	case SourceType_ShardContent:
		shardContent = src.Content
		urlStr := d.pattern.FindString(shardContent)
		if urlStr == "" {
			return Video{}, errors.New("获取视频链接失败")
		}
		rawUrlStr, err = d.GetRedirectUrl(urlStr)
		if err != nil {
			return Video{}, err
		}
	case SourceType_VideoID:
		rawUrlStr = fmt.Sprintf("https://www.iesdouyin.com/web/api/v2/aweme/iteminfo/?item_ids=%s", src.Content)
	default:
		return Video{}, fmt.Errorf("unsupported src type")
	}

	body, err := d.GetVideoInfo(rawUrlStr)
	if err != nil {
		return Video{}, err
	}
	d.printf("获取抖音视频成功 -> [resp=%s]", body)

	var info iteminfo.ItemInfo

	err = json.Unmarshal([]byte(body), &info)

	if err != nil {
		return Video{}, err
	}
	if info.StatusCode != 0 {
		return Video{}, fmt.Errorf("resp err, resp: %s", body)
	}

	item := info.ItemList[0]
	video := Video{
		RawLink:      shardContent,
		VideoRawAddr: urlStr,
		PlayRawAddr:  rawUrlStr,
		Images:       []ImageItem{},
	}

	video.PlayAddr = strings.ReplaceAll(item.Video.PlayAddr.URLList[0], "playwm", "play")

	//获取播放时长，视频有播放时长，图文类无播放时长
	if item.Duration != 0 {
		video.VideoType = VideoPlayType
	} else {
		video.VideoType = ImagePlayType
		images := item.Images
		for _, image := range images {
			imageRes := image.URLList[0]
			video.Images = append(video.Images, ImageItem{
				ImageUrl: imageRes,
				ImageId:  image.URI,
			})
		}
	}
	//获取播放地址
	video.PlayId = item.Video.PlayAddr.URI
	//获取视频唯一id
	video.VideoId = item.AwemeID
	//获取封面
	video.Cover = item.Video.Cover.URLList[0]
	//获取原始封面
	video.OriginCover = item.Video.OriginCover.URLList[0]
	//获取作者信息
	video.Author.Id = item.Author.Uid
	video.Author.ShortId = item.Author.ShortID
	video.Author.Nickname = item.Author.Nickname
	video.Author.Signature = item.Author.Signature
	//获取视频描述
	video.Desc = item.Desc
	//获取作者大头像
	video.Author.AvatarLarger = item.Author.AvatarLarger.URLList[0]
	d.printf("解析后数据 [video=%s]", video.String())

	return video, nil
}

func (d *DouYin) printf(format string, v ...any) {
	if d.isDebug {
		d.log.Printf(format, v...)
	}
}
