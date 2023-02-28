package douyin

import "testing"

func TestGetPlayAddr(t *testing.T) {
	playStr, err := GetVideoDetail("7183314664921615649")
	if err != nil {
		t.Log(err)
	}
	t.Log(playStr)
}
