package main

import (
	"fmt"
	"math/rand"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/nanlei2000/douyin_download/pkg/douyin"
	"github.com/urfave/cli/v2"
)

func HandleDouyinCmd(c *cli.Context, verbose bool, downloadUserPost bool, path string) error {
	dy := douyin.NewDouYin()
	dy.IsDebug(verbose)

	// https://www.douyin.com/user/MS4wLjABAAAAZimxk0o3KWTEJNNrzwSF3HBjCy4TkS6mpPyHNxEYC2A?relation=1
	if downloadUserPost {
		userLink := c.Args().Get(0)
		url, err := url.Parse(userLink)
		if err != nil {
			return err
		}
		parts := strings.Split(url.Path, "/")
		secUid := parts[len(parts)-1]

		if len(secUid) == 0 {
			return fmt.Errorf("url is invalid")
		}

		idList, err := dy.GetAllVideoIDList(secUid)

		if err != nil {
			return err
		}

		c := make(chan struct{}, MAX_CONCURRENT_NUM)
		defer close(c)

		var wg sync.WaitGroup
		for _, id := range idList {
			wg.Add(1)
			go func(id string) {
				c <- struct{}{}
				defer func() {
					wg.Done()
					<-c
				}()

				ran := rand.Int31n(100)
				time.Sleep(time.Duration(ran) * time.Millisecond)

				video, err := dy.Get(douyin.Source{
					Type:    douyin.SourceType_VideoID,
					Content: id,
				})
				if err != nil {
					fmt.Printf("get video info failed, id: %s, err: %s", id, err)
				}
				_, err = video.Download(path)
				if err != nil {
					fmt.Printf("download video failed, id: %s, err: %s", id, err)
				}
			}(id)
		}
		wg.Wait()

		return nil
	}

	// use shardContent
	shareContent := strings.Join(c.Args().Slice(), "")

	video, err := dy.Get(douyin.Source{
		Type:    douyin.SourceType_ShardContent,
		Content: shareContent,
	})
	if err != nil {
		return err
	}
	_, err = video.Download(path)
	if err != nil {
		return err
	}

	return nil
}
