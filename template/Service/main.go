package Template

import (
	"micro-falsework/common"
	"strings"
)

// #Var_ToTitle# #Var_ToLower# #Var_ProjectName#

var Path = `/main.go`
var Code = `
package main

import (
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2"
	"#Var_ProjectName#/handler"

	#Var_ToLower# "#Var_ProjectName#/proto/#Var_ToLower#"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.#Var_ToLower#"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	#Var_ToLower#.Register#Var_ToTitle#Handler(service.Server(), new(handler.#Var_ToTitle#))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
`

func Init(projectName, projectPath string) {
	VarToTitle := strings.Title(projectName)
	VarToLower := strings.ToLower(projectName)
	common.WriterFile("./"+projectPath+common.PathOrFilePathFormat(Path, VarToLower), common.CodeFormat(Code, VarToTitle, VarToLower, projectPath))
}
