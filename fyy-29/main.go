package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 读取文件
// []byte 自定义缓存读取数据
// bufio逐行读取
// ioutil 读取所有文件内容

// []byte 自定义缓存读取数据
func readFromByte(fpath string) {
	file, err := os.Open(fpath)
	defer file.Close()
	if err != nil {
		fmt.Printf("打开文件 %s 错误, %v\n", fpath, err)
		return
	}
	var buf []byte
	buf = make([]byte, 128)
	for {
		n, err := file.Read(buf) // 读取的数据会在buf数组当中，返回读取的字节个数
		if err == io.EOF {
			fmt.Println("文件读取已经完成")
			return
		}
		if err != nil {
			fmt.Println("读取文件失败,失败原因: ", err)
			return
		}
		fmt.Printf("读取了%d个字节\n", n)
		fmt.Println(string(buf[:n])) // 取出buf中的数据并转换成string类型
	}
}

// bufio逐行读取
func readFrombufio(fpath string) {
	file, err := os.Open(fpath)
	defer file.Close()
	if err != nil {
		fmt.Printf("打开文件 %s 错误, %v\n", fpath, err)
		return
	}
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n') // 返回字符串
		if err == io.EOF {
			fmt.Println("文件读取已经完成")
			return
		}
		if err != nil {
			fmt.Println("读取文件内容失败，", err)
			return
		}
		fmt.Print(str) // 注意，本身自己会分行读写
	}
}

// ioutil读取文件内容
func readFromIoutil(fpath string) {
	buf, err := ioutil.ReadFile(fpath)
	if err != nil {
		fmt.Println("读取文件内容失败，", err)
		return
	}
	fmt.Print(string(buf)) // 返回值是[]byte类型
}

func main() {
	//readFromByte("./main.go")
	//readFrombufio("./main.go")
	readFromIoutil("./maifn.go")
}
