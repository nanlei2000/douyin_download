package weibo

import (
	"testing"
)

func TestGetAllImageWallObjId(t *testing.T) {
	t.Skip()
	w := Weibo{}
	images, err := w.GetAllImageWallPid("2286073303")

	if err != nil {
		t.Fatal(err)
	}
	t.Log(images)
	t.Log(len(images.IdList))
}

func TestFloatFormat(t *testing.T) {
	f := float64(0.01)

	t.Logf("%f", f)
	t.Logf("%.f", f)
	t.Logf("%v", f)
}
