package utils

import (
	"fmt"

	"github.com/beego/beego/v2/core/logs"
)

func Logger(fileName string, lineNumber int, requestId string, logLevel string, log string) {
	lineNumberStr := fmt.Sprintf("%d", lineNumber)
	switch logLevel {
	case "debug":
		logs.Debug("File: %s:%s - %s: %s", fileName, lineNumberStr, requestId, log)
	case "error":
		logs.Error("File: %s:%s - %s: %s", fileName, lineNumberStr, requestId, log)
	default:
		logs.Info("File: %s:%s - %s: %s", fileName, lineNumberStr, requestId, log)
	}
}
