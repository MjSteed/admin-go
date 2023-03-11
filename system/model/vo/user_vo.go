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

type UserAuthInfo struct {
	UserId    string   `json:"userId"`
	Username  string   `json:"username"`
	Nickname  string   `json:"nickname"`
	Password  string   `json:"password"`
	Status    int      `json:"status"`
	DeptId    int64    `json:"deptId"`
	Roles     []string `json:"roles"`
	Perms     []string `json:"perms"`
	DataScope int      `json:"dataScope"`
}
