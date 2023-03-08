package service

import (
	"errors"

	"github.com/MjSteed/vue3-element-admin-go/common"
	"github.com/MjSteed/vue3-element-admin-go/system/model"
	"github.com/MjSteed/vue3-element-admin-go/system/model/dto"
	"github.com/MjSteed/vue3-element-admin-go/system/model/vo"
)

type userService struct{}

var UserService = new(userService)

// 用户分页列表
func (service *userService) ListPages(pageReq dto.UserPageReq) (list []vo.SysUser, total int64, err error) {
	tx := common.DB.Table("sys_user")
	tx = tx.Joins("LEFT JOIN sys_dept ON sys_user.dept_id = sys_dept.id")
	tx = tx.Joins("LEFT JOIN sys_user_role ON sys_user.id = sys_user_role.user_id")
	tx = tx.Joins("LEFT JOIN sys_role ON sys_user_role.role_id = sys_role.id")
	if pageReq.Keywords != "" {
		keywords := "%" + pageReq.Keywords + "%"
		tx = tx.Where("(sys_user.username like ? or sys_user.nickname LIKE ? or sys_user.mobile LIKE ?)", keywords, keywords, keywords)
	}
	if pageReq.Status != "" {
		tx = tx.Where("sys_user.status = ?", pageReq.Status)
	}
	if pageReq.DeptId != 0 {
		tx = tx.Where("concat(',',concat(sys_dept.tree_path,',',sys_dept.id),',') like concat('%,',?,',%')", pageReq.DeptId)
	}
	tx = tx.Select("count(distinct sys_user.id)").Count(&total)
	tx = tx.Select("sys_user.id,sys_user.username,sys_user.nickname,sys_user.mobile,sys_user.gender,sys_user.avatar,sys_user.STATUS,sys_dept.NAME as dept_name,GROUP_CONCAT(sys_role.NAME) as role_names,sys_user.create_time")
	err = tx.Debug().Group("sys_user.id").Limit(pageReq.PageSize).Offset(pageReq.PageSize * (pageReq.PageNum - 1)).Find(&list).Error
	return
}

// 获取用户表单数据
func (service *userService) GetById(id int64) (data model.SysUser, err error) {
	err = common.DB.First(&data, id).Error
	return
}

// 新增用户
func (service *userService) Save(data model.SysUser) (err error) {
	var c int64
	err = common.DB.Model(&data).Where("username = ?", data.Username).Count(&c).Error
	if err != nil {
		return
	}
	if c > 0 {
		err = errors.New("用户名已存在")
		return
	}
	err = common.DB.Create(&data).Error
	//TODO 保存用户角色
	return
}

// 修改用户
func (service *userService) Update(data model.SysUser) (err error) {
	var c int64
	err = common.DB.Model(&data).Where("username = ?", data.Username).Where("id != ?", data.Id).Count(&c).Error
	if err != nil {
		return
	}
	if c > 0 {
		err = errors.New("用户名已存在")
		return
	}
	err = common.DB.Updates(&data).Error
	//TODO 保存用户角色
	return
}

// 删除用户
func (service *userService) DeleteByIds(ids []int64) (err error) {
	err = common.DB.Where("id in ?", ids).Delete(&model.SysRole{}).Error
	return
}

// 修改用户密码
func (service *userService) UpdatePassword(id int64, password string) (err error) {
	err = common.DB.Model(&model.SysUser{Id: id}).Update("password", password).Error
	return
}

// 根据用户名获取认证信息
func (service *userService) GetAuthInfo(username string) (data model.SysUser, err error) {
	return
}
