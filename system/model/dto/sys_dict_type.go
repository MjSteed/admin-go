package dto

type DictTypePageReq struct {
	Name     string `form:"name"`
	PageNum  int    `form:"pageNum"`
	PageSize int    `form:"pageSize"`
}
