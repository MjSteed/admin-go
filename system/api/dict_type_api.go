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

type DictTypeApi struct{}

// 字典类型分页列表
// @Router    /api/v1/dict/types/pages [get]
func (a DictTypeApi) ListPages(c *gin.Context) {
	var pageParam dto.DictTypePageReq
	err := c.ShouldBindQuery(&pageParam)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	service := service.DictTypeService{}
	list, total, err := service.ListDictTypePages(pageParam)
	if err != nil {
		vo.FailMsg("查询失败", c)
		return
	}
	vo.SuccessData(vo.PageResult{List: list, Total: total}, c)
}

// 字典数据表单数据
// @Router    /api/v1/dict/types/{id}/form [get]
func (a DictTypeApi) GetForm(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	service := service.DictTypeService{}
	dictItem, err := service.GetDictType(id)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	vo.SuccessData(dictItem, c)
}

// 新增字典数据
// @Router    /api/v1/dict/types [post]
func (a DictTypeApi) Save(c *gin.Context) {
	var d model.SysDictType
	err := c.ShouldBindJSON(d)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	service := service.DictTypeService{}
	err = service.SaveDictType(d)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	vo.Success(c)
}

// 修改字典数据
// @Router    /api/v1/dict/types [put]
func (a DictTypeApi) Update(c *gin.Context) {
	var d model.SysDictType
	err := c.ShouldBindJSON(d)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	service := service.DictTypeService{}
	err = service.UpdateDictType(d)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	vo.Success(c)
}

// 删除字典
// @Router    /api/v1/dict/types/{ids} [delete]
func (a DictTypeApi) BatchDelete(c *gin.Context) {
	idsStr := strings.Split(c.Param("ids"), ",")
	ids := make([]int64, len(idsStr))
	for _, v := range idsStr {
		id, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			continue
		}
		ids = append(ids, id)
	}
	service := service.DictTypeService{}
	err := service.DeleteDictTypes(ids)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	vo.Success(c)
}

// 删除字典
// @Router    /api/v1/dict/types/{typeCode}/items [get]
func (a DictTypeApi) ListDictItemsByTypeCode(c *gin.Context) {
	typeCode := c.Param("typeCode")
	service := service.DictTypeService{}
	dicts, err := service.ListDictItemsByTypeCode(typeCode)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	vo.SuccessData(dicts, c)
}
