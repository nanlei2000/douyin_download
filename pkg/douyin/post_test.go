package douyin

import (
	"testing"
)

func TestGetAllVideoIDList(t *testing.T) {
	list, err := GetAllVideoIDList("MS4wLjABAAAAF0HqDm-8U9TiT_9AfqqPGiNbrP0c93AdB3_oRG7Em_Q")

	if err != nil {
		t.Fatalf("err: %s", err)
	}

	if len(list) != 73 {
		t.Fatalf("get all id failed, expect: 73, actual: %d", len(list))
	}
	t.Logf("list: %s", list)
}
