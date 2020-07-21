package common

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// 检查文件或目录路径是否存在
func CheckPathExist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

// 如果文件不存在就创建
func CreateFileIfNotExist(filePath string) {
	var err error

	// 目录不存在就创建
	dir := filepath.Dir(filePath)
	if !CheckPathExist(dir) {
		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}

	// 文件不存在就创建
	if !CheckPathExist(filePath) {
		_, err = os.Create(filePath)
		if err != nil {
			panic(err)
		}
	}
}

// 添加内容到文件开头
func AppendContentToFileStart(filePath string, msg string) (err error) {
	// 读原来的内容
	buf, err := ioutil.ReadFile(filePath)
	f, err := os.OpenFile(filePath, os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer f.Close()

	// 写入文件
	_, err = f.Write([]byte(msg))
	if err != nil {
		return
	}
	_, err = f.Write(buf)
	if err != nil {
		return
	}

	return
}

// 添加内容到文件开头
func WriteMultiBufToFile(filePath string, buffers ...[]byte) (err error) {
	// 打开
	f, err := os.OpenFile(filePath, os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer f.Close()

	// 写
	for _, buf := range buffers {
		_, err = f.Write(buf)
		if err != nil {
			return
		}
	}

	return
}

// 清除文件内容
func ClearFileContent(filePath string) (err error) {
	//err = common.WriteMultiBufToFile(c.storageFilePath, []byte{0})
	err = ioutil.WriteFile(filePath, []byte{}, 0644)
	return
}
