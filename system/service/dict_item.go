package service

import (
	"errors"

	"github.com/MjSteed/vue3-element-admin-go/common"
	"github.com/MjSteed/vue3-element-admin-go/system/model"
	"gorm.io/gorm"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteSysDictionary
//@description: 创建字典数据
//@param: sysDictionary model.SysDictionary
//@return: err error

type DictionaryService struct {
}

func (dictionaryService *DictionaryService) CreateSysDictionary(sysDictItem model.SysDictItem) (err error) {
	if !errors.Is(common.DB.First(&model.SysDictItem{}, "type_code = ?", sysDictItem.TypeCode).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同的type，不允许创建")
	}
	err = common.DB.Create(&sysDictItem).Error
	return err
}
