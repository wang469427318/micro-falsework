package Template

import (
	"micro-falsework/common"
	"strings"
)

// #Var_ToTitle# #Var_ToLower#

var DockerPath = `/Dockerfile`
var DockerCode =`
FROM alpine
ADD #Var_ToLower#-service /#Var_ToLower#-service
ENTRYPOINT [ "/#Var_ToLower#-service" ]
`

func DockerInit(projectName string) {
	VarToTitle := strings.Title(projectName)
	VarToLower := strings.ToLower(projectName)
	common.WriterFile("./"+VarToLower+common.PathOrFilePathFormat(DockerPath, VarToLower), common.CodeFormat(DockerCode, VarToTitle, VarToLower))
}