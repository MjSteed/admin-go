package vo

// 树形下拉选项对象
type TreeOption struct {
	Label    string       `json:"label"`
	Value    interface{}  `json:"value"`
	Children []TreeOption `json:"children"`
}
