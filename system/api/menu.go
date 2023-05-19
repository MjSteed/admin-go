package api

import (
	"strconv"
	"strings"

	"github.com/MjSteed/vue3-element-admin-go/common/model/vo"
	"github.com/MjSteed/vue3-element-admin-go/system/model/dto"
	"github.com/MjSteed/vue3-element-admin-go/system/service"
	"github.com/gin-gonic/gin"
)

// MenuApi 菜单路由接口
type MenuApi struct {
	menuService *service.MenuService
}

// MenuApi 实例化路由接口
func NewMenuApi(menuService *service.MenuService) MenuApi {
	return MenuApi{menuService: menuService}
}

// 资源(菜单+权限)列表
// @Router    /api/v1/menus/resources [get]
func (a *MenuApi) ListResources(c *gin.Context) {
	vo.SuccessData(a.menuService.ListResources(c), c)
}

// 菜单列表
// @Router    /api/v1/menus [get]
func (a *MenuApi) List(c *gin.Context) {
	var pageParam dto.DeptPageReq
	err := c.ShouldBindQuery(&pageParam)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	list, err := a.menuService.ListPages(c, pageParam)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	vo.SuccessData(list, c)
}

// 菜单下拉列表
// @Router    /api/v1/menus/options [get]
func (a *MenuApi) ListOptions(c *gin.Context) {
	vo.SuccessData(a.menuService.Options(c), c)
}

// 路由列表
// @Router    /api/v1/menus/routes [get]
func (a *MenuApi) ListRoutes(c *gin.Context) {
	vo.SuccessData(a.menuService.ListRoutes(c), c)
}

// 路由列表
// @Router    /api/v1/menus/:id [get]
func (a *MenuApi) GetById(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	menu, err := a.menuService.GetById(c, id)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	vo.SuccessData(menu, c)
}

// 新增菜单
// @Router    /api/v1/menus [post]
func (a *MenuApi) Save(c *gin.Context) {
	var d dto.MenuForm
	err := c.ShouldBindJSON(&d)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	m := d.ToMenu()
	err = a.menuService.Save(c, &m)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	vo.Success(c)
}

// 修改菜单
// @Router    /api/v1/menus [put]
func (a *MenuApi) Update(c *gin.Context) {
	var d dto.MenuForm
	err := c.ShouldBindJSON(&d)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	m := d.ToMenu()
	err = a.menuService.Save(c, &m)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	vo.Success(c)
}

// 删除部门
// @Router    /api/v1/dept/{ids} [delete]
func (a *MenuApi) BatchDelete(c *gin.Context) {
	idsStr := strings.Split(c.Param("ids"), ",")
	ids := make([]int64, len(idsStr))
	for _, v := range idsStr {
		id, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			continue
		}
		ids = append(ids, id)
	}
	err := a.menuService.DeleteByIds(c, ids)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	vo.Success(c)
}
