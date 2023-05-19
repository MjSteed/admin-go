package service

import (
	"errors"
	"strings"
	"time"

	"github.com/MjSteed/vue3-element-admin-go/common"
	"github.com/MjSteed/vue3-element-admin-go/common/model/vo"
	"github.com/MjSteed/vue3-element-admin-go/system/dao"
	"github.com/MjSteed/vue3-element-admin-go/system/model"
	"github.com/MjSteed/vue3-element-admin-go/system/model/dto"
	s_vo "github.com/MjSteed/vue3-element-admin-go/system/model/vo"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/exp/slices"
)

// MenuService 菜单服务
type MenuService struct {
	log     *zap.Logger
	menuDao *dao.MenuDao
}

// NewMenuService 实例化
func NewMenuService(log *zap.Logger, menuDao *dao.MenuDao) *MenuService {
	return &MenuService{log: log, menuDao: menuDao}
}

const (
	// 菜单缓存key
	router_cache_key = "sys:routers"
)

// ListPages 获取菜单表格列表
func (s *MenuService) ListPages(ctx *gin.Context, pageReq dto.DeptPageReq) (list []s_vo.Menu, err error) {
	s.log.Debug("查询菜单表格参数", zap.Int("PageNum", pageReq.PageNum), zap.Int("PageSize", pageReq.PageSize), zap.String("Keywords", pageReq.Keywords))
	menus, err := s.menuDao.Lists(ctx, pageReq)
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
			list = append(list, s.recur(ctx, parentId, menus)...)
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

// GetById 根据id获取菜单详情
func (s *MenuService) GetById(ctx *gin.Context, id int64) (s_vo.SysMenu, error) {
	m, err := s.menuDao.FindByID(ctx, id)
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

// recur 递归生成部门层级列表
func (s *MenuService) recur(ctx *gin.Context, parentId int64, menus []model.SysMenu) (vos []s_vo.Menu) {
	for _, v := range menus {
		if v.ParentId != parentId {
			continue
		}
		vo := s_vo.Menu{}
		vo = vo.Format(v)
		vo.Children = s.recur(ctx, v.Id, menus)
		vos = append(vos, vo)
	}
	return
}

// Options 获取菜单下拉列表
func (s *MenuService) Options(ctx *gin.Context) (list []vo.TreeOption) {
	var menus []model.SysMenu
	err := common.DB.Model(&model.SysMenu{}).Order("`sort` ASC").Find(&menus).Error
	if err != nil {
		return
	}
	list = s.recurTreeOptions(ctx, ROOT_NODE_ID, menus)
	return
}

// recurTreeOptions 递归生成菜单下拉层级列表
func (s *MenuService) recurTreeOptions(ctx *gin.Context, parentId int64, menus []model.SysMenu) (options []vo.TreeOption) {
	if len(menus) <= 0 {
		return
	}
	for _, v := range menus {
		if v.ParentId != parentId {
			continue
		}
		op := vo.TreeOption{Label: v.Name, Value: v.Id, Children: s.recurTreeOptions(ctx, v.Id, menus)}
		options = append(options, op)
	}
	return
}

// Save 保存菜单
func (s *MenuService) Save(ctx *gin.Context, data *model.SysMenu) (err error) {
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
	s.menuDao.Save(ctx, *data)
	if err != nil {
		return
	}
	common.CacheDel(router_cache_key)
	return
}

// DeleteByIds 批量刪除
func (s *MenuService) DeleteByIds(ctx *gin.Context, ids []int64) error {
	err := s.menuDao.Delete(ctx, ids)
	if err != nil {
		return err
	}
	common.CacheDel(router_cache_key)
	return err
}

// ListRoutes 路由列表
func (s *MenuService) ListRoutes(ctx *gin.Context) (vos []s_vo.Route) {
	err := common.CacheGet(router_cache_key, &vos)
	if err != nil {
		return nil
	}
	if len(vos) > 0 {
		s.log.Debug("路由缓存获取成功")
		return
	}
	s.log.Debug("缓存获取失败，从数据库获取路由")
	menus, err := s.menuDao.ListRoutes(ctx)
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
	vos = s.recurRoutes(ctx, ROOT_NODE_ID, routes)
	err = common.CacheSet(router_cache_key, &vos, time.Minute*10)
	if err != nil {
		s.log.Error("设置缓存错误", zap.Error(err))
	}
	return
}

// recurRoutes 递归生成菜单路由层级列表
func (s *MenuService) recurRoutes(ctx *gin.Context, parentId int64, menus []model.Route) (list []s_vo.Route) {
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
		children := s.recurRoutes(ctx, v.Id, menus)
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

// ListResources 获取菜单资源树形列表
func (s *MenuService) ListResources(ctx *gin.Context) []vo.TreeOption {
	return s.Options(ctx)
}

// UpdateVisible 修改菜单显示状态
// @param id 菜单id
// @param visible 是否显示(1->显示；2->隐藏)
func (s *MenuService) UpdateVisible(ctx *gin.Context, id int64, visible int) bool {
	err := s.menuDao.UpdateVisible(ctx, id, visible)
	if err != nil {
		return false
	}
	return true
}

// ListRolePerms 获取角色权限列表
func (s *MenuService) ListRolePerms(ctx *gin.Context, roles []string) (perms []string) {
	perms, _ = s.menuDao.GetPermsByRoles(ctx, roles)
	return
}
