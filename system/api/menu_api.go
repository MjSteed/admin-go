package api

import (
	"strconv"
	"strings"

	"github.com/MjSteed/vue3-element-admin-go/common/model/vo"
	"github.com/MjSteed/vue3-element-admin-go/system/model/dto"
	"github.com/MjSteed/vue3-element-admin-go/system/service"
	"github.com/gin-gonic/gin"
)

type MenuApi struct{}

// 资源(菜单+权限)列表
// @Router    /api/v1/menus/resources [get]
func (a MenuApi) ListResources(c *gin.Context) {
	service := service.MenuService{}
	vo.SuccessData(service.ListResources(), c)
}

// 菜单列表
// @Router    /api/v1/menus [get]
func (a MenuApi) List(c *gin.Context) {
	var pageParam dto.DeptPageReq
	err := c.ShouldBindQuery(&pageParam)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	service := service.MenuService{}
	list, err := service.ListPages(pageParam)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	vo.SuccessData(list, c)
}

// 菜单下拉列表
// @Router    /api/v1/menus/options [get]
func (a MenuApi) ListOptions(c *gin.Context) {
	service := service.MenuService{}
	vo.SuccessData(service.ListOptions(), c)
}

// 路由列表
// @Router    /api/v1/menus/routes [get]
func (a MenuApi) ListRoutes(c *gin.Context) {
	service := service.MenuService{}
	vo.SuccessData(service.ListRoutes(), c)
}

// 新增菜单
// @Router    /api/v1/menus [post]
func (a MenuApi) Save(c *gin.Context) {
	var d dto.MenuForm
	err := c.ShouldBindJSON(d)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	service := service.MenuService{}
	err = service.Save(d.ToMenu())
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	vo.Success(c)
}

// 修改菜单
// @Router    /api/v1/menus [put]
func (a MenuApi) Update(c *gin.Context) {
	var d dto.MenuForm
	err := c.ShouldBindJSON(d)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	service := service.MenuService{}
	err = service.Save(d.ToMenu())
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	vo.Success(c)
}

// 删除部门
// @Router    /api/v1/dept/{ids} [delete]
func (a MenuApi) BatchDelete(c *gin.Context) {
	idsStr := strings.Split(c.Param("ids"), ",")
	ids := make([]int64, len(idsStr))
	for _, v := range idsStr {
		id, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			continue
		}
		ids = append(ids, id)
	}
	service := service.MenuService{}
	err := service.DeleteByIds(ids)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	vo.Success(c)
}
