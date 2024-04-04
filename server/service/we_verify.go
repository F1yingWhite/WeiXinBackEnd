package service

import (
	"crypto/sha1"
	"fmt"
	"net/http"
	"sort"
	"strings"

	"github.com/gin-gonic/gin"
)

func WX_authorization(c *gin.Context) {
	data := c.Request.URL.Query()
	if len(data) == 0 {
		c.String(http.StatusOK, "hello, this is handle view")
		return
	}

	signature := data.Get("signature")
	timestamp := data.Get("timestamp")
	nonce := data.Get("nonce")
	echostr := data.Get("echostr")
	token := "xyh"

	list := []string{token, timestamp, nonce}
	sort.Strings(list)
	hash := sha1.New()
	hash.Write([]byte(strings.Join(list, "")))
	hashcode := fmt.Sprintf("%x", hash.Sum(nil))
	fmt.Printf("handle/GET func: hashcode, signature: %s, %s\n", hashcode, signature)

	if hashcode == signature {
		c.String(http.StatusOK, echostr)
	} else {
		c.String(http.StatusOK, "")
	}
}
