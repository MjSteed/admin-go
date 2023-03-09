package utils

import "testing"

func Test_BcryptMakeCheckStr(t *testing.T) {
	pwd := "$2a$10$xVWsNOhHrCxh5UbpCE7/HuJ.PAOKcYAqRxD2CO2nVnJS.IAXkr5aq"
	b := BcryptMakeCheckStr("123456", pwd)
	if !b {
		t.Error("校验失败")
	}
	b = BcryptMakeCheckStr("1234567", pwd)
	if b {
		t.Error("校验失败")
	}
}
