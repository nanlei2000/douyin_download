package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

func main() {
	var isDebug bool
	var path string

	app := &cli.App{
		Name:  "dydl",
		Usage: "下载抖音视频",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "d",
				Aliases:     []string{"debug"},
				Value:       false,
				Usage:       "toggle debug mode",
				Destination: &isDebug,
			},
			&cli.StringFlag{
				Name:        "p",
				Aliases:     []string{"path"},
				Value:       "./video/",
				Usage:       "path to download",
				Destination: &path,
			},
		},
		Action: func(c *cli.Context) error {
			shareUrl := ""
			if c.NArg() > 0 {
				shareUrl = strings.Join(c.Args().Slice(), "")
			} else {
				return fmt.Errorf("url is required!/n")
			}

			dy := NewDouYin()
			dy.isDebug = isDebug
			video, err := dy.Get(shareUrl)
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
