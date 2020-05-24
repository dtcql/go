package main

import (
	"github.com/dtcql/go-log/filelogger"
)

func main() {
	//输入log日志的存储路径
	fLogger := filelogger.NewFileLogger("./")

	writer := "Jeffrey"
	msg := "This is a test message"

	fLogger.Log("ERROR", "%s. Writer:%s", msg, writer)
	fLogger.Log("INFO", "%s. Writer:%s", msg, writer)
	fLogger.Log("Debug", "%s. Writer:%s", msg, writer)
	fLogger.Log("FATAL", "%s. Writer:%s", msg, writer)

	//关闭文件
	fLogger.Close()
}
