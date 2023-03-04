package service

import (
	"errors"

	"github.com/MjSteed/vue3-element-admin-go/common"
	"github.com/MjSteed/vue3-element-admin-go/system/model"
	"github.com/MjSteed/vue3-element-admin-go/system/model/dto"
	"gorm.io/gorm"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteSysDictionary
//@description: 创建字典数据
//@param: sysDictionary model.SysDictionary
//@return: err error

type DictItemService struct {
}

func (dictItemService *DictItemService) SaveDictItem(sysDictItem model.SysDictItem) (err error) {
	if !errors.Is(common.DB.First(&model.SysDictItem{}, "type_code = ?", sysDictItem.TypeCode).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同的type，不允许创建")
	}
	err = common.DB.Create(&sysDictItem).Error
	return err
}

func (dictItemService *DictItemService) UpdateDictItem(sysDictItem model.SysDictItem) (err error) {
	if !errors.Is(common.DB.First(&model.SysDictItem{}, sysDictItem.Id).Error, gorm.ErrRecordNotFound) {
		return errors.New("改字典不存在，无法更新")
	}
	err = common.DB.Save(&sysDictItem).Error
	return err
}

// @param: ids 待删除的字典数据项ID
func (dictItemService *DictItemService) DeleteDictItems(ids []int64) (err error) {
	if ids == nil || len(ids) <= 0 {
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

// 字典数据项分页列表
func (dictItemService *DictItemService) ListDictItemPages(pageReq dto.DictItemPageReq) (list []model.SysDictItem, total int64, err error) {
	tx := common.DB.Model(&model.SysDictItem{})
	if pageReq.Keywords != "" {
		tx = tx.Where("`name` like ?", "%"+pageReq.Keywords+"%")
	}
	if pageReq.TypeCode != "" {
		tx = tx.Where("`typeCode` like ?", "%"+pageReq.TypeCode+"%")
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
func (dictItemService *DictItemService) GetDictItem(id int64) (dictItem model.SysDictItem, err error) {
	err = common.DB.Model(&dictItem).First(&dictItem, id).Error
	return dictItem, err
}
