package dao

import (
	"context"

	"github.com/MjSteed/vue3-element-admin-go/system/model"
	"github.com/MjSteed/vue3-element-admin-go/system/model/dto"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type DeptDao struct {
	db  *gorm.DB
	log *zap.Logger
}

func NewDeptDao(db *gorm.DB, log *zap.Logger) *DeptDao {
	return &DeptDao{db: db, log: log}
}

// Save 创建或更新部门
func (d *DeptDao) Save(ctx context.Context, dept *model.SysDept) (*model.SysDept, error) {
	err := d.db.WithContext(ctx).Create(*dept).Error
	if err != nil {
		d.log.Error("创建或更新部门失败", zap.Any("dept", dept), zap.Error(err))
		return nil, err
	}
	return dept, nil
}

// Delete 删除部门
func (d *DeptDao) Delete(ctx context.Context, ids []int64) error {
	err := d.db.WithContext(ctx).Delete(&model.SysDept{}, ids).Error
	if err != nil {
		d.log.Error("删除部门失败", zap.Int64s("ids", ids), zap.Error(err))
		return err
	}
	return nil
}

// ListPages 查询部门列表分页
func (d *DeptDao) ListPages(ctx context.Context, pageReq *dto.DeptPageReq) ([]model.SysDept, error) {
	tx := d.db.WithContext(ctx)
	if pageReq.Keywords != "" {
		tx = tx.Where("`name` like ?", "%"+pageReq.Keywords+"%")
	}
	if pageReq.Status != "" {
		tx = tx.Where("`status` = ?", pageReq.Status)
	}
	var depts []model.SysDept
	err := tx.Order("`sort` ASC").Find(&depts).Error
	if err != nil {
		d.log.Error("查询部门失败", zap.Any("pageReq", pageReq), zap.Error(err))
		return nil, err
	}
	return depts, nil
}

// ListAll 查询所有部门
func (d *DeptDao) ListAll(ctx context.Context) ([]model.SysDept, error) {
	var depts []model.SysDept
	err := d.db.WithContext(ctx).Where("`status` = 1").Order("`sort` ASC").Find(&depts).Error
	if err != nil {
		d.log.Error("查询所有部门失败", zap.Error(err))
		return nil, err
	}
	return depts, nil
}

// FindById 根据ID查询部门
func (d *DeptDao) FindById(ctx context.Context, id int64) (*model.SysDept, error) {
	var dept model.SysDept
	err := d.db.WithContext(ctx).Where("`id` = ?", id).First(&dept).Error
	if err != nil {
		d.log.Error("根据ID查询部门失败", zap.Int64("id", id), zap.Error(err))
		return nil, err
	}
	return &dept, nil
}
