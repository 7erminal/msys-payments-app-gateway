package utils

import (
	"github.com/beego/beego/v2/core/logs"
)

func Logger(fileName string, lineNumber int, requestId string, logLevel string, log string) {
	switch logLevel {
	case "debug":
		logs.Debug("File: %s:%s - %s: %s", fileName, lineNumber, requestId, log)
	case "error":
		logs.Error("File: %s:%s - %s: %s", fileName, lineNumber, requestId, log)
	default:
		logs.Info("File: %s:%s - %s: %s", fileName, lineNumber, requestId, log)
	}
}
