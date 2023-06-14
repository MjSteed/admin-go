package dao

import (
	"context"
	"errors"
	"github.com/MjSteed/vue3-element-admin-go/system/model"
	"github.com/MjSteed/vue3-element-admin-go/system/model/dto"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// DictTypeDao 字典类型
type DictTypeDao struct {
	db  *gorm.DB
	log *zap.Logger
}

// NewDictTypeDao 实例化
func NewDictTypeDao(db *gorm.DB, log *zap.Logger) *DictTypeDao {
	return &DictTypeDao{db: db, log: log}
}

// Create 创建
// @param dictType SysDictType
// @return error
func (d *DictTypeDao) Create(ctx context.Context, dictType *model.SysDictType) error {
	err := d.db.WithContext(ctx).Create(dictType).Error
	if err != nil {
		d.log.Warn("创建字典类型失败", zap.Any("dictType", dictType), zap.Error(err))
		return err
	}
	return nil
}

// Update 更新
// @param dictType SysDictType
// @return error
func (d *DictTypeDao) Update(ctx context.Context, dictType *model.SysDictType) error {
	err := d.db.WithContext(ctx).Updates(dictType).Error
	if err != nil {
		d.log.Warn("更新字典类型失败", zap.Any("dictType", dictType), zap.Error(err))
		return err
	}
	return nil
}

// Delete 删除
// @param ids []int64
// @return error
func (d *DictTypeDao) Delete(ctx context.Context, ids []int64) error {
	if ids == nil || len(ids) <= 0 {
		return errors.New("删除数据为空")
	}
	err := d.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var codes []string
		dictType := model.SysDictType{}
		err := tx.Model(&dictType).Where("id in ?", ids).Select("code").Find(&codes).Error
		if err != nil {
			return err
		}
		err = tx.Where("id in ?", ids).Delete(&dictType).Error
		if err != nil {
			return err
		}
		//err = DictItemService.DeleteDictItemsByCode(tx, codes)
		return err
	})
	return err
}

// FindById 根据ID查询
// @param id int64
// @return *SysDictType
// @return error
func (d *DictTypeDao) FindById(ctx context.Context, id int64) (*model.SysDictType, error) {
	var dictType model.SysDictType
	err := d.db.WithContext(ctx).First(&dictType, id).Error
	if err != nil {
		d.log.Warn("查询字典类型失败", zap.Any("dictType", dictType), zap.Error(err))
		return nil, err
	}
	return &dictType, nil
}

// ListPages 分页查询
// @param pageReq dto.DictTypePageReq
// @return []SysDictType
// @return int64
// @return error
func (d *DictTypeDao) ListPages(ctx context.Context, pageReq *dto.DictTypePageReq) (list []model.SysDictType, total int64, err error) {
	tx := d.db.Model(&model.SysDictType{})
	if pageReq.Name != "" {
		tx = tx.Where("`name` like ?", "%"+pageReq.Name+"%").Or("`code` like ?", "%"+pageReq.Name+"%")
	}
	err = tx.Count(&total).Error
	if err != nil {
		return
	}
	err = tx.Limit(pageReq.PageSize).Offset(pageReq.PageSize * (pageReq.PageNum - 1)).Find(&list).Error
	return list, total, err
}
