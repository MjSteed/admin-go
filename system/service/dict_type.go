package service

import (
	"errors"

	"github.com/MjSteed/vue3-element-admin-go/common"
	"github.com/MjSteed/vue3-element-admin-go/system/model"
	"github.com/MjSteed/vue3-element-admin-go/system/model/dto"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

//@description: 创建字典数据

type DictTypeService struct {
}

func (dictTypeService *DictTypeService) SaveDictType(sysdictType model.SysDictType) (err error) {
	err = common.DB.Create(&sysdictType).Error
	return err
}

func (dictTypeService *DictTypeService) UpdateDictType(sysdictType model.SysDictType) (err error) {
	if !errors.Is(common.DB.First(&model.SysDictType{}, sysdictType.Id).Error, gorm.ErrRecordNotFound) {
		return errors.New("字典类型不存在，无法更新")
	}
	oldCode := sysdictType.Code
	err = common.DB.Save(&sysdictType).Error
	if err != nil || oldCode == sysdictType.Code {
		return err
	}
	//字典类型code变化，同步修改字典项的类型code
	dictItemService := DictItemService{}
	return dictItemService.UpdateOldCodeToNew(oldCode, sysdictType.Code)
}

// @param: ids 待删除的字典类型ID
func (dictTypeService *DictTypeService) DeleteDictTypes(ids []int64) (err error) {
	if ids == nil || len(ids) <= 0 {
		return errors.New("删除数据为空")
	}
	var types []model.SysDictType
	tx := common.DB.Clauses(clause.Returning{Columns: []clause.Column{{Name: "code"}}}).Delete(&types, ids)
	if err = tx.Error; err != nil {
		err = tx.Error
		return err
	}
	if tx.RowsAffected == 0 {
		err = errors.New("删除失败")
		return err
	}
	if len(types) <= 0 {
		return err
	}
	codes := make([]string, len(types))
	for _, v := range types {
		codes = append(codes, v.Code)
	}
	//删除字典数据项
	dictItemService := DictItemService{}
	return dictItemService.DeleteDictItemsByCode(codes)
}

// 字典分页列表
func (dictTypeService *DictTypeService) ListDictTypePages(pageReq dto.DictTypePageReq) (list []model.SysDictType, total int64, err error) {
	tx := common.DB.Model(&model.SysDictType{})
	if pageReq.Keywords != "" {
		tx = tx.Where("`name` like ?", "%"+pageReq.Keywords+"%").Or("`code` like ?", "%"+pageReq.Keywords+"%")
	}
	err = tx.Count(&total).Error
	if err != nil {
		return
	}
	err = tx.Limit(pageReq.PageSize).Offset(pageReq.PageSize * (pageReq.PageNum - 1)).Find(&list).Error
	return list, total, err
}

// 获取字典类型表单详情
// @param id 字典类型ID
func (dictTypeService *DictTypeService) GetDictType(id int64) (dictType model.SysDictType, err error) {
	err = common.DB.Model(&dictType).First(&dictType, id).Error
	return dictType, err
}

// 获取字典类型的数据项
// @param typeCode
func (dictTypeService *DictTypeService) ListDictItemsByTypeCode(typeCode string) (dicts []model.SysDictItem, err error) {
	err = common.DB.Model(&model.SysDictItem{}).Where("type_code = ?", typeCode).Find(dicts).Error
	if err != nil {
		return
	}
	return
}
