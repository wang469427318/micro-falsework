package model

import (
	"fmt"
	"micro-falsework/common"
	"strings"
)

// #Var_ToTitle# #Var_ToLower#
var Path = `/domain/model`
var FilePath = `/domain/model/#Var_ToLower#.go`
var Code = fmt.Sprintf(
	`
package model

type #Var_ToTitle# struct{
	ID int64 %v
}`, StrGorm)

var StrGorm = "`gorm:\"primary_key;not_null;auto_increment\"`"

func Init(projectName string) {
	VarToTitle := strings.Title(projectName)
	VarToLower := strings.ToLower(projectName)
	common.MkDirAll("./" + VarToLower + common.PathOrFilePathFormat(Path, VarToLower))
	common.WriterFile("./"+VarToLower+common.PathOrFilePathFormat(FilePath, VarToLower), common.CodeFormat(Code, VarToTitle, VarToLower))
}
