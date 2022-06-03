package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

const MAX_CONCURRENT_NUM = 5

func main() {
	var verbose bool
	var path string
	var downloadUserPost bool
	var weiboMode bool

	app := &cli.App{
		Name:  "dydl",
		Usage: "下载抖音视频",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "up",
				Aliases:     []string{"user_post"},
				Value:       false,
				Usage:       "下载所有发布的内容",
				Destination: &downloadUserPost,
			},
			&cli.BoolFlag{
				Name:        "wb",
				Aliases:     []string{"weibo"},
				Value:       false,
				Usage:       "下载微博图片",
				Destination: &weiboMode,
			},
			&cli.StringFlag{
				Name:        "p",
				Aliases:     []string{"path"},
				Value:       "./video/",
				Usage:       "文件下载路径",
				Destination: &path,
			},
			&cli.BoolFlag{
				Name:        "v",
				Aliases:     []string{"verbose"},
				Value:       false,
				Usage:       "切换 verbose 模式",
				Destination: &verbose,
			},
		},
		Action: func(c *cli.Context) error {
			if c.NArg() == 0 {
				return fmt.Errorf("url is required")
			}

			if weiboMode {
				return HandleWeiboCmd(c.Args().Get(0), downloadUserPost, path)
			}

			return HandleDouyinCmd(c, verbose, downloadUserPost, path)
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
