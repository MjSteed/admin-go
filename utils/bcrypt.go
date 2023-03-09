package utils

import (
	"github.com/MjSteed/vue3-element-admin-go/common"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func BcryptMakeStr(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		common.LOG.Warn("加密失败", zap.Error(err))
	}
	return string(hash)
}

// 校验密码
func BcryptMakeCheckStr(pwd string, hashedPwd string) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, []byte(pwd))
	return err == nil
}
