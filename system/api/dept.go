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

type deptApi struct {
}

var DeptApi = new(deptApi)

// 获取部门列表
// @Router    /api/v1/dept [get]
func (a *deptApi) ListPages(c *gin.Context) {
	var pageParam dto.DeptPageReq
	err := c.ShouldBindQuery(&pageParam)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	list, err := service.DeptService.ListDepts(pageParam)
	if err != nil {
		vo.FailMsg("查询失败", c)
		return
	}
	vo.SuccessData(list, c)
}

// 获取部门下拉选项
// @Router    /api/v1/dept/options [get]
func (a *deptApi) ListOptions(c *gin.Context) {
	list, err := service.DeptService.ListDeptOptions()
	if err != nil {
		vo.FailMsg("查询失败", c)
		return
	}
	vo.SuccessData(list, c)
}

// 获取部门详情
// @Router    /api/v1/dept/{id}/form [get]
func (a *deptApi) GetForm(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	item, err := service.DeptService.GetDeptForm(id)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	vo.SuccessData(item, c)
}

// 新增部门
// @Router    /api/v1/dept [post]
func (a *deptApi) Save(c *gin.Context) {
	var d model.SysDept
	err := c.ShouldBindJSON(&d)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	id, err := service.DeptService.SaveDept(&d)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	vo.SuccessData(id, c)
}

// 修改部门
// @Router    /api/v1/dept [put]
func (a *deptApi) Update(c *gin.Context) {
	var d model.SysDept
	err := c.ShouldBindJSON(&d)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	id, err := service.DeptService.UpdateDept(&d)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	vo.SuccessData(id, c)
}

// 删除部门
// @Router    /api/v1/dept/{ids} [delete]
func (a *deptApi) BatchDelete(c *gin.Context) {
	idsStr := strings.Split(c.Param("ids"), ",")
	ids := make([]int64, len(idsStr))
	for _, v := range idsStr {
		id, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			continue
		}
		ids = append(ids, id)
	}
	err := service.DeptService.DeleteByIds(ids)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	vo.Success(c)
}
