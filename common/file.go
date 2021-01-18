package common

import (
	"bufio"
	"log"
	"os"
	"strings"
)

const VARTOLOWER = "#Var_ToLower#"
const VARTOTITLE = "#Var_ToTitle#"

// 判断文件夹是否存在
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

// 创建目录
func MkDirAll(path string) {
	if ok, err := PathExists(path); ok {
		if err != nil {
			log.Fatal(err)
		}
		err := os.RemoveAll(path)
		if err != nil {
			log.Fatal(err)
		}
	}
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}

// 创建文件
func CreateFile(filePath string) (*os.File, error) {
	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return nil, err
	}
	return f, nil
}

// 写入数据
func WriterFile(filePath string, code string) {
	f, err := CreateFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	write := bufio.NewWriter(f)
	_, err = write.WriteString(code)
	if err != nil {
		log.Fatal(err)
	}
	err = write.Flush()
	if err != nil {
		log.Fatal(err)
	}
}

// 替换Code中的变量
func CodeFormat(sourceValue, varToTitle, varToLower string) (code string) {
	code = strings.ReplaceAll(sourceValue, VARTOLOWER, varToLower)
	code = strings.ReplaceAll(code, VARTOTITLE, varToTitle)
	return code
}
func PathOrFilePathFormat(sourceValue, varToLower string) string {
	return strings.ReplaceAll(sourceValue, VARTOLOWER, varToLower)
}
