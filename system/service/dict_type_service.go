package service

import (
	"errors"

	"github.com/MjSteed/vue3-element-admin-go/common"
	"github.com/MjSteed/vue3-element-admin-go/system/model"
	"github.com/MjSteed/vue3-element-admin-go/system/model/dto"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

//@description: 创建字典数据

type dictTypeService struct {
}

var DictTypeService = new(dictTypeService)

func (dictTypeService *dictTypeService) SaveDictType(sysdictType *model.SysDictType) (err error) {
	err = common.DB.Create(&sysdictType).Error
	return err
}

func (dictTypeService *dictTypeService) UpdateDictType(data *model.SysDictType) (err error) {
	oldData := model.SysDictType{Id: data.Id}
	err = common.DB.First(&oldData).Error
	common.LOG.Debug("查询字典数据", zap.Any("旧数据", oldData))
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("字典类型不存在，无法更新")
	}
	oldCode := oldData.Code
	err = common.DB.Updates(&data).Error
	if err != nil || oldCode == data.Code {
		return err
	}
	//字典类型code变化，同步修改字典项的类型code
	return DictItemService.UpdateOldCodeToNew(oldCode, data.Code)
}

// @param: ids 待删除的字典类型ID
func (dictTypeService *dictTypeService) DeleteDictTypes(ids []int64) (err error) {
	if ids == nil || len(ids) <= 0 {
		return errors.New("删除数据为空")
	}

	err = common.DB.Transaction(func(tx *gorm.DB) error {
		var codes []string
		dictType := model.SysDictType{}
		err = tx.Model(&dictType).Where("id in ?", ids).Select("code").Find(&codes).Error
		if err != nil {
			return err
		}
		err = tx.Where("id in ?", ids).Delete(&dictType).Error
		if err != nil {
			return err
		}
		err = DictItemService.DeleteDictItemsByCode(tx, codes)
		return err
	})
	return
}

// 字典分页列表
func (dictTypeService *dictTypeService) ListDictTypePages(pageReq dto.DictTypePageReq) (list []model.SysDictType, total int64, err error) {
	tx := common.DB.Model(&model.SysDictType{})
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

// 获取字典类型表单详情
// @param id 字典类型ID
func (dictTypeService *dictTypeService) GetDictType(id int64) (dictType model.SysDictType, err error) {
	err = common.DB.Model(&dictType).First(&dictType, id).Error
	return dictType, err
}

// 获取字典类型的数据项
// @param typeCode
func (dictTypeService *dictTypeService) ListDictItemsByTypeCode(typeCode string) (dicts []model.SysDictItem, err error) {
	err = common.DB.Model(&model.SysDictItem{}).Where("type_code = ?", typeCode).Find(dicts).Error
	if err != nil {
		return
	}
	return
}
