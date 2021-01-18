package handler

import (
	"micro-falsework/common"
	"strings"
)

// #Var_ToTitle# #Var_ToLower#

var Path = `/handler`
var FilePath = `/handler/#Var_ToLower#.go`
var Code= `
package handler

import (
    "#Var_ToLower#/domain/service"
)
type #Var_ToTitle#  struct{
     #Var_ToTitle#DataService service.I#Var_ToTitle#DataService
}
`
func Init(projectName string) {
	VarToTitle := strings.Title(projectName)
	VarToLower := strings.ToLower(projectName)
	common.MkDirAll("./" + VarToLower + common.PathOrFilePathFormat(Path, VarToLower))
	common.WriterFile("./"+VarToLower+common.PathOrFilePathFormat(FilePath, VarToLower), common.CodeFormat(Code, VarToTitle, VarToLower))
}
