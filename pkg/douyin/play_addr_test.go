package douyin

import "testing"

func TestGetPlayAddr(t *testing.T) {
	playStr, err := GetPlayAddr("7204783511956442372")
	if err != nil {
		t.Log(err)
	}
	t.Log(playStr)
}
