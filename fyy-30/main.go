package main

<<<<<<< HEAD
// 写入文件 , os.OpenFile(name string, flag int, perm FileMode) (*File,error){...}

// flag 的选项如下,使用16进制表示，可以使用逻辑运算来表示这个flag的值
// os.O_WRONLY 只写
// os.O_CREATE 创建文件
// os.O_RDONLY 只读
// os.O_RDWR  读写
// os.O_TRUNC  清空
// os.O_APPEND  追加

// perm： 文件权限， 一个八进制数， r,w,x 分别对应为 读，写，执行
=======
import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

// 文件写入

// 1. []byte方式写入
func writeFromByte(){
	fileObj,err := os.OpenFile("./error.txt",os.O_WRONLY|os.O_CREATE|os.O_APPEND,0644)
	defer fileObj.Close()
	if err != nil{
		fmt.Println("打开写入文件时发生错误")
		return
	}
	n,err:=fileObj.WriteString("abcdefg\n")
	if err !=nil{
		fmt.Println("写入字符错误")
		return
	}
	fmt.Println("写入的字节数为: ",n)
}

// 2. bufio方式写入
func writeFromBufio(){
	fileObj,err := os.OpenFile("./error.txt",os.O_WRONLY|os.O_CREATE|os.O_APPEND,0644)
	defer fileObj.Close()
	writer:=bufio.NewWriter(fileObj)
	n,err:=writer.WriteString("bufio write\n")
	if err !=nil{
		fmt.Println("写入字符错误")
		return
	}
	fmt.Println("写入的字节数为: ",n)
	writer.Flush()
}

// 3. ioutil方式写入, 似乎无法追加写入
func writeFromIoutil(){
	err:=ioutil.WriteFile("./error.txt",[]byte("write from ioutil"),0644)
	if err != nil{
		fmt.Println("写入字符错误")
		return
	}
}

func main(){
	//writeFromByte()
	//writeFromBufio()
	writeFromIoutil()
}
>>>>>>> 307a332f9e84cbfa6e57944e315cab6e0999f332
