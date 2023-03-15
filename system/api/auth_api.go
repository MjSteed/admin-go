package api

import (
	"strconv"
	"time"

	"github.com/MjSteed/vue3-element-admin-go/common"
	"github.com/MjSteed/vue3-element-admin-go/common/model/vo"
	"github.com/MjSteed/vue3-element-admin-go/system/model/dto"
	"github.com/MjSteed/vue3-element-admin-go/system/service"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type authApi struct{}

var AuthApi = new(authApi)

// 登录
// @Router    /api/v1/auth/login [POST]
func (api *authApi) Login(c *gin.Context) {
	login := dto.Login{}
	c.ShouldBindQuery(&login)
	if user, err := service.UserService.Login(login.Username, login.Password); err != nil {
		vo.FailMsg(err.Error(), c)
	} else {
		config := common.Config.Jwt
		token := jwt.NewWithClaims(
			jwt.SigningMethodHS256,
			jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(config.Ttl) * time.Second)),
				ID:        strconv.FormatInt(user.Id, 10),
				Issuer:    config.Issuer,
				NotBefore: jwt.NewNumericDate(time.Now().Add(-1000 * time.Second)),
			},
		)
		tokenStr, err := token.SignedString([]byte(config.Secret))
		if err != nil {
			vo.FailMsg(err.Error(), c)
			return
		}
		res := make(map[string]string)
		res["accessToken"] = tokenStr
		vo.SuccessData(res, c)
	}
}
