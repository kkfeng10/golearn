package main

// log 控制打印输出的格式字段(const内容)
// const (
// 	// 将下面的位使用或运算符连接在一起，可以控制要输出的信息。没有
// 	// 办法控制这些信息出现的顺序（下面会给出顺序）或者打印的格式
// 	// （格式在注释里描述）。这些项后面会有一个冒号：
// 	// 2009/01/23 01:23:23.123123 /a/b/c/d.go:23: message
// 	// 日期: 2009/01/23
// 	Ldate = 1 << iota
// 	// 时间: 01:23:23
// 	Ltime
// 	// 毫秒级时间: 01:23:23.123123。该设置会覆盖 Ltime 标志
// 	Lmicroseconds
// 	// 完整路径的文件名和行号: /a/b/c/d.go:23
// 	Llongfile
// 	// 最终的文件名元素和行号: d.go:23
// 	// 覆盖 Llongfile
// 	Lshortfile
// 	// 标准日志记录器的初始值
// 	LstdFlags = Ldate | Ltime
// 	)

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

// // 1. 最简单的用法，直接输出到控制台日志
// func init() {
// 	log.SetPrefix("TRACE: ")
// 	log.SetFlags( log.Ldate | log.Ltime | log.Llongfile )
// }

// func main() {
// 	log.Println("message")
// 	log.Fatalln("fatal")
// 	log.Panicln("Panic")  // 直接退出
// }

// 2.需求：
// 1. debug日志忽略输出，但是必须要在代码中具备，且直接输出到控制台
// 2. info日志直接输出到控制台
// 3. warning日志直接输出到控制台
// 4. error日志除了要输出到控制台，也要记录在文件中

// var (
// 	// Debug 级别日志
// 	Debug *log.Logger
// 	// Info 级别日志
// 	Info *log.Logger
// 	// Warning 级别日志
// 	Warning *log.Logger
// 	// Error 级别日志
// 	Error *log.Logger
// )

// func init() {
// 	fileobj, err := os.OpenFile("./error.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0664)
// 	if err != nil {
// 		fmt.Println("打开文件失败")
// 	}
// 	defer fileobj.Close()
// 	// ioutil.Discard os.Stdout,io.MultiWriter 都实现了 io.Writer 接口的方法
// 	Debug = log.New(ioutil.Discard, "DEBUG: ", log.Ldate|log.Ltime|log.Llongfile)
// 	Info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Llongfile)
// 	Warning = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Llongfile)
// 	Error = log.New(io.MultiWriter(fileobj, os.Stderr), "ERROR: ", log.Ldate|log.Ltime|log.Llongfile)
// }

// func main() {
// 	Debug.Println("Debug log")
// 	Info.Println("Info log")
// 	Warning.Println("Warning log")
// 	Error.Println("Error log")
// }

// 使用struct 方式编写自己的log日志
type yflog struct {
	fileObj *os.File
}

var mylog yflog

func init() {
	fileObj, err := os.OpenFile("./error.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0664)
	if err != nil {
		fmt.Println("打开文件失败")
		os.Exit(-1)
	}
	mylog = yflog{fileObj: fileObj}
}

func (y *yflog) Debug(v interface{}) {
	debug := log.New(ioutil.Discard, "DEBUG: ", log.Ldate|log.Ltime|log.Llongfile)
	debug.Printf("%v\n", v)
}

func (y *yflog) Info(v interface{}) {
	debug := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Llongfile)
	debug.Printf("%v\n", v)
}
func (y *yflog) Warning(v interface{}) {
	debug := log.New(io.MultiWriter(os.Stdout, y.fileObj), "WARNING: ", log.Ldate|log.Ltime|log.Llongfile)
	debug.Printf("%v\n", v)
}
func (y *yflog) Error(v interface{}) {
	debug := log.New(io.MultiWriter(os.Stdout, y.fileObj), "ERROR: ", log.Ldate|log.Ltime|log.Llongfile)
	debug.Printf("%v\n", v)
}

func main() {
	defer mylog.fileObj.Close()
	mylog.Debug("很好，非常好")
	mylog.Warning("告警了，老哥")
	mylog.Error("很差，非常差")
	var s string
	rd := bufio.NewReader(os.Stdin)
	s, _ = rd.ReadString('\n') // 需要传入一个行终止符来等待用户输入
	mylog.Info(s)
}
