package service

import (
	"github.com/MjSteed/vue3-element-admin-go/system/dao"
	"github.com/MjSteed/vue3-element-admin-go/system/model"
	"github.com/MjSteed/vue3-element-admin-go/system/model/dto"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// DictItemService 字典数据项服务
// @description: 创建字典数据
type DictItemService struct {
	log         *zap.Logger
	dictItemDao *dao.DictItemDao
}

// NewDictItemService 实例化
func NewDictItemService(log *zap.Logger, dictItemDao *dao.DictItemDao) *DictItemService {
	return &DictItemService{log: log, dictItemDao: dictItemDao}
}

// Create 创建字典数据项
func (s *DictItemService) Create(ctx *gin.Context, data *model.SysDictItem) (err error) {
	err = s.dictItemDao.Create(ctx, data)
	return err
}

// Update 更新字典数据项
func (s *DictItemService) Update(ctx *gin.Context, sysDictItem *model.SysDictItem) (err error) {
	err = s.dictItemDao.Update(ctx, sysDictItem)
	return err
}

// UpdateOldCodeToNew 字典类型code变化，同步修改字典项的类型code
// @param oldCode 旧code
// @param newCode 新code
func (s *DictItemService) UpdateOldCodeToNew(ctx *gin.Context, oldCode string, newCode string) (err error) {
	if oldCode == newCode {
		return nil
	}
	return s.dictItemDao.UpdateOldCodeToNew(ctx, oldCode, newCode)
}

// Delete 删除字典数据项
// @param: ids 待删除的字典数据项ID
func (s *DictItemService) Delete(ctx *gin.Context, ids []int64) (err error) {
	if len(ids) <= 0 {
		s.log.Debug("删除数据为空")
		return nil
	}
	return s.dictItemDao.Delete(ctx, ids)
}

// DeleteByCode 删除字典数据项
// @param: codes 待删除的字典数据项编码
func (s *DictItemService) DeleteByCode(ctx *gin.Context, codes []string) (err error) {
	if len(codes) <= 0 {
		s.log.Debug("删除数据为空")
		return nil
	}
	return s.dictItemDao.DeleteByCodes(ctx, codes)
}

// ListPages 字典数据项分页列表
func (s *DictItemService) ListPages(ctx *gin.Context, pageReq dto.DictItemPageReq) (list []model.SysDictItem, total int64, err error) {
	list, total, err = s.dictItemDao.ListPages(ctx, &pageReq)
	return list, total, err
}

// GetById 字典数据项表单详情
// @param id 字典数据项ID
func (s *DictItemService) GetById(ctx *gin.Context, id int64) (model.SysDictItem, error) {
	dictItem, err := s.dictItemDao.FindByID(ctx, id)
	return *dictItem, err
}
