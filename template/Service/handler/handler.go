package handler

import (
	"micro-falsework/common"
	"strings"
)

// #Var_ToTitle# #Var_ToLower# #Var_ProjectName#

var Path = `/handler`
var FilePath = `/handler/#Var_ToLower#.go`
var Code= `
package handler

import (
    "#Var_ProjectName#/domain/service"
)
type #Var_ToTitle#  struct{
     #Var_ToTitle#DataService service.I#Var_ToTitle#DataService
}
`
func Init(projectName,projectPath string) {
	VarToTitle := strings.Title(projectName)
	VarToLower := strings.ToLower(projectName)
	common.MkDirAll("./" + projectPath + common.PathOrFilePathFormat(Path, VarToLower))
	common.WriterFile("./"+projectPath+common.PathOrFilePathFormat(FilePath, VarToLower), common.CodeFormat(Code, VarToTitle, VarToLower,projectPath))
}
