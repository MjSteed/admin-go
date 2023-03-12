package dto

type DictItemPageReq struct {
	Name     string `form:"name"`
	TypeCode string `form:"typeCode"`
	PageNum  int    `form:"pageNum"`
	PageSize int    `form:"pageSize"`
}
