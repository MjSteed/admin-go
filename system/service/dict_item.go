package service

import (
	"context"
	"github.com/MjSteed/vue3-element-admin-go/system/dao"
	"github.com/MjSteed/vue3-element-admin-go/system/model"
	"github.com/MjSteed/vue3-element-admin-go/system/model/dto"
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
func (s *DictItemService) Create(ctx context.Context, data *model.SysDictItem) (err error) {
	err = s.dictItemDao.Create(ctx, data)
	return err
}

// Update 更新字典数据项
func (s *DictItemService) Update(ctx context.Context, sysDictItem *model.SysDictItem) (err error) {
	err = s.dictItemDao.Update(ctx, sysDictItem)
	return err
}

// UpdateOldCodeToNew 字典类型code变化，同步修改字典项的类型code
// @param oldCode 旧code
// @param newCode 新code
func (s *DictItemService) UpdateOldCodeToNew(ctx context.Context, oldCode string, newCode string) (err error) {
	if oldCode == newCode {
		return nil
	}
	return s.dictItemDao.UpdateOldCodeToNew(ctx, oldCode, newCode)
}

// Delete 删除字典数据项
// @param: ids 待删除的字典数据项ID
func (s *DictItemService) Delete(ctx context.Context, ids []int64) (err error) {
	if len(ids) <= 0 {
		s.log.Debug("删除数据为空")
		return nil
	}
	return s.dictItemDao.Delete(ctx, ids)
}

// DeleteByCode 删除字典数据项
// @param: codes 待删除的字典数据项编码
func (s *DictItemService) DeleteByCode(ctx context.Context, codes []string) (err error) {
	if len(codes) <= 0 {
		s.log.Debug("删除数据为空")
		return nil
	}
	return s.dictItemDao.DeleteByCodes(ctx, codes)
}

// ListPages 字典数据项分页列表
func (s *DictItemService) ListPages(ctx context.Context, pageReq dto.DictItemPageReq) (list []model.SysDictItem, total int64, err error) {
	list, total, err = s.dictItemDao.ListPages(ctx, &pageReq)
	return list, total, err
}

// GetById 字典数据项表单详情
// @param id 字典数据项ID
func (s *DictItemService) GetById(ctx context.Context, id int64) (model.SysDictItem, error) {
	dictItem, err := s.dictItemDao.FindByID(ctx, id)
	return *dictItem, err
}

// GetByCode 根据code获取字典数据项列表
// @param code 字典数据项编码
func (s *DictItemService) GetByCode(ctx context.Context, code string) ([]model.SysDictItem, error) {
	return s.dictItemDao.GetByCode(ctx, code)
}
