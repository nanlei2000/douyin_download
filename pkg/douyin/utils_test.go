package douyin

import "testing"

func TestGenRandomStr(t *testing.T) {
	str := genRandomStr(10)
	t.Log(str)
}
