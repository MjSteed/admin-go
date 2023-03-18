package middleware

import (
	"errors"
	"fmt"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/MjSteed/vue3-element-admin-go/common"
	"github.com/MjSteed/vue3-element-admin-go/common/model/vo"
	"github.com/MjSteed/vue3-element-admin-go/utils"
	"github.com/gin-gonic/gin"
)

// MD5参数签名校验
func SignMd5() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := verifySign(c)
		if err != nil {
			vo.FailMsg(err.Error(), c)
			c.Abort()
			return
		}
		c.Next()
	}
}

// 验证签名
func verifySign(c *gin.Context) error {
	_ = c.Request.ParseForm()
	req := c.Request.Form
	ak := strings.Join(c.Request.Form["ak"], "")
	sn := strings.Join(c.Request.Form["sn"], "")
	ts := strings.Join(c.Request.Form["ts"], "")

	// 验证来源
	appSecret, ok := common.Config.Sign.Auth[ak]
	if !ok {
		return errors.New("ak Error")
	}

	// 验证过期时间
	timestamp := time.Now().Unix()
	exp := common.Config.Sign.Exp
	tsInt, _ := strconv.ParseInt(ts, 10, 64)
	if tsInt > timestamp || timestamp-tsInt >= exp {
		return errors.New("ts Error")
	}

	// 验证签名
	if sn == "" || sn != createSign(req, appSecret) {
		return errors.New("sn Error")
	}

	return nil
}

// 创建签名
func createSign(params url.Values, appSecret string) string {
	// 自定义MD5组合
	return utils.MD5(createEncryptStr(params) + appSecret)
}

func createEncryptStr(params url.Values) string {
	var key []string
	var str = ""
	for k := range params {
		if k != "sn" {
			key = append(key, k)
		}
	}
	sort.Strings(key)
	for i := 0; i < len(key); i++ {
		if i == 0 {
			str = fmt.Sprintf("%v=%v", key[i], params.Get(key[i]))
		} else {
			str = str + fmt.Sprintf("&%v=%v", key[i], params.Get(key[i]))
		}
	}
	return str
}
