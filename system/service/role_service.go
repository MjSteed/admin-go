package service

import (
	"errors"

	"github.com/MjSteed/vue3-element-admin-go/common"
	"github.com/MjSteed/vue3-element-admin-go/common/model/vo"
	"github.com/MjSteed/vue3-element-admin-go/system/model"
	"github.com/MjSteed/vue3-element-admin-go/system/model/dto"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// 角色业务接口层
type RoleService struct {
}

// 超级管理员角色编码
const ROOT_ROLE_CODE = "ROOT"

// 角色分页列表
func (service *RoleService) ListPages(pageReq dto.DeptPageReq) (list []model.SysRole, total int64, err error) {
	common.LOG.Debug("查询角色分页参数", zap.Int("PageNum", pageReq.PageNum), zap.Int("PageSize", pageReq.PageSize), zap.String("Keywords", pageReq.Keywords))
	tx := common.DB.Model(&list)
	if pageReq.Keywords != "" {
		tx = tx.Where("`name` like ?", "%"+pageReq.Keywords+"%").Or("`code` like ?", "%"+pageReq.Keywords+"%")
	}
	//TODO 超级管理员判断
	err = tx.Count(&total).Error
	if err != nil {
		return
	}
	err = tx.Limit(pageReq.PageSize).Offset(pageReq.PageSize * (pageReq.PageNum - 1)).Find(&list).Error
	return
}

// 角色下拉列表
func (service *RoleService) ListOptions() (list []vo.TreeOption) {
	var roles []model.SysRole
	err := common.DB.Model(&model.SysRole{}).Where("`code` != ?", ROOT_ROLE_CODE).Find(&roles).Error
	if err != nil {
		common.LOG.Error("查询角色下拉列表失败", zap.Error(err))
		return
	}
	for _, v := range roles {
		list = append(list, vo.TreeOption{Label: v.Name, Value: v.Id})
	}
	return
}

// 角色详情
func (service *RoleService) GetById(id int64) (r model.SysRole) {
	common.DB.First(&r, id)
	return
}

// 新增或更新
func (service *RoleService) Save(data model.SysRole) (err error) {
	tx := common.DB.Model(&data)
	if data.Id > 0 {
		tx = tx.Where("id != ?", data.Id)
	}
	var c int64
	err = tx.Where("`code` = ?", data.Code).Or("`name` = ?", data.Name).Count(&c).Error
	if err != nil {
		common.LOG.Error("查询角色是否重复失败", zap.Error(err))
		return
	}
	if c > 0 {
		err = errors.New("角色名称或角色编码重复，请检查！")
		return
	}
	err = common.DB.Model(&data).Save(&data).Error
	return
}

// 修改角色状态
func (service *RoleService) UpdateStatus(id int64, status int) (err error) {
	err = common.DB.Model(&model.SysRole{Id: id}).Update("status", status).Error
	return
}

// 批量删除
func (service *RoleService) DeleteByIds(ids []int64) error {
	var c int64
	common.DB.Model(&model.SysUserRole{}).Where("role_id in ?", ids).Count(&c)
	if c > 0 {
		return errors.New("角色已分配用户，无法删除，请检查！")
	}
	err := common.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Where("role_id in ?", ids).Delete(&model.SysRoleMenu{}).Error
		if err != nil {
			common.LOG.Warn("删除sys_user_role表失败", zap.Error(err))
			return err
		}
		err = tx.Where("id in ?", ids).Delete(&model.SysRole{}).Error
		return err
	})
	return err
}

// 获取角色的资源ID集合,资源包括菜单和权限
func (service *RoleService) GetRoleMenuIds(id int64) (menus []int64) {
	sql := `SELECT
				rm.menu_id
			FROM
				sys_role_menu rm
				INNER JOIN sys_menu m ON rm.menu_id = m.id
			WHERE rm.role_id = ?`
	common.DB.Raw(sql, id).Scan(&menus)
	return
}

// 修改角色的资源权限
func (service *RoleService) UpdateRoleMenus(id int64, menuIds []int64) error {
	return common.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Where("`role_id` = ?", id).Delete(&model.SysUserRole{}).Error
		if err != nil {
			common.LOG.Warn("删除sys_user_role表失败", zap.Error(err))
			return err
		}
		roleMenus := make([]model.SysRoleMenu, len(menuIds))
		for _, v := range menuIds {
			roleMenus = append(roleMenus, model.SysRoleMenu{RoleId: id, MenuId: v})
		}
		return tx.Create(&roleMenus).Error
	})
}

// 获取最大范围的数据权限
func (service *RoleService) GetMaximumDataScope(roles []string) (max int, err error) {
	if len(roles) > 0 {
		common.DB.Raw("SELECT min(data_scope) FROM sys_role where code in ?", roles).Find(&max)
	} else {
		common.DB.Raw("SELECT min(data_scope) FROM sys_role where id=-1").Find(&max)
	}
	return
}
