package dao

import (
	"context"

	"github.com/MjSteed/vue3-element-admin-go/system/model"
	"github.com/MjSteed/vue3-element-admin-go/system/model/dto"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// DictItemDao 字典项
type DictItemDao struct {
	db  *gorm.DB
	log *zap.Logger
}

// NewDictItemDao 实例化
func NewDictItemDao(db *gorm.DB, log *zap.Logger) *DictItemDao {
	return &DictItemDao{db: db, log: log}
}

// Create 创建
// @param dictItem SysDictItem
// @return error
func (d *DictItemDao) Create(ctx context.Context, dictItem *model.SysDictItem) error {
	err := d.db.WithContext(ctx).Create(dictItem).Error
	if err != nil {
		d.log.Warn("创建字典项失败", zap.Any("dictItem", dictItem), zap.Error(err))
		return err
	}
	return nil
}

// Update 更新
// @param dictItem SysDictItem
// @return error
func (d *DictItemDao) Update(ctx context.Context, dictItem *model.SysDictItem) error {
	err := d.db.WithContext(ctx).Model(dictItem).Updates(dictItem).Error
	if err != nil {
		d.log.Warn("更新字典项失败", zap.Any("dictItem", dictItem), zap.Error(err))
		return err
	}
	return nil
}

// UpdateOldCodeToNew 更新旧编码为新编码
// @param oldCode 旧编码
// @param newCode 新编码
// @return error
func (d *DictItemDao) UpdateOldCodeToNew(ctx context.Context, oldCode, newCode string) error {
	err := d.db.WithContext(ctx).Model(&model.SysDictItem{}).Where("type_code = ?", oldCode).Update("type_code", newCode).Error
	if err != nil {
		d.log.Warn("更新字典项失败", zap.String("oldCode", oldCode), zap.String("newCode", newCode), zap.Error(err))
		return err
	}
	return nil
}

// Delete 删除
// @param ids []int64
// @return error
func (d *DictItemDao) Delete(ctx context.Context, ids []int64) error {
	err := d.db.WithContext(ctx).Delete(&model.SysDictItem{}, ids).Error
	if err != nil {
		d.log.Warn("删除字典项失败", zap.Any("ids", ids), zap.Error(err))
		return err
	}
	return nil
}

// DeleteByCodes 根据字典编码删除
// @param codes []string
// @return error
func (d *DictItemDao) DeleteByCodes(ctx context.Context, codes []string) error {
	err := d.db.WithContext(ctx).Where("type_code IN ?", codes).Delete(&model.SysDictItem{}).Error
	if err != nil {
		d.log.Warn("删除字典项失败", zap.Any("codes", codes), zap.Error(err))
		return err
	}
	return nil
}

// FindByID 根据ID查询
// @param id 主键ID
// @return *SysDictItem
// @return error
func (d *DictItemDao) FindByID(ctx context.Context, id int64) (*model.SysDictItem, error) {
	var dictItem model.SysDictItem
	err := d.db.WithContext(ctx).First(&dictItem, id).Error
	if err != nil {
		d.log.Warn("查询字典项失败", zap.Int64("id", id), zap.Error(err))
		return nil, err
	}
	return &dictItem, nil
}

// ListPages 分页查询
// @param  pageReq 分页参数
// @return []SysDictItem 列表
// @return total 总数
// @return error
func (d *DictItemDao) ListPages(ctx context.Context, pageReq *dto.DictItemPageReq) ([]model.SysDictItem, int64, error) {
	var dictItems []model.SysDictItem
	var total int64
	tx := d.db.WithContext(ctx).Model(&model.SysDictItem{})
	if pageReq.Name != "" {
		tx = tx.Where("`name` like ?", "%"+pageReq.Name+"%")
	}
	if pageReq.TypeCode != "" {
		tx = tx.Where("`type_code` like ?", "%"+pageReq.TypeCode+"%")
	}
	err := tx.Count(&total).Error
	if err != nil {
		d.log.Warn("查询字典项总数失败", zap.Any("pageReq", pageReq), zap.Error(err))
		return nil, 0, err
	}
	err = tx.Limit(pageReq.PageSize).Offset(pageReq.PageSize * (pageReq.PageNum - 1)).Find(&dictItems).Error
	if err != nil {
		d.log.Warn("查询字典项失败", zap.Any("pageReq", pageReq), zap.Error(err))
		return nil, 0, err
	}
	return dictItems, total, nil
}

// GetByCode 根据编码查询
// @param code 编码
// @return []SysDictItem
// @return error
func (d *DictItemDao) GetByCode(ctx context.Context, code string) ([]model.SysDictItem, error) {
	var dictItems []model.SysDictItem
	err := d.db.WithContext(ctx).Where("type_code = ?", code).Find(&dictItems).Error
	if err != nil {
		d.log.Warn("查询字典项失败", zap.String("code", code), zap.Error(err))
		return nil, err
	}
	return dictItems, nil
}
