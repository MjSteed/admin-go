package dto

type DictTypePageReq struct {
	Keywords string `form:"keywords"`
	PageNum  int    `form:"pageNum"`
	PageSize int    `form:"pageSize"`
}
