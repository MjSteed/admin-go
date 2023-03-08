package vo

import "time"

type SysUser struct {
	Id          int64     `json:"id"`
	Username    string    `json:"username"`
	Nickname    string    `json:"nickname"`
	Mobile      string    `json:"mobile"`
	GenderLabel string    `json:"genderLabel"`
	Avatar      string    `json:"avatar"`
	Email       string    `json:"email"`
	Status      int       `json:"status"`
	DeptName    string    `json:"deptName"`
	RoleNames   string    `json:"roleNames"`
	CreateTime  time.Time `json:"createTime"`
}
