package main

import "github.com/nanlei2000/douyin_download/pkg/weibo"

func HandleWeiboCmd(link string, downloadUserPost bool, path string) error {
	src := weibo.Source{
		Type: weibo.DownLoadType_Show,
		Link: link,
	}
	if downloadUserPost {
		src.Type = weibo.DownLoadType_ImageWall
	}

	w := weibo.Weibo{}
	err := w.DownLoad(src, path)

	if err != nil {
		return err
	}

	return nil
}
