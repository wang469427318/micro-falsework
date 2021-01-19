package Template

import (
	"micro-falsework/common"
	"strings"
)

// #Var_ToTitle# #Var_ToLower#
var MakePath = `/Makefile`
var MakeCode=`
.PHONY: proto
proto:
	protoc --plugin=protoc-gen-micro=/go/bin/protoc-gen-micro --proto_path=./ ./proto/#Var_ToLower#/#Var_ToLower#..proto --micro_out ./

.PHONY: build
build: 

	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o #Var_ToLower#-service *.go

.PHONY: docker
docker:
	docker build . -t #Var_ToLower#-service:latest
`
func MakeInit(projectName,projectPath string) {
	VarToTitle := strings.Title(projectName)
	VarToLower := strings.ToLower(projectName)
	common.WriterFile("./"+projectPath+common.PathOrFilePathFormat(MakePath, VarToLower), common.CodeFormat(MakeCode, VarToTitle, VarToLower,projectPath))
}