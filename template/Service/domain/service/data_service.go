package service

import (
	"micro-falsework/common"
	"strings"
)

// #Var_ToTitle# #Var_ToLower# #Var_ProjectName#

var Path = `/domain/service`
var FilePath = `/domain/service/#Var_ToLower#_data_service.go`
var Code = `
package service


import (
	"#Var_ProjectName#/domain/model"
	"#Var_ProjectName#/domain/repository"
)

type I#Var_ToTitle#DataService interface {
	Add#Var_ToTitle#(*model.#Var_ToTitle#) (int64 , error)
	Delete#Var_ToTitle#(int64) error
	Update#Var_ToTitle#(*model.#Var_ToTitle#) error
	Find#Var_ToTitle#ByID(int64) (*model.#Var_ToTitle#, error)
	FindAll#Var_ToTitle#() ([]model.#Var_ToTitle#, error)
}


//创建
func New#Var_ToTitle#DataService(#Var_ToLower#Repository repository.I#Var_ToTitle#Repository) I#Var_ToTitle#DataService{
	return &#Var_ToTitle#DataService{ #Var_ToLower#Repository }
}

type #Var_ToTitle#DataService struct {
	#Var_ToTitle#Repository repository.I#Var_ToTitle#Repository
}


//插入
func (u *#Var_ToTitle#DataService) Add#Var_ToTitle#(#Var_ToLower# *model.#Var_ToTitle#) (int64 ,error) {
	 return u.#Var_ToTitle#Repository.Create#Var_ToTitle#(#Var_ToLower#)
}

//删除
func (u *#Var_ToTitle#DataService) Delete#Var_ToTitle#(#Var_ToLower#ID int64) error {
	return u.#Var_ToTitle#Repository.Delete#Var_ToTitle#ByID(#Var_ToLower#ID)
}

//更新
func (u *#Var_ToTitle#DataService) Update#Var_ToTitle#(#Var_ToLower# *model.#Var_ToTitle#) error {
	return u.#Var_ToTitle#Repository.Update#Var_ToTitle#(#Var_ToLower#)
}

//查找
func (u *#Var_ToTitle#DataService) Find#Var_ToTitle#ByID(#Var_ToLower#ID int64) (*model.#Var_ToTitle#, error) {
	return u.#Var_ToTitle#Repository.Find#Var_ToTitle#ByID(#Var_ToLower#ID)
}

//查找
func (u *#Var_ToTitle#DataService) FindAll#Var_ToTitle#() ([]model.#Var_ToTitle#, error) {
	return u.#Var_ToTitle#Repository.FindAll()
}
`


func Init(projectName,projectPath string) {
	VarToTitle := strings.Title(projectName)
	VarToLower := strings.ToLower(projectName)
	common.MkDirAll("./" + projectPath + common.PathOrFilePathFormat(Path, VarToLower))
	common.WriterFile("./"+projectPath+common.PathOrFilePathFormat(FilePath, VarToLower), common.CodeFormat(Code, VarToTitle, VarToLower,projectPath))
}