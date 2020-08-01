package file

import (
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
)

/**
获取文件大小
 */
func GetSize(f multipart.File) (int, error) {
	content, err := ioutil.ReadAll(f)

	return len(content), err
}
/**
获取后缀名
 */
func GetExt(fileName string) string {
	return path.Ext(fileName)
}
/**
判断文件是否存在
 */
func CheckExist(src string) bool {
	_, err := os.Stat(src)

	return os.IsNotExist(err)
}
/**
判断权限
 */
func CheckPermission(src string) bool {
	_, err := os.Stat(src)

	return os.IsPermission(err)
}
/**

 */
func IsNotExistMkDir(src string) error {
	if notExist := CheckExist(src); notExist == true {
		if err := MkDir(src); err != nil {
			return err
		}
	}

	return nil
}
/**
新建问价夹
 */
func MkDir(src string) error {
	err := os.MkdirAll(src, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
/**
打开文件
 */
func Open(name string, flag int, perm os.FileMode) (*os.File, error) {
	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}

	return f, nil
}