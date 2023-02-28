package douyin

import "testing"

func TestGetPlayAddr(t *testing.T) {
	playStr, err := GetVideoDetail("7015482603633511683")
	if err != nil {
		t.Log(err)
	}
	t.Log(playStr)
}
