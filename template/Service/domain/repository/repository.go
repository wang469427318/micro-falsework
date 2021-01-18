package repository

import (
	"micro-falsework/common"
	"strings"
)

// #Var_ToTitle# #Var_ToLower#

var Path = `/domain/repository`
var FilePath = `/domain/repository/#Var_ToLower#_repository.go`
var Code = `
package repository

import (
	"github.com/jinzhu/gorm"
	"#Var_ToLower#/domain/model"
)
type I#Var_ToTitle#Repository interface{
    InitTable() error
    Find#Var_ToTitle#ByID(int64) (*model.#Var_ToTitle#, error)
	Create#Var_ToTitle#(*model.#Var_ToTitle#) (int64, error)
	Delete#Var_ToTitle#ByID(int64) error
	Update#Var_ToTitle#(*model.#Var_ToTitle#) error
	FindAll()([]model.#Var_ToTitle#,error)

}
//创建#Var_ToLower#Repository
func New#Var_ToTitle#Repository(db *gorm.DB) I#Var_ToTitle#Repository  {
	return &#Var_ToTitle#Repository{mysqlDb:db}
}

type #Var_ToTitle#Repository struct {
	mysqlDb *gorm.DB
}

//初始化表
func (u *#Var_ToTitle#Repository)InitTable() error  {
	return u.mysqlDb.CreateTable(&model.#Var_ToTitle#{}).Error
}

//根据ID查找#Var_ToTitle#信息
func (u *#Var_ToTitle#Repository)Find#Var_ToTitle#ByID(#Var_ToLower#ID int64) (#Var_ToLower# *model.#Var_ToTitle#,err error) {
	#Var_ToLower# = &model.#Var_ToTitle#{}
	return #Var_ToLower#, u.mysqlDb.First(#Var_ToLower#,#Var_ToLower#ID).Error
}

//创建#Var_ToTitle#信息
func (u *#Var_ToTitle#Repository) Create#Var_ToTitle#(#Var_ToLower# *model.#Var_ToTitle#) (int64, error) {
	return #Var_ToLower#.ID, u.mysqlDb.Create(#Var_ToLower#).Error
}

//根据ID删除#Var_ToTitle#信息
func (u *#Var_ToTitle#Repository) Delete#Var_ToTitle#ByID(#Var_ToLower#ID int64) error {
	return u.mysqlDb.Where("id = ?",#Var_ToLower#ID).Delete(&model.#Var_ToTitle#{}).Error
}

//更新#Var_ToTitle#信息
func (u *#Var_ToTitle#Repository) Update#Var_ToTitle#(#Var_ToLower# *model.#Var_ToTitle#) error {
	return u.mysqlDb.Model(#Var_ToLower#).Update(#Var_ToLower#).Error
}

//获取结果集
func (u *#Var_ToTitle#Repository) FindAll()(#Var_ToLower#All []model.#Var_ToTitle#,err error) {
	return #Var_ToLower#All, u.mysqlDb.Find(&#Var_ToLower#All).Error
}

`

func Init(projectName string) {
	VarToTitle := strings.Title(projectName)
	VarToLower := strings.ToLower(projectName)
	common.MkDirAll("./" + VarToLower + common.PathOrFilePathFormat(Path, VarToLower))
	common.WriterFile("./"+VarToLower+common.PathOrFilePathFormat(FilePath, VarToLower), common.CodeFormat(Code, VarToTitle, VarToLower))
}
