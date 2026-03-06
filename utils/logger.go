package utils

import (
	"github.com/beego/beego/v2/core/logs"
)

func Logger(requestId string, logLevel string, log string) {
	switch logLevel {
	case "debug":
		logs.Debug("%s: %s", requestId, log)
	case "error":
		logs.Error("%s: %s", requestId, log)
	default:
		logs.Info("%s: %s", requestId, log)
	}
}
