package dao

import (
	"context"

	"github.com/MjSteed/vue3-element-admin-go/system/model"
	"github.com/MjSteed/vue3-element-admin-go/system/model/dto"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type MenuDao struct {
	db  *gorm.DB
	log *zap.Logger
}

// NewMenuDao 实例化
func NewMenuDao(db *gorm.DB, log *zap.Logger) *MenuDao {
	return &MenuDao{db: db, log: log}
}

// Lists 列表
// @param pageReq 分页参数
// @return []SysMenu
// @return error
func (d *MenuDao) Lists(ctx context.Context, pageReq dto.DeptPageReq) ([]model.SysMenu, error) {
	tx := d.db.WithContext(ctx).Model(&model.SysMenu{})
	if pageReq.Keywords != "" {
		tx = tx.Where("`name` like ?", "%"+pageReq.Keywords+"%")
	}
	var menus []model.SysMenu
	err := tx.Order("`sort` ASC").Find(&menus).Error
	if err != nil {
		d.log.Warn("查询菜单失败", zap.Any("pageReq", pageReq), zap.Error(err))
		return nil, err
	}
	return menus, nil
}

// FindByID 根据ID查询
// @param id 主键ID
// @return *SysMenu
// @return error
func (d *MenuDao) FindByID(ctx context.Context, id int64) (*model.SysMenu, error) {
	var menu model.SysMenu
	err := d.db.WithContext(ctx).First(&menu, id).Error
	if err != nil {
		d.log.Warn("查询菜单失败", zap.Int64("id", id), zap.Error(err))
		return nil, err
	}
	return &menu, nil
}

// Save 保存
// @param menu SysMenu
// @return *SysMenu
// @return error
func (d *MenuDao) Save(ctx context.Context, menu model.SysMenu) (*model.SysMenu, error) {
	err := d.db.WithContext(ctx).Create(&menu).Error
	if err != nil {
		d.log.Warn("保存菜单失败", zap.Any("menu", menu), zap.Error(err))
		return nil, err
	}
	return &menu, nil
}

// ListAll 查询所有
// @return []SysMenu
// @return error
func (d *MenuDao) ListAll(ctx context.Context) ([]model.SysMenu, error) {
	var menus []model.SysMenu
	err := d.db.WithContext(ctx).Order("`sort` ASC").Find(&menus).Error
	if err != nil {
		d.log.Warn("查询菜单失败", zap.Error(err))
		return nil, err
	}
	return menus, nil
}

// Delete 删除
// @param ids 主键ID集合
// @return error
func (d *MenuDao) Delete(ctx context.Context, ids []int64) error {
	err := d.db.WithContext(ctx).Delete(&model.SysMenu{}, ids).Error
	if err != nil {
		d.log.Warn("删除菜单失败", zap.Any("ids", ids), zap.Error(err))
		return err
	}
	return nil
}

// UpdateVisible 更新菜单可见性
// @param id 主键ID
// @param visible 可见性 (1->显示；2->隐藏)
// @return error
func (d *MenuDao) UpdateVisible(ctx context.Context, id int64, visible int) error {
	err := d.db.WithContext(ctx).Model(&model.SysMenu{}).Where("id = ?", id).Update("visible", visible).Error
	if err != nil {
		d.log.Warn("更新菜单可见性失败", zap.Int64("id", id), zap.Int("visible", visible), zap.Error(err))
		return err
	}
	return nil
}

// GetPermsByRoles 根据角色ID集合查询权限标识集合
// @param roleCodes 角色code集合
// @return []string
// @return error
func (d *MenuDao) GetPermsByRoles(ctx context.Context, roleCodes []string) ([]string, error) {
	sql := `
	SELECT
            DISTINCT t1.perm
        FROM
            sys_menu t1
                INNER JOIN sys_role_menu t2
                INNER JOIN sys_role t3
        WHERE
            t1.type = ?
          AND t1.perm IS NOT NULL
	`
	var condition string
	if len(roleCodes) > 0 {
		condition += " AND t3.CODE IN ("
		for i, v := range roleCodes {
			if i != 0 {
				condition += ","
			}
			condition += "'" + v + "'"
		}
		condition += ")"
	} else {
		condition += "AND t1.id = -1"
	}
	var perms []string
	err := d.db.Raw(sql+condition, 4).Scan(&perms).Error
	if err != nil {
		d.log.Warn("查询权限标识失败", zap.Any("roleCodes", roleCodes), zap.Error(err))
		return nil, err
	}
	return perms, nil
}

// ListRoutes 查询路由
// @return []SysMenu
// @return error
func (d *MenuDao) ListRoutes(ctx context.Context) ([]model.SysMenu, error) {
	var menus []model.SysMenu
	err := d.db.WithContext(ctx).Where("type != ?", model.BUTTON).Preload("SysRoles").Find(&menus).Error
	if err != nil {
		d.log.Warn("查询路由失败", zap.Error(err))
		return nil, err
	}
	return menus, nil
}
