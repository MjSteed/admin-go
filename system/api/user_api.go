package api

import (
	"strconv"
	"strings"

	"github.com/MjSteed/vue3-element-admin-go/common"
	"github.com/MjSteed/vue3-element-admin-go/common/model/vo"
	"github.com/MjSteed/vue3-element-admin-go/system/model/dto"
	"github.com/MjSteed/vue3-element-admin-go/system/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type userApi struct{}

var UserApi = new(userApi)

// 用户分页列表
// @Router    /api/v1/users/pages [get]
func (a *userApi) List(c *gin.Context) {
	var pageParam dto.UserPageReq
	err := c.ShouldBindQuery(&pageParam)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	list, total, err := service.UserService.ListPages(pageParam)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	vo.SuccessData(vo.PageResult{List: list, Total: total}, c)
}

// 用户表单数据
// @Router    /api/v1/users/:id/form [get]
func (a *userApi) GetForm(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	r, err := service.UserService.GetById(id)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	vo.SuccessData(r, c)
}

// 新增用户
// @Router    /api/v1/users [POST]
func (a *userApi) Save(c *gin.Context) {
	var d dto.UserForm
	err := c.ShouldBindJSON(d)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	err = service.UserService.Save(d)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	vo.Success(c)
}

// 修改用户
// @Router    /api/v1/users/:id [PUT]
func (a *userApi) Update(c *gin.Context) {
	var d dto.UserForm
	err := c.ShouldBindJSON(d)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	err = service.UserService.Update(d)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	vo.Success(c)
}

// 删除用户
// @Router    /api/v1/users/:ids [delete]
func (a *userApi) BatchDelete(c *gin.Context) {
	idsStr := strings.Split(c.Param("ids"), ",")
	ids := make([]int64, len(idsStr))
	for _, v := range idsStr {
		id, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			continue
		}
		ids = append(ids, id)
	}
	err := service.UserService.DeleteByIds(ids)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	vo.Success(c)
}

// 删除用户
// @Router    /api/v1/users/:id/password [delete]
func (a *userApi) UpdatePassword(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	password := c.Param("password")
	err = service.UserService.UpdatePassword(id, password)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	vo.Success(c)
}

// 获取登录用户信息
// @Router    /api/v1/users/me [GET]
func (a *userApi) GetUserLoginInfo(c *gin.Context) {
	id, err := strconv.ParseInt(c.Keys["id"].(string), 10, 64)
	common.LOG.Debug("从jwt中获取到用户", zap.Int64("id", id))
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	user, err := service.UserService.GetUserInfo(id)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	vo.SuccessData(user, c)
}
