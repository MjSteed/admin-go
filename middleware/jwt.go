package middleware

import (
	"github.com/MjSteed/vue3-element-admin-go/common"
	"github.com/MjSteed/vue3-element-admin-go/common/model/vo"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.Request.Header.Get("Authorization")
		if tokenStr == "" {
			vo.FailDetail(gin.H{"reload": true}, "未登录或非法访问", c)
			c.Abort()
			return
		}
		token, err := jwt.ParseWithClaims(tokenStr, &jwt.RegisteredClaims{}, func(t *jwt.Token) (interface{}, error) {
			return []byte(common.Config.Jwt.Secret), nil
		})
		common.LOG.Debug("token", zap.Any("Token", token))
		if err != nil {
			vo.FailDetail(gin.H{"reload": true}, err.Error(), c)
			c.Abort()
			return
		}
		claims := token.Claims.(*jwt.RegisteredClaims)
		common.LOG.Debug("claims", zap.Any("claims", claims))
		if claims.Issuer != common.Config.Jwt.Issuer {
			vo.FailDetail(gin.H{"reload": true}, "未登录或非法访问", c)
			c.Abort()
			return
		}
		c.Set("token", token)
		c.Set("id", claims.ID)
		c.Next()
	}
}
