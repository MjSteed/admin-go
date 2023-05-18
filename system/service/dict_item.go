package service

import (
	"errors"

	"github.com/MjSteed/vue3-element-admin-go/common"
	"github.com/MjSteed/vue3-element-admin-go/system/model"
	"github.com/MjSteed/vue3-element-admin-go/system/model/dto"
	"gorm.io/gorm"
)

//@description: 创建字典数据

type dictItemService struct {
}

var DictItemService = new(dictItemService)

func (dictItemService *dictItemService) SaveDictItem(data *model.SysDictItem) (err error) {
	err = common.DB.Create(&data).Error
	return err
}

func (dictItemService *dictItemService) UpdateDictItem(sysDictItem *model.SysDictItem) (err error) {
	err = common.DB.Updates(&sysDictItem).Error
	return err
}

// 字典类型code变化，同步修改字典项的类型code
// @param oldCode 旧code
// @param newCode 新code
func (dictItemService *dictItemService) UpdateOldCodeToNew(oldCode string, newCode string) (err error) {
	if oldCode == newCode {
		return nil
	}
	return common.DB.Model(&model.SysDictItem{}).Where("type_code = ?", oldCode).Updates(model.SysDictItem{TypeCode: newCode}).Error
}

// @param: ids 待删除的字典数据项ID
func (dictItemService *dictItemService) DeleteDictItems(ids []int64) (err error) {
	if len(ids) <= 0 {
		return errors.New("删除数据为空")
	}
	tx := common.DB.Delete(&model.SysDictItem{}, ids)
	if err = tx.Error; err != nil {
		err = tx.Error
		return err
	}
	if tx.RowsAffected == 0 {
		err = errors.New("删除失败")
		return err
	}
	return nil
}

// @param: ids 待删除的字典数据项ID
func (dictItemService *dictItemService) DeleteDictItemsByCode(tx *gorm.DB, codes []string) (err error) {
	if len(codes) <= 0 {
		return errors.New("删除数据为空")
	}
	err = tx.Where("type_code in ?", codes).Delete(&model.SysDictItem{}).Error
	if err != nil {
		err = tx.Error
		return err
	}
	return nil
}

// 字典数据项分页列表
func (dictItemService *dictItemService) ListDictItemPages(pageReq dto.DictItemPageReq) (list []model.SysDictItem, total int64, err error) {
	tx := common.DB.Model(&model.SysDictItem{})
	if pageReq.Name != "" {
		tx = tx.Where("`name` like ?", "%"+pageReq.Name+"%")
	}
	if pageReq.TypeCode != "" {
		tx = tx.Where("`type_code` like ?", "%"+pageReq.TypeCode+"%")
	}
	err = tx.Count(&total).Error
	if err != nil {
		return
	}
	err = tx.Limit(pageReq.PageSize).Offset(pageReq.PageSize * (pageReq.PageNum - 1)).Find(&list).Error
	return list, total, err
}

// 字典数据项表单详情
// @param id 字典数据项ID
func (dictItemService *dictItemService) GetDictItem(id int64) (dictItem model.SysDictItem, err error) {
	err = common.DB.Model(&dictItem).First(&dictItem, id).Error
	return dictItem, err
}
