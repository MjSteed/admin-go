package api

import (
	"strconv"
	"strings"

	"github.com/MjSteed/vue3-element-admin-go/common/model/vo"
	"github.com/MjSteed/vue3-element-admin-go/system/model"
	"github.com/MjSteed/vue3-element-admin-go/system/model/dto"
	"github.com/MjSteed/vue3-element-admin-go/system/service"
	"github.com/gin-gonic/gin"
)

type RoleApi struct{}

// 角色分页列表
// @Router    /api/v1/roles/pages [get]
func (a RoleApi) List(c *gin.Context) {
	var pageParam dto.DeptPageReq
	err := c.ShouldBindQuery(&pageParam)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	list, total, err := service.RoleService.ListPages(pageParam)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	vo.SuccessData(vo.PageResult{List: list, Total: total}, c)
}

// 角色下拉列表
// @Router    /api/v1/roles/options [get]
func (a RoleApi) ListOptions(c *gin.Context) {
	vo.SuccessData(service.RoleService.ListOptions(), c)
}

// 角色详情
// @Router    /api/v1/roles/:id [get]
func (a RoleApi) GetForm(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	r := service.RoleService.GetById(id)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	vo.SuccessData(r, c)
}

// 新增角色
// @Router    /api/v1/roles [post]
func (a RoleApi) Save(c *gin.Context) {
	var d model.SysRole
	err := c.ShouldBindJSON(d)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	err = service.RoleService.Save(d)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	vo.Success(c)
}

// 修改菜单
// @Router    /api/v1/roles [put]
func (a RoleApi) Update(c *gin.Context) {
	var d model.SysRole
	err := c.ShouldBindJSON(d)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	err = service.RoleService.Save(d)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	vo.Success(c)
}

// 删除角色
// @Router    /api/v1/roles/:ids [delete]
func (a RoleApi) BatchDelete(c *gin.Context) {
	idsStr := strings.Split(c.Param("ids"), ",")
	ids := make([]int64, len(idsStr))
	for _, v := range idsStr {
		id, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			continue
		}
		ids = append(ids, id)
	}
	err := service.RoleService.DeleteByIds(ids)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	vo.Success(c)
}

// 修改角色状态
// @Router    /api/v1/:id/status [PUT]
func (a RoleApi) UpdateRoleStatus(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	status, err := strconv.ParseInt(c.Param("status"), 10, 8)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	err = service.RoleService.UpdateStatus(id, int(status))
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	vo.Success(c)
}

// 获取角色的菜单ID集合
// @Router    /api/v1/:id/menuIds [GET]
func (a RoleApi) GetRoleMenuIds(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	menus := service.RoleService.GetRoleMenuIds(id)
	vo.SuccessData(menus, c)
}

// 分配角色的资源权限
// @Router    /api/v1/:id/menus [PUT]
func (a RoleApi) UpdateRoleMenus(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	var menus []int64
	err = c.BindJSON(&menus)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	err = service.RoleService.UpdateRoleMenus(id, menus)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	vo.Success(c)
}
