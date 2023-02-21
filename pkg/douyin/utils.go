package douyin

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/corpix/uarand"
)

func genRandomStr(length int) string {
	randomStr := ""
	baseStr := "ABCDEFGHIGKLMNOPQRSTUVWXYZabcdefghigklmnopqrstuvwxyz0123456789="
	baseStrLen := len(baseStr)

	for i := 0; i < length; i++ {
		randomStr += string(baseStr[rand.Int31n(int32(baseStrLen))])
	}
	return randomStr
}

func SetupHeaders() http.Header {
	h := http.Header{}
	h.Add("User-Agent", uarand.GetRandom())
	h.Add("Cookie", fmt.Sprintf("msToken=%s", genRandomStr(107)))

	return h
}
