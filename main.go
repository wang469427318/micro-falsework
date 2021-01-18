package main

import (
	"log"
	Template "micro-falsework/template/Service"
	"micro-falsework/template/Service/domain/model"
	"micro-falsework/template/Service/domain/repository"
	"micro-falsework/template/Service/domain/service"
	"micro-falsework/template/Service/handler"
	proto "micro-falsework/template/Service/proto/product"
	"os"
	"strings"
)

// #Var_ToTitle# #Var_ToLower#
func main() {
	pathStr, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(pathStr)
	log.Println(strings.LastIndex(pathStr, "/"))
	log.Println(pathStr[strings.LastIndex(pathStr, "/")+1:])
	//common.MkDirAll(path)
	project := "user"
	model.Init(project)
	repository.Init(project)
	service.Init(project)
	handler.Init(project)
	proto.Init(project)
	Template.DockerInit(project)
	Template.MakeInit(project)
	Template.Init(project)
}
