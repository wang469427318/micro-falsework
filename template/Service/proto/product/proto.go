package proto

import (
	"micro-falsework/common"
	"strings"
)

// #Var_ToTitle# #Var_ToLower# #Var_ProjectName#

var Path = `/proto/#Var_ToLower#/`
var FilePath = `/proto/#Var_ToLower#/#Var_ToLower#.proto`
var Code = `
syntax = "proto3";

package go.micro.service.#Var_ToLower#;

service #Var_ToTitle# {
	rpc Find(Request) returns (Response) {}
}

message Request {
	string id = 1;
}

message Response {
	string msg = 1;
	int32 code = 2;
}
`
func Init(projectName,projectPath string) {
	VarToTitle := strings.Title(projectName)
	VarToLower := strings.ToLower(projectName)
	common.MkDirAll("./" + projectPath + common.PathOrFilePathFormat(Path, VarToLower))
	common.WriterFile("./"+projectPath+common.PathOrFilePathFormat(FilePath, VarToLower), common.CodeFormat(Code, VarToTitle, VarToLower,projectPath))
}