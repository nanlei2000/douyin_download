package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/urfave/cli/v2"
)

const MAX_CONCURRENT_NUM = 4

func main() {
	var isDebug bool
	var path string
	var useID bool

	app := &cli.App{
		Name:  "dydl",
		Usage: "下载抖音视频",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "v",
				Aliases:     []string{"verbose"},
				Value:       false,
				Usage:       "切换 verbose 模式",
				Destination: &isDebug,
			},
			&cli.BoolFlag{
				Name:        "id",
				Value:       false,
				Usage:       "使用 id 下载",
				Destination: &useID,
			},
			&cli.StringFlag{
				Name:        "p",
				Aliases:     []string{"path"},
				Value:       "./video/",
				Usage:       "文件下载路径",
				Destination: &path,
			},
		},
		Action: func(c *cli.Context) error {
			if c.NArg() == 0 {
				return fmt.Errorf("url is required")
			}

			if useID {
				idStr := c.Args().Get(0)
				idList := strings.Split(idStr, ",")

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

						dy := NewDouYin()
						dy.isDebug = isDebug
						video, err := dy.Get(Source{
							Type:    SourceType_VideoID,
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
			shareContent := ""
			if c.NArg() > 0 {
				shareContent = strings.Join(c.Args().Slice(), "")
			} else {
				return fmt.Errorf("url is required")
			}

			dy := NewDouYin()
			dy.isDebug = isDebug
			video, err := dy.Get(Source{
				Type:    SourceType_ShardContent,
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
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
