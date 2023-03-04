package service

import (
	"errors"

	"github.com/MjSteed/vue3-element-admin-go/common"
	"github.com/MjSteed/vue3-element-admin-go/system/model"
	"github.com/MjSteed/vue3-element-admin-go/system/model/dto"
	"gorm.io/gorm"
)

//@description: 创建字典数据

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

// 字典类型code变化，同步修改字典项的类型code
// @param oldCode 旧code
// @param newCode 新code
func (dictItemService *DictItemService) UpdateOldCodeToNew(oldCode string, newCode string) (err error) {
	if oldCode == newCode {
		return nil
	}
	return common.DB.Model(&model.SysDictItem{}).Where("type_code = ?", oldCode).Updates(model.SysDictItem{TypeCode: newCode}).Error
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

// @param: ids 待删除的字典数据项ID
func (dictItemService *DictItemService) DeleteDictItemsByCode(codes []string) (err error) {
	if codes == nil || len(codes) <= 0 {
		return errors.New("删除数据为空")
	}
	tx := common.DB.Where("type_code in ?", codes).Delete(&model.SysDictItem{})
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
func (dictItemService *DictItemService) GetDictItem(id int64) (dictItem model.SysDictItem, err error) {
	err = common.DB.Model(&dictItem).First(&dictItem, id).Error
	return dictItem, err
}
