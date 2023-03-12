package api

import (
	"strconv"
	"strings"

	"github.com/MjSteed/vue3-element-admin-go/common/model/vo"
	"github.com/MjSteed/vue3-element-admin-go/system/model/dto"
	"github.com/MjSteed/vue3-element-admin-go/system/service"
	"github.com/gin-gonic/gin"
)

type menuApi struct{}

var MenuApi = new(menuApi)

// 资源(菜单+权限)列表
// @Router    /api/v1/menus/resources [get]
func (a *menuApi) ListResources(c *gin.Context) {
	vo.SuccessData(service.MenuService.ListResources(), c)
}

// 菜单列表
// @Router    /api/v1/menus [get]
func (a *menuApi) List(c *gin.Context) {
	var pageParam dto.DeptPageReq
	err := c.ShouldBindQuery(&pageParam)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	list, err := service.MenuService.ListPages(pageParam)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	vo.SuccessData(list, c)
}

// 菜单下拉列表
// @Router    /api/v1/menus/options [get]
func (a *menuApi) ListOptions(c *gin.Context) {
	vo.SuccessData(service.MenuService.ListOptions(), c)
}

// 路由列表
// @Router    /api/v1/menus/routes [get]
func (a *menuApi) ListRoutes(c *gin.Context) {
	vo.SuccessData(service.MenuService.ListRoutes(), c)
}

// 路由列表
// @Router    /api/v1/menus/:id [get]
func (a *menuApi) GetById(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	menu, err := service.MenuService.GetById(id)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	vo.SuccessData(menu, c)
}

// 新增菜单
// @Router    /api/v1/menus [post]
func (a *menuApi) Save(c *gin.Context) {
	var d dto.MenuForm
	err := c.ShouldBindJSON(&d)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	m := d.ToMenu()
	err = service.MenuService.Save(&m)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	vo.Success(c)
}

// 修改菜单
// @Router    /api/v1/menus [put]
func (a *menuApi) Update(c *gin.Context) {
	var d dto.MenuForm
	err := c.ShouldBindJSON(&d)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	m := d.ToMenu()
	err = service.MenuService.Save(&m)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	vo.Success(c)
}

// 删除部门
// @Router    /api/v1/dept/{ids} [delete]
func (a *menuApi) BatchDelete(c *gin.Context) {
	idsStr := strings.Split(c.Param("ids"), ",")
	ids := make([]int64, len(idsStr))
	for _, v := range idsStr {
		id, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			continue
		}
		ids = append(ids, id)
	}
	err := service.MenuService.DeleteByIds(ids)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	vo.Success(c)
}
