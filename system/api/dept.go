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

type DeptApi struct {
	deptSerivce *service.DeptService
}

// 获取部门列表
// @Router    /api/v1/dept [get]
func (a *DeptApi) ListPages(c *gin.Context) {
	var pageParam dto.DeptPageReq
	err := c.ShouldBindQuery(&pageParam)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	list, err := a.deptSerivce.ListPages(c, pageParam)
	if err != nil {
		vo.FailMsg("查询失败", c)
		return
	}
	vo.SuccessData(list, c)
}

// 获取部门下拉选项
// @Router    /api/v1/dept/options [get]
func (a *DeptApi) ListOptions(c *gin.Context) {
	list, err := a.deptSerivce.Options(c)
	if err != nil {
		vo.FailMsg("查询失败", c)
		return
	}
	vo.SuccessData(list, c)
}

// 获取部门详情
// @Router    /api/v1/dept/{id}/form [get]
func (a *DeptApi) GetForm(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	item, err := a.deptSerivce.GetForm(c, id)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	vo.SuccessData(item, c)
}

// 新增部门
// @Router    /api/v1/dept [post]
func (a *DeptApi) Save(c *gin.Context) {
	var d model.SysDept
	err := c.ShouldBindJSON(&d)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	id, err := a.deptSerivce.Create(c, &d)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	vo.SuccessData(id, c)
}

// 修改部门
// @Router    /api/v1/dept [put]
func (a *DeptApi) Update(c *gin.Context) {
	var d model.SysDept
	err := c.ShouldBindJSON(&d)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	id, err := a.deptSerivce.Update(c, &d)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	vo.SuccessData(id, c)
}

// 删除部门
// @Router    /api/v1/dept/{ids} [delete]
func (a *DeptApi) BatchDelete(c *gin.Context) {
	idsStr := strings.Split(c.Param("ids"), ",")
	ids := make([]int64, len(idsStr))
	for _, v := range idsStr {
		id, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			continue
		}
		ids = append(ids, id)
	}
	err := a.deptSerivce.DeleteByIds(c, ids)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	vo.Success(c)
}
