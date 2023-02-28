package douyin

import (
	"testing"
)

func TestVideo_Download_FromID(t *testing.T) {
	v := Video{
		VideoId:  "7015873671797411111",
		PlayAddr: "https://v26-web.douyinvod.com/3b4717f38b67e080bd0e58886d7c51f3/63fdbca4/video/tos/cn/tos-cn-ve-15c001-alinc2/owStiIJdnofpi8QhADegGE7PkCbHHEBQDJAZBf/?a=6383&ch=26&cr=3&dr=0&lr=all&cd=0%7C0%7C0%7C3&cv=1&br=1852&bt=1852&cs=0&ds=4&ft=7yV4ZBo7UUmfTbdPD0PD1YjCV3R4tG-NyBS9eFc2lKpr12nz&mime_type=video_mp4&qs=0&rc=ZTk1ZWczPGc4aWk4ZDtnOUBpamx0ZTs6Zm07aTMzNGkzM0AuLmFjNWJjNWIxNjU2LzFfYSM1MDM0cjRvZHJgLS1kLTBzcw%3D%3D&l=20230228153435EED9780FFF6EDB03498B&btag=8000",
		Author: struct {
			SecUid   string
			Nickname string
		}{
			SecUid:   "s",
			Nickname: "test",
		},
	}
	DownloadVideo(v, "./video")
}
