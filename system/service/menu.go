package service

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/MjSteed/vue3-element-admin-go/common"
	"github.com/MjSteed/vue3-element-admin-go/common/model/vo"
	"github.com/MjSteed/vue3-element-admin-go/system/model"
	"github.com/MjSteed/vue3-element-admin-go/system/model/dto"
	s_vo "github.com/MjSteed/vue3-element-admin-go/system/model/vo"
	"go.uber.org/zap"
	"golang.org/x/exp/slices"
)

// 菜单业务接口
type menuService struct{}

var MenuService = new(menuService)

const (
	router_cache_key = "sys:routers"
)

// 获取菜单表格列表
func (service *menuService) ListPages(pageReq dto.DeptPageReq) (list []s_vo.Menu, err error) {
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
		var cacheIds []int64
		for _, v := range menus {
			cacheIds = append(cacheIds, v.Id)
		}
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

// 根据id获取菜单详情
func (service *menuService) GetById(id int64) (s_vo.SysMenu, error) {
	var m model.SysMenu
	err := common.DB.First(&m, id).Error
	if err != nil {
		return s_vo.SysMenu{}, err
	}
	vo := s_vo.SysMenu{
		Id:          m.Id,
		Name:        m.Name,
		ParentId:    m.ParentId,
		Type:        m.Type.String(),
		Path:        m.Path,
		Component:   m.Component,
		Perm:        m.Perm,
		Visible:     m.Visible,
		Sort:        m.Sort,
		Icon:        m.Icon,
		RedirectUrl: m.RedirectUrl,
		CreateTime:  m.CreateTime,
		UpdateTime:  m.UpdateTime,
	}
	return vo, nil
}

// 递归生成部门层级列表
func (service *menuService) recur(parentId int64, menus []model.SysMenu) (vos []s_vo.Menu) {
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
func (service *menuService) ListOptions() (list []vo.TreeOption) {
	var menus []model.SysMenu
	err := common.DB.Model(&model.SysMenu{}).Order("`sort` ASC").Find(&menus).Error
	if err != nil {
		return
	}
	list = service.recurTreeOptions(ROOT_NODE_ID, menus)
	return
}

// 递归生成菜单下拉层级列表
func (service *menuService) recurTreeOptions(parentId int64, menus []model.SysMenu) (options []vo.TreeOption) {
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
func (service *menuService) Save(data *model.SysMenu) (err error) {
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
		err = common.DB.Model(&data).Create(&data).Error
	}
	if err != nil {
		return
	}
	common.CacheDel(router_cache_key)
	return
}

// 批量刪除
func (service *menuService) DeleteByIds(ids []int64) error {
	err := common.DB.Where("id in ?", ids).Delete(&model.SysMenu{}).Error
	if err != nil {
		return err
	}
	common.CacheDel(router_cache_key)
	return err
}

// 路由列表
func (service *menuService) ListRoutes() (vos []s_vo.Route) {
	err := common.CacheGet(router_cache_key, &vos)
	if err != nil {
		return nil
	}
	if len(vos) > 0 {
		common.LOG.Debug("路由缓存获取成功")
		return
	}
	common.LOG.Debug("缓存获取失败，从数据库获取路由")
	var menus []model.SysMenu
	err = common.DB.Model(&model.SysMenu{}).Where("type != ?", model.BUTTON).Preload("SysRoles").Find(&menus).Error
	if err != nil {
		return nil
	}
	var routes []model.Route
	for _, v := range menus {
		r := model.Route{
			Id:          v.Id,
			ParentId:    v.ParentId,
			Name:        v.Name,
			Type:        v.Type.String(),
			Path:        v.Path,
			Component:   v.Component,
			Perm:        v.Perm,
			Visible:     v.Visible,
			Sort:        v.Sort,
			Icon:        v.Icon,
			RedirectUrl: v.RedirectUrl,
		}
		if len(v.SysRoles) > 0 {
			var roles []string
			for _, sr := range v.SysRoles {
				roles = append(roles, sr.Code)
			}
			r.Roles = roles
		}
		routes = append(routes, r)
	}
	vos = service.recurRoutes(ROOT_NODE_ID, routes)
	err = common.CacheSet(router_cache_key, &vos, time.Minute*10)
	if err != nil {
		common.LOG.Error("设置缓存错误", zap.Error(err))
	}
	return
}

// 递归生成菜单路由层级列表
func (service *menuService) recurRoutes(parentId int64, menus []model.Route) (list []s_vo.Route) {
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
		if v.Type == model.MENU.String() {
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
func (service *menuService) ListResources() []vo.TreeOption {
	return service.ListOptions()
}

// 修改菜单显示状态
// @param id 菜单id
// @param visible 是否显示(1->显示；2->隐藏)
func (service *menuService) UpdateVisible(id int64, visible int) bool {
	err := common.DB.Model(&model.SysMenu{Id: id}).Update("visible", visible).Error
	if err != nil {
		fmt.Println("修改失败")
		return false
	}
	return true
}

func (service *menuService) ListRolePerms(roles []string) (perms []string) {
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
