package dto

type UserPageReq struct {
	Keywords string `form:"keywords"`
	Status   string `form:"status"`
	DeptId   int64  `form:"deptId"`
	PageNum  int    `form:"pageNum"`
	PageSize int    `form:"pageSize"`
}
