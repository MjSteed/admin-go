package service

import (
	"context"
	"strconv"

	"github.com/MjSteed/vue3-element-admin-go/common/model/vo"
	"github.com/MjSteed/vue3-element-admin-go/system/dao"
	"github.com/MjSteed/vue3-element-admin-go/system/model"
	"github.com/MjSteed/vue3-element-admin-go/system/model/dto"
	s_vo "github.com/MjSteed/vue3-element-admin-go/system/model/vo"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/exp/slices"
)

// 部门业务
type DeptService struct {
	deptDao *dao.DeptDao
	log     *zap.Logger
}

func NewDeptService(deptDao *dao.DeptDao, logger *zap.Logger) *DeptService {
	return &DeptService{deptDao: deptDao, log: logger}
}

// 根节点ID
const ROOT_NODE_ID = 0

// ListPages 部门列表分页
func (s *DeptService) ListPages(ctx *gin.Context, pageReq dto.DeptPageReq) (list []s_vo.Dept, err error) {
	depts, err := s.deptDao.ListPages(ctx, &pageReq)
	if len(depts) > 0 {
		var cacheDeptIds []int64
		for _, v := range depts {
			cacheDeptIds = append(cacheDeptIds, v.Id)
		}
		for _, v := range depts {
			parentId := v.ParentId
			//不在缓存ID列表的parentId是顶级节点ID，以此作为递归开始
			if slices.Contains(cacheDeptIds, parentId) {
				continue
			}
			list = append(list, s.recur(parentId, depts)...)
			cacheDeptIds = append(cacheDeptIds, parentId)
		}
	}
	if len(list) <= 0 {
		//列表为空说明所有的节点都是独立的
		for _, v := range depts {
			vo := s_vo.Dept{}
			list = append(list, vo.Format(v))
		}
	}
	return
}

// recurDepts 递归生成部门层级列表
func (s *DeptService) recur(parentId int64, depts []model.SysDept) (vos []s_vo.Dept) {
	for _, v := range depts {
		if v.ParentId != parentId {
			continue
		}
		vo := s_vo.Dept{}
		vo = vo.Format(v)
		vo.Children = s.recur(v.Id, depts)
		vos = append(vos, vo)
	}
	return
}

// Options 部门树形下拉选项
func (s *DeptService) Options(ctx *gin.Context) (list []vo.TreeOption, err error) {
	depts, err := s.deptDao.ListAll(ctx)
	if err != nil {
		return
	}
	list = s.recurTreeOptions(ROOT_NODE_ID, depts)
	return
}

// recurTreeOptions 递归生成部门表格层级列表
func (s *DeptService) recurTreeOptions(parentId int64, depts []model.SysDept) (options []vo.TreeOption) {
	if len(depts) <= 0 {
		return
	}
	for _, v := range depts {
		if v.ParentId != parentId {
			continue
		}
		op := vo.TreeOption{Label: v.Name, Value: v.Id}
		op.Children = s.recurTreeOptions(v.Id, depts)
		options = append(options, op)
	}
	return
}

// Create 新增部门
func (s *DeptService) Create(ctx *gin.Context, dept *model.SysDept) (id int64, err error) {
	dept.TreePath = s.generateTreePath(ctx, dept.ParentId)
	dept, err = s.deptDao.Save(ctx, dept)
	return dept.Id, err
}

// Update 修改部门
func (s *DeptService) Update(ctx *gin.Context, dept *model.SysDept) (id int64, err error) {
	dept.TreePath = s.generateTreePath(ctx, dept.ParentId)
	dept, err = s.deptDao.Save(ctx, dept)
	return dept.Id, err
}

func (s *DeptService) generateTreePath(ctx context.Context, parentId int64) (path string) {
	if ROOT_NODE_ID == parentId {
		path = strconv.FormatInt(parentId, 10)
	} else {
		parent, err := s.deptDao.FindById(ctx, parentId)
		if err != nil {
			return
		}
		path = parent.TreePath + "," + strconv.FormatInt(parent.Id, 10)
	}
	return
}

// 删除部门
// @param ids 部门id列表
func (s *DeptService) DeleteByIds(ctx *gin.Context, ids []int64) (err error) {
	err = s.deptDao.Delete(ctx, ids)
	return
}

// 获取部门详情
// @param id 部门id
func (s *DeptService) GetForm(ctx *gin.Context, id int64) (dept *model.SysDept, err error) {
	dept, err = s.deptDao.FindById(ctx, id)
	return
}
