package dto

import "github.com/MjSteed/vue3-element-admin-go/system/model"

type UserPageReq struct {
	Keywords string `form:"keywords"`
	Status   string `form:"status"`
	DeptId   int64  `form:"deptId"`
	PageNum  int    `form:"pageNum"`
	PageSize int    `form:"pageSize"`
}

// 用户表单对象
type UserForm struct {
	Id       int64   `form:"id"`
	Username string  `form:"username"`
	Nickname string  `form:"nickname"`
	Mobile   string  `form:"mobile"`
	Gender   int     `form:"gender"`
	Avatar   string  `form:"avatar"`
	Email    string  `form:"email"`
	Status   int     `form:"status"`
	DeptId   int64   `form:"deptId"`
	RoleIds  []int64 `form:"roleIds"`
}

func (m *UserForm) ToUser() model.SysUser {
	d := model.SysUser{
		Id:       m.Id,
		Username: m.Username,
		Nickname: m.Nickname,
		Gender:   m.Gender,
		DeptId:   m.DeptId,
		Avatar:   m.Avatar,
		Mobile:   m.Mobile,
		Status:   m.Status,
		Email:    m.Email,
	}
	return d
}
