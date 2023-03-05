package vo

// 树形下拉选项对象
type TreeOption struct {
	Label    string
	Value    interface{}
	Children []TreeOption
}
