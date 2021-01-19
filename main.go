package main

import (
	"flag"
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
	_pathStr := flag.String("path", "", "请填写路径")
	flag.Parse()
	pathStr := *_pathStr
	projectPath := *_pathStr
	osPathStr, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(osPathStr)
	log.Println(pathStr)
	log.Println(strings.LastIndex(pathStr, "/"))
	log.Println(pathStr[strings.LastIndex(pathStr, "/")+1:])
	project := pathStr[strings.LastIndex(pathStr, "/")+1:]

	model.Init(project, projectPath)
	repository.Init(project, projectPath)
	service.Init(project, projectPath)
	handler.Init(project, projectPath)
	proto.Init(project, projectPath)
	Template.DockerInit(project, projectPath)
	Template.MakeInit(project, projectPath)
	Template.Init(project, projectPath)
	Template.ModInit(project, projectPath)
}
