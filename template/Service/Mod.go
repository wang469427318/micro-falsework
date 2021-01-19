package Template

import (
	"micro-falsework/common"
	"strings"
)

// #Var_ToTitle# #Var_ToLower# #Var_ProjectName#

var ModPath = `/go.mod`
var ModCode = `
module #Var_ProjectName#

go 1.13

require (
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/protobuf v1.4.0
	github.com/jinzhu/gorm v1.9.16
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/registry/consul/v2 v2.9.1
	github.com/micro/go-plugins/wrapper/monitoring/prometheus/v2 v2.9.1
	github.com/micro/go-plugins/wrapper/ratelimiter/uber/v2 v2.9.1
	github.com/micro/go-plugins/wrapper/trace/opentracing/v2 v2.9.1
	github.com/opentracing/opentracing-go v1.2.0
)
`

func ModInit(projectName, projectPath string) {
	VarToLower := strings.ToLower(projectName)
	common.WriterFile("./"+projectPath+common.PathOrFilePathFormat(ModPath, VarToLower), common.CodeFormat(ModCode, "", "", projectPath))
}
