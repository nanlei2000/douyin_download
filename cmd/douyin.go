package main

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"sync"
	"time"

	"github.com/nanlei2000/douyin_download/pkg/douyin"
	"github.com/urfave/cli/v2"
)

const (
	MAX_CONCURRENT_NUM = 2
	RETRY_COUNT        = 5
)

func HandleDouyinCmd(c *cli.Context, verbose bool, downloadUserPost bool, path string) error {
	dy := douyin.NewDouYin()
	dy.IsDebug(verbose)

	// https://www.douyin.com/user/MS4wLjABAAAAZimxk0o3KWTEJNNrzwSF3HBjCy4TkS6mpPyHNxEYC2A?relation=1
	if downloadUserPost {
		userLink := c.Args().Get(0)
		idList, err := dy.GetAllVideoIDList(userLink)

		if err != nil {
			return err
		}

		c := make(chan struct{}, MAX_CONCURRENT_NUM)
		defer close(c)

		var wg sync.WaitGroup
		for _, id := range idList {
			id := id
			wg.Add(1)
			go func() {
				c <- struct{}{}
				defer func() {
					wg.Done()
					<-c
				}()

				run := func() error {
					ran := rand.Int31n(100) + 200
					time.Sleep(time.Duration(ran) * time.Millisecond)

					video, err := dy.Get(douyin.Source{
						Type:    douyin.SourceType_VideoID,
						Content: id,
					})
					if err != nil {
						return fmt.Errorf("get video info failed, id: %s, err: %s", id, err)
					}
					_, err = video.Download(path)
					if err != nil {
						return fmt.Errorf("download video failed, id: %s, err: %s", id, err)
					}

					return nil
				}

				var lastErr error
				for i := 0; i < RETRY_COUNT; i++ {
					if err := run(); err != nil {
						lastErr = err
						continue
					}
					break
				}

				if lastErr != nil {
					log.Printf("run fail, err: %s", lastErr)
				}
			}()
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
