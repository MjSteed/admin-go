package service

import (
	"errors"
	"fmt"
	"strings"

	"github.com/MjSteed/vue3-element-admin-go/common"
	"github.com/MjSteed/vue3-element-admin-go/common/model/vo"
	"github.com/MjSteed/vue3-element-admin-go/system/model"
	"github.com/MjSteed/vue3-element-admin-go/system/model/dto"
	s_vo "github.com/MjSteed/vue3-element-admin-go/system/model/vo"
	"go.uber.org/zap"
	"golang.org/x/exp/slices"
)

// 菜单业务接口
type MenuService struct{}

// 获取菜单表格列表
func (service *MenuService) ListPages(pageReq dto.DeptPageReq) (list []s_vo.Menu, err error) {
	common.LOG.Debug("查询菜单表格参数", zap.Int("PageNum", pageReq.PageNum), zap.Int("PageSize", pageReq.PageSize), zap.String("Keywords", pageReq.Keywords))
	tx := common.DB.Model(&model.SysMenu{})
	if pageReq.Keywords != "" {
		tx = tx.Where("`name` like ?", "%"+pageReq.Keywords+"%")
	}
	var menus []model.SysMenu
	err = tx.Order("`sort` ASC").Find(&menus).Error
	if err != nil {
		return
	}
	if len(menus) > 0 {
		cacheIds := make([]int64, len(menus))
		for _, v := range menus {
			parentId := v.ParentId
			//不在缓存ID列表的parentId是顶级节点ID，以此作为递归开始
			if slices.Contains(cacheIds, parentId) {
				continue
			}
			list = append(list, service.recur(parentId, menus)...)
			cacheIds = append(cacheIds, parentId)
		}
	}
	if len(list) <= 0 {
		//列表为空说明所有的节点都是独立的
		for _, v := range menus {
			vo := s_vo.Menu{}
			list = append(list, vo.Format(v))
		}
	}
	return list, err
}

// 递归生成部门层级列表
func (service *MenuService) recur(parentId int64, menus []model.SysMenu) (vos []s_vo.Menu) {
	for _, v := range menus {
		if v.ParentId != parentId {
			continue
		}
		vo := s_vo.Menu{}
		vo = vo.Format(v)
		vo.Children = service.recur(v.Id, menus)
		vos = append(vos, vo)
	}
	return
}

// 获取菜单下拉列表
func (service *MenuService) ListOptions() (list []vo.TreeOption) {
	var menus []model.SysMenu
	err := common.DB.Model(&model.SysMenu{}).Order("`sort` ASC").Find(&menus).Error
	if err != nil {
		return
	}
	list = service.recurTreeOptions(ROOT_NODE_ID, menus)
	return
}

// 递归生成菜单下拉层级列表
func (service *MenuService) recurTreeOptions(parentId int64, menus []model.SysMenu) (options []vo.TreeOption) {
	if len(menus) <= 0 {
		return
	}
	for _, v := range menus {
		if v.ParentId != parentId {
			continue
		}
		op := vo.TreeOption{Label: v.Name, Value: v.Id, Children: service.recurTreeOptions(v.Id, menus)}
		options = append(options, op)
	}
	return
}

// 保存菜单
func (service *MenuService) Save(data model.SysMenu) (err error) {
	switch data.Type {
	case 2:
		//目录
		if strings.HasPrefix(data.Path, "/") {
			data.Component = "Layout"
		} else {
			return errors.New("目录路由路径格式错误，必须以/开始")
		}
	case 3:
		//外链
		data.Component = ""
	}
	if data.Id > 0 {
		err = common.DB.Model(&data).Updates(&data).Error
	} else {
		err = common.DB.Model(&data).Save(&data).Error
	}
	return
}

// 批量刪除
func (service *MenuService) DeleteByIds(ids []int64) error {
	return common.DB.Model(&model.SysMenu{}).Delete(ids).Error
}

// 路由列表
func (service *MenuService) ListRoutes() []s_vo.Route {
	//TODO 增加缓存
	sql := `SELECT
			t1.id,
			t1.name,
			t1.parent_id,
           	t1.path,
           	t1.component,
           	t1.icon,
			t1.sort,
			t1.visible,
			t1.redirect_url,
			t1.type,
			t3.code
		FROM
			sys_menu t1
				LEFT JOIN sys_role_menu t2 ON t1.id = t2.menu_id
				LEFT JOIN sys_role t3 ON t2.role_id = t3.id
		WHERE
			t1.type != ?
		ORDER BY t1.sort asc`
	var menus []model.Route
	err := common.DB.Raw(sql, 4).Scan(&menus).Error
	if err != nil {
		return nil
	}
	return service.recurRoutes(ROOT_NODE_ID, menus)
}

// 递归生成菜单路由层级列表
func (service *MenuService) recurRoutes(parentId int64, menus []model.Route) (list []s_vo.Route) {
	if len(menus) <= 0 {
		return
	}
	for _, v := range menus {
		if v.ParentId != parentId {
			continue
		}
		vo := s_vo.Route{
			Path:      v.Path,
			Redirect:  v.RedirectUrl,
			Component: v.Component,
			Meta: s_vo.Meta{
				Title:     v.Name,
				Icon:      v.Icon,
				Roles:     v.Roles,
				Hidden:    v.Visible == 0,
				KeepAlive: true,
			},
		}
		if v.Type == 1 {
			vo.Name = v.Path
		}
		children := service.recurRoutes(v.Id, menus)
		alwaysShow := false
		for _, c := range children {
			if !c.Meta.Hidden {
				// 含有子节点的目录设置为可见
				alwaysShow = true
				break
			}
		}
		vo.Meta.AlwaysShow = alwaysShow
		vo.Children = children
		list = append(list, vo)
	}
	return
}

// 获取菜单资源树形列表
func (service *MenuService) ListResources() []vo.TreeOption {
	return service.ListOptions()
}

// 修改菜单显示状态
// @param id 菜单id
// @param visible 是否显示(1->显示；2->隐藏)
func (service *MenuService) UpdateVisible(id int64, visible int) bool {
	err := common.DB.Model(&model.SysMenu{Id: id}).Update("visible", visible).Error
	if err != nil {
		fmt.Println("修改失败")
		return false
	}
	return true
}

func (service *MenuService) ListRolePerms(roles []string) (perms []string) {
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
	if len(roles) > 0 {
		condition += " AND t3.CODE IN ("
		for i, v := range roles {
			if i != 0 {
				condition += ","
			}
			condition += "'" + v + "'"
		}
		condition += ")"
	} else {
		condition += "AND t1.id = -1"
	}
	err := common.DB.Raw(sql+condition, 4).Scan(&perms).Error
	if err != nil {
		return
	}
	return
}
