package filelogger

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"time"
)

//Log Level 常量，日志级别，用于判断是否写到err log里
var logLevelMap = map[string]int64{
	"DEBUG": 0,
	"INFO":  1,
	"ERROR": 2,
	"FATAL": 3,
}

//FileLogger 文件日志结构体
type FileLogger struct {
	FilePath   string
	FileName   string
	fileObj    *os.File
	errFileObj *os.File
	lastDate   string
}

//NewFileLogger 创建文件日志结构体
func NewFileLogger(fp string) *FileLogger {
	fl := &FileLogger{
		FilePath: fp,
		FileName: "default",
	}
	err := fl.initFile()
	if err != nil {
		panic(err)
	}
	return fl
}

//打开YYYYMMDD和YYYYMMDD-Err日志文件
func (f *FileLogger) initFile() error {
	nowDate := time.Now().Format("20060102")
	fullFileName := path.Join(f.FilePath, nowDate)
	fileObj, err := os.OpenFile(fullFileName+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		fmt.Printf("open log file %s.log failed, err:%v\n", fullFileName, err)
		return err
	}
	errFileObj, err := os.OpenFile(fullFileName+"-Err.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		fmt.Printf("open error log file %sErr.log failed, err:%v\n", fullFileName, err)
		return err
	}
	f.fileObj = fileObj
	f.errFileObj = errFileObj
	return nil
}

//Close 关闭日志文件
func (f *FileLogger) Close() {
	f.fileObj.Close()
	f.errFileObj.Close()
}

//Log 向日志文件中写入日志，如果日志级别高于error，将写入-Err日志。
func (f *FileLogger) Log(lv string, format string, a ...interface{}) {
	//获得格式化后的日志消息
	msg := fmt.Sprintf(format, a...)

	//获得当前时间，并比较上一次日志日期，如日期相同，追加写入；日期不同则新创建新日期的日志文件。
	now := time.Now()
	if f.lastDate != now.Format("20060102") {
		f.Close()
		f.initFile()
	}

	//获得运行时信息并写入日志。
	funcName, prgName, lineNo := getRunTimeInfo(2)
	fmt.Fprintf(f.fileObj, "[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), lv, prgName, funcName, lineNo, msg)
	if logLevelMap[lv] > logLevelMap["INFO"] {
		fmt.Fprintf(f.errFileObj, "[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), lv, prgName, funcName, lineNo, msg)
	}
	//保存日志日期，留待下次循环开始比较日期
	f.lastDate = now.Format("20060102")
}

//getRunTimeInfo 得到运行时的信息
func getRunTimeInfo(callLayer int) (funcName, prgName string, lineNo int) {
	pc, prg, lineNo, ok := runtime.Caller(callLayer)
	if !ok {
		fmt.Printf("runtime.Caller failed!\n")
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	prgName = path.Base(prg)
	return
}
