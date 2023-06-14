package service

import (
	"context"
	"errors"
	"github.com/MjSteed/vue3-element-admin-go/system/dao"

	"github.com/MjSteed/vue3-element-admin-go/common"
	"github.com/MjSteed/vue3-element-admin-go/common/model/vo"
	"github.com/MjSteed/vue3-element-admin-go/system/model"
	"github.com/MjSteed/vue3-element-admin-go/system/model/dto"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// DictTypeService 字典类型
type DictTypeService struct {
	log             *zap.Logger
	dictTypeDao     *dao.DictTypeDao
	dictItemService *DictItemService
}

// NewDictTypeService 实例化
func NewDictTypeService(log *zap.Logger, dictTypeDao *dao.DictTypeDao, dictItemService *DictItemService) *DictTypeService {
	return &DictTypeService{log: log, dictTypeDao: dictTypeDao, dictItemService: dictItemService}
}

// Create 创建
func (s *DictTypeService) Create(ctx context.Context, data *model.SysDictType) (err error) {
	err = s.dictTypeDao.Create(ctx, data)
	return err
}

// Update 更新
func (s *DictTypeService) Update(ctx context.Context, data *model.SysDictType) (err error) {
	oldData, err := s.dictTypeDao.FindById(ctx, data.Id)
	common.LOG.Debug("查询字典数据", zap.Any("旧数据", oldData))
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("字典类型不存在，无法更新")
	}
	oldCode := oldData.Code
	err = s.dictTypeDao.Update(ctx, data)
	if err != nil || oldCode == data.Code {
		return err
	}
	//字典类型code变化，同步修改字典项的类型code
	return s.dictItemService.UpdateOldCodeToNew(ctx, oldCode, data.Code)
}

// Delete 删除
// @param: ids 待删除的字典类型ID
func (s *DictTypeService) Delete(ctx context.Context, ids []int64) (err error) {
	if ids == nil || len(ids) <= 0 {
		return errors.New("删除数据为空")
	}
	return s.dictTypeDao.Delete(ctx, ids)
}

// ListPages 字典分页列表
func (s *DictTypeService) ListPages(ctx context.Context, pageReq dto.DictTypePageReq) (list []model.SysDictType, total int64, err error) {
	return s.dictTypeDao.ListPages(ctx, &pageReq)
}

// FindById 获取字典类型表单详情
// @param id 字典类型ID
func (s *DictTypeService) FindById(ctx context.Context, id int64) (model.SysDictType, error) {
	t, err := s.dictTypeDao.FindById(ctx, id)
	return *t, err
}

// ListDictItemsByTypeCode 获取字典类型的数据项
// @param typeCode
func (s *DictTypeService) ListDictItemsByTypeCode(ctx context.Context, typeCode string) (dicts []vo.TreeOption, err error) {
	list, err := s.dictItemService.GetByCode(ctx, typeCode)
	if err != nil {
		return
	}
	for _, v := range list {
		dicts = append(dicts, vo.TreeOption{Label: v.Name, Value: v.Value})
	}
	return
}
