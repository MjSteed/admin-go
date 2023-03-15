package utils

import "os"

// 检查目录是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// 创建目录，如果已存在目录，则跳过
func CreateDir(path string) {
	if ok, _ := PathExists(path); !ok {
		_ = os.Mkdir(path, os.ModePerm)
	}
}
