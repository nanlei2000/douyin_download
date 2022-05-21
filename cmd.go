package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	var isDebug bool

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
		},
		Action: func(c *cli.Context) error {
			shareUrl := ""
			if c.NArg() > 0 {
				shareUrl = c.Args().Get(0)
			} else {
				return fmt.Errorf("url is required!/n")
			}

			dy := NewDouYin()
			dy.isDebug = isDebug
			video, err := dy.Get(shareUrl)
			if err != nil {
				return err
			}
			p, err := video.Download("./video/")

			fmt.Printf("dist file: %s/n", p)

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
