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

// DictItemApi 字典数据路由接口
type DictItemApi struct {
	dictItemService *service.DictItemService
}

// NewDictItemApi 实例化路由接口
func NewDictItemApi(dictItemService *service.DictItemService) DictItemApi {
	return DictItemApi{dictItemService: dictItemService}
}

// 字典数据分页列表
// @Router    /api/v1/dict/items/pages [get]
func (a *DictItemApi) ListPages(c *gin.Context) {
	var pageParam dto.DictItemPageReq
	err := c.ShouldBindQuery(&pageParam)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	list, total, err := a.dictItemService.ListPages(c, pageParam)
	if err != nil {
		vo.FailMsg("查询失败", c)
		return
	}
	vo.SuccessData(vo.PageResult{List: list, Total: total}, c)
}

// 字典数据表单数据
// @Router    /api/v1/dict/items/{id}/form [get]
func (a *DictItemApi) GetForm(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	dictItem, err := a.dictItemService.GetById(c, id)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	vo.SuccessData(dictItem, c)
}

// 新增字典数据
// @Router    /api/v1/dict/items [post]
func (a *DictItemApi) Save(c *gin.Context) {
	var dictItem model.SysDictItem
	err := c.ShouldBindJSON(&dictItem)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	err = a.dictItemService.Create(c, &dictItem)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	vo.Success(c)
}

// 修改字典数据
// @Router    /api/v1/dict/items [put]
func (a *DictItemApi) Update(c *gin.Context) {
	var dictItem model.SysDictItem
	err := c.ShouldBindJSON(&dictItem)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	err = a.dictItemService.Update(c, &dictItem)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	vo.Success(c)
}

// 删除字典
// @Router    /api/v1/dict/items/{ids} [delete]
func (a *DictItemApi) BatchDelete(c *gin.Context) {
	idsStr := strings.Split(c.Param("ids"), ",")
	ids := make([]int64, len(idsStr))
	for _, v := range idsStr {
		id, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			continue
		}
		ids = append(ids, id)
	}
	err := a.dictItemService.Delete(c, ids)
	if err != nil {
		vo.FailMsg(err.Error(), c)
		return
	}
	vo.Success(c)
}
