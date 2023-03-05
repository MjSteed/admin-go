package dto

type DeptPageReq struct {
	Keywords string `form:"keywords"`
	Status   string `form:"status"`
	PageNum  int    `form:"pageNum"`
	PageSize int    `form:"pageSize"`
}
