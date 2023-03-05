package service

import (
	"strconv"

	"github.com/MjSteed/vue3-element-admin-go/common"
	"github.com/MjSteed/vue3-element-admin-go/common/model/vo"
	"github.com/MjSteed/vue3-element-admin-go/system/model"
	"github.com/MjSteed/vue3-element-admin-go/system/model/dto"
	s_vo "github.com/MjSteed/vue3-element-admin-go/system/model/vo"
	"golang.org/x/exp/slices"
)

// 部门业务
type DeptService struct{}

// 根节点ID
const ROOT_NODE_ID = 0

// 部门列表
func (deptService *DeptService) ListDepts(pageReq dto.DeptPageReq) (list []s_vo.Dept, total int64, err error) {
	tx := common.DB.Model(&model.SysDept{})
	if pageReq.Keywords != "" {
		tx = tx.Where("`name` like ?", "%"+pageReq.Keywords+"%")
	}
	//TODO 0值处理
	if pageReq.Status != "" {
		tx = tx.Where("`status` = ?", pageReq.Status)
	}
	err = tx.Count(&total).Error
	if err != nil {
		return
	}
	var depts []model.SysDept
	err = tx.Order("`sort` ASC").Limit(pageReq.PageSize).Offset(pageReq.PageSize * (pageReq.PageNum - 1)).Find(&depts).Error
	if len(depts) > 0 {
		cacheDeptIds := make([]int64, len(depts))
		for _, v := range depts {
			parentId := v.ParentId
			//不在缓存ID列表的parentId是顶级节点ID，以此作为递归开始
			if slices.Contains(cacheDeptIds, parentId) {
				continue
			}
			list = append(list, deptService.recurDepts(parentId, depts)...)
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
	return list, total, err
}

// 递归生成部门层级列表
func (deptService *DeptService) recurDepts(parentId int64, depts []model.SysDept) (vos []s_vo.Dept) {
	for _, v := range depts {
		if v.ParentId != parentId {
			continue
		}
		vo := s_vo.Dept{}
		vo = vo.Format(v)
		vo.Children = deptService.recurDepts(v.Id, depts)
		vos = append(vos, vo)
	}
	return
}

// 部门树形下拉选项
func (deptService *DeptService) ListDeptOptions() (list []vo.TreeOption, err error) {
	var depts []model.SysDept
	err = common.DB.Model(&model.SysDept{}).Where("`status` = 1").Order("`sort` ASC").Find(&depts).Error
	if err != nil {
		return
	}
	list = deptService.recurDeptTreeOptions(ROOT_NODE_ID, depts)
	return
}

// 递归生成部门表格层级列表
func (deptService *DeptService) recurDeptTreeOptions(parentId int64, depts []model.SysDept) (options []vo.TreeOption) {
	if len(depts) <= 0 {
		return
	}
	for _, v := range depts {
		if v.ParentId != parentId {
			continue
		}
		op := vo.TreeOption{Label: v.Name, Value: v.Id}
		op.Children = deptService.recurDeptTreeOptions(v.Id, depts)
		options = append(options, op)
	}
	return
}

// 新增部门
func (deptService *DeptService) SaveDept(dept model.SysDept) (id int64, err error) {
	dept.TreePath = deptService.generateTreePath(dept.ParentId)
	err = common.DB.Model(&dept).Save(&dept).Error
	return dept.Id, err
}

// 修改部门
func (deptService *DeptService) UpdateDept(dept model.SysDept) (id int64, err error) {
	dept.TreePath = deptService.generateTreePath(dept.ParentId)
	err = common.DB.Model(&dept).Updates(&dept).Error
	return dept.Id, err
}

func (deptService *DeptService) generateTreePath(parentId int64) (path string) {
	if ROOT_NODE_ID == parentId {
		path = strconv.FormatInt(parentId, 10)
	} else {
		parent := model.SysDept{}
		err := common.DB.Model(&parent).First(&parent, parentId).Error
		if err != nil {
			return
		}
		path = parent.TreePath + "," + strconv.FormatInt(parent.Id, 10)
	}
	return
}

// 删除部门
// @param ids 部门id列表
func (deptService *DeptService) DeleteByIds(ids []int64) (err error) {
	err = common.DB.Model(&model.SysDept{}).Delete(ids).Error
	return
}

// 获取部门详情
// @param id 部门id
func (deptService *DeptService) GetDeptForm(id int64) (dept model.SysDept, err error) {
	err = common.DB.Model(&dept).First(&dept, id).Error
	return
}
