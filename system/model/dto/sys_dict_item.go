package dto

type DictItemPageReq struct {
	Keywords string `form:"keywords"`
	TypeCode string `form:"typeCode"`
	PageNum  int    `form:"pageNum"`
	PageSize int    `form:"pageSize"`
}
