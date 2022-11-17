package gofsextra

import (
	"os"

	"github.com/markthree/go-fs-extra/utils"
)

// 复制文件
func CopyFile(src, dest string) error {
	printError := utils.CreatePrintError("CopyFile")
	data, err := os.ReadFile(src)
	if err != nil {
		printError(err)
		return err
	}
	return os.WriteFile(dest, data, os.ModePerm)
}

// 是否是文件夹
func IsDir(path string) bool {
	printError := utils.CreatePrintError("IsDir")
	fi, err := os.Stat(path)
	if err != nil {
		printError(err)
		return false
	}
	return fi.IsDir()
}

// 写文件
func WriteFile[T string | []byte](path string, data T) error {
	printError := utils.CreatePrintError("WriteFile")
	err := os.WriteFile(path, []byte(data), os.ModePerm)
	if err != nil {
		printError(err)
		return err
	}
	return nil
}

// 读文件
func ReadFile[T string | []byte](path string) (T, error) {
	printError := utils.CreatePrintError("ReadFile")
	data, err := os.ReadFile(path)
	if err != nil {
		printError(err)
		return T(""), err
	}
	return T(data), nil
}

// TODO
// 确保文件存在
func EnsureFile(path string) error {
	return nil
}

// TODO
// 确保目录存在
func EnsureDir(path string) error {
	return nil
}

// 路径是否存在
func PathExists(path string) bool {
	printError := utils.CreatePrintError("PathExists")
	_, err := os.Stat(path)

	if err == nil {
		return true
	} else if os.IsExist(err) {
		return true
	} else if os.IsNotExist(err) {
		return false
	} else {
		printError(err)
		return false
	}
}
