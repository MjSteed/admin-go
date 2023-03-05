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
}

// 获取部门列表
// @Router    /api/v1/dept [get]
func (a DeptApi) ListPages(c *gin.Context) {
	var pageParam dto.DeptPageReq
	err := c.ShouldBindQuery(&pageParam)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	service := service.DeptService{}
	list, total, err := service.ListDepts(pageParam)
	if err != nil {
		vo.FailMsg("查询失败", c)
		return
	}
	vo.SuccessData(vo.PageResult{List: list, Total: total}, c)
}

// 获取部门下拉选项
// @Router    /api/v1/dept/options [get]
func (a DeptApi) ListOptions(c *gin.Context) {
	service := service.DeptService{}
	list, err := service.ListDeptOptions()
	if err != nil {
		vo.FailMsg("查询失败", c)
		return
	}
	vo.SuccessData(list, c)
}

// 获取部门详情
// @Router    /api/v1/dept/{id}/form [get]
func (a DeptApi) GetForm(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	service := service.DeptService{}
	item, err := service.GetDeptForm(id)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	vo.SuccessData(item, c)
}

// 新增部门
// @Router    /api/v1/dept [post]
func (a DeptApi) Save(c *gin.Context) {
	var d model.SysDept
	err := c.ShouldBindJSON(d)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	service := service.DeptService{}
	id, err := service.SaveDept(d)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	vo.SuccessData(id, c)
}

// 修改部门
// @Router    /api/v1/dept [put]
func (a DeptApi) Update(c *gin.Context) {
	var d model.SysDept
	err := c.ShouldBindJSON(d)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	service := service.DeptService{}
	id, err := service.UpdateDept(d)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	vo.SuccessData(id, c)
}

// 删除部门
// @Router    /api/v1/dept/{ids} [delete]
func (a DeptApi) BatchDelete(c *gin.Context) {
	idsStr := strings.Split(c.Param("ids"), ",")
	ids := make([]int64, len(idsStr))
	for _, v := range idsStr {
		id, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			continue
		}
		ids = append(ids, id)
	}
	service := service.DeptService{}
	err := service.DeleteByIds(ids)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	vo.Success(c)
}
